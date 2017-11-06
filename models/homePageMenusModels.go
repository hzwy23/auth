package models

import (
	"encoding/json"

	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/auth-core/entity"
)

const redirect = `
<script type="text/javascript">
    $.Hconfirm({
		cancelBtn:false,
        header:"Get Special Page Failure",
        body:"Get special page failed, Please contact your administrator",
        callback:function () {
            window.location.href="/"
        }
    })
</script>
`

type HomePageMenusModel struct {
	mur UserRolesModel
	mut UserThemeModel
	mts ThemeResourceModel
	mrs RoleAndResourceModel
	ms  ResourceModel
}

func (this *HomePageMenusModel) Get(id, typeId, useId string) ([]byte, error) {

	// 首先获取用户主题信息
	theme, err := this.mut.Get(useId)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// 获取这个主题的所有资源信息
	theme_resource, err := this.mts.Get(theme.ThemeId)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var mres = make(map[string]entity.ResData)

	// 如果是超级管理员用户，
	// 获取系统中所有的菜单信息
	if utils.IsAdmin(useId) {
		resdata, err := this.ms.GetChildren(id)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range resdata {
			if val.Restype == typeId {
				mres[val.ResId] = val
			}
		}
	} else {
		// 获取这个用户的角色信息
		roles, err := this.mur.GetRolesByUser(useId)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		var role_list []string
		for _, val := range roles {
			role_list = append(role_list, val.Role_id)
		}

		// 获取角色拥有的资源信息
		role_resource, err := this.mrs.Gets(role_list, id, typeId)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range role_resource {
			mres[val.ResId] = val
		}
	}

	// 获取角色拥有的资源信息
	var rst []entity.HomePageMenuData
	for _, t_res := range theme_resource {
		if val, ok := mres[t_res.ResId]; ok {
			var one entity.HomePageMenuData
			one.Res_id = t_res.ResId
			one.Res_up_id = val.ResUpid
			one.Res_name = val.ResName
			one.Group_id = t_res.GroupId
			one.Res_bg_color = t_res.ResBgColor
			one.Res_class = t_res.ResClass
			one.Res_img = t_res.ResImg
			one.Res_url = t_res.ResUrl
			one.Res_open_type = t_res.ResOpenType
			one.New_iframe = t_res.NewIframe
			rst = append(rst, one)
		}
	}
	return json.Marshal(rst)
}

func (this *HomePageMenusModel) GetUrl(user_id, id string) (string, error) {
	row := dbobj.QueryRow(sys_rdbms_011, user_id, id)
	var url string
	err := row.Scan(&url)
	if err != nil {
		logger.Error("Get Special Page failure, user_id is :", user_id, "menu id is :", id, "details error is:", err)
		return redirect, err
	}
	return url, nil
}

func (this *HomePageMenusModel) GetAllMenusExceptButton(userId string, menuId string) ([]byte, error) {
	// 首先获取用户主题信息
	theme, err := this.mut.Get(userId)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// 获取这个主题的所有资源信息
	theme_resource, err := this.mts.Get(theme.ThemeId)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var mres = make(map[string]entity.ResData)
	var mtheme = make(map[string]bool)
	for _, val := range theme_resource {
		mtheme[val.ResId] = true
	}

	var rst []entity.HomePageMenuData

	// 如果是超级管理员用户，
	// 获取系统中所有的菜单信息
	if utils.IsAdmin(userId) {
		// 获取所有的下级资源
		resdata, err := this.ms.GetChildExceptButton(menuId)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range resdata {
			// 当菜单资源没有配置主题信息，或者是菜单结点时
			if _, yes := mtheme[val.ResId]; !yes || val.Restype == "4" {
				var one entity.HomePageMenuData
				one.Res_id = val.ResId
				one.Res_up_id = val.ResUpid
				one.Res_name = val.ResName
				one.Method = val.Method
				one.Res_attr = val.ResAttr
				rst = append(rst, one)
			} else {
				mres[val.ResId] = val
			}
		}
	} else {
		// 获取这个用户的角色信息
		roles, err := this.mur.GetRolesByUser(userId)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		var role_list []string
		for _, val := range roles {
			role_list = append(role_list, val.Role_id)
		}

		// 获取角色拥有的资源信息
		role_resource, err := this.mrs.Gets(role_list, menuId)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range role_resource {
			if val.Restype == "2" {
				continue
			} else if _, yes := mtheme[val.ResId]; !yes || val.Restype == "4" {
				var one entity.HomePageMenuData
				one.Res_id = val.ResId
				one.Res_up_id = val.ResUpid
				one.Res_name = val.ResName
				one.Method = val.Method
				one.Res_attr = val.ResAttr
				rst = append(rst, one)
			} else {
				mres[val.ResId] = val
			}
		}
	}

	// 获取角色拥有的资源信息
	for _, t_res := range theme_resource {
		if val, ok := mres[t_res.ResId]; ok {
			var one entity.HomePageMenuData
			one.Res_id = t_res.ResId
			one.Res_up_id = val.ResUpid
			one.Res_name = val.ResName
			one.Group_id = t_res.GroupId
			one.Res_bg_color = t_res.ResBgColor
			one.Res_class = t_res.ResClass
			one.Res_img = t_res.ResImg
			one.Res_url = t_res.ResUrl
			one.Res_open_type = t_res.ResOpenType
			one.New_iframe = t_res.NewIframe
			one.Method = val.Method
			one.Res_attr = val.ResAttr
			rst = append(rst, one)
		}
	}

	return json.Marshal(rst)
}
