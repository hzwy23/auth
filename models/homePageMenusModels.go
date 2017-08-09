package models

import (
	"encoding/json"

	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/hzwy23/dbobj"
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

type homePageMenuData struct {
	Res_id        string
	Res_name      string
	Res_url       string
	Res_bg_color  string
	Res_class     string
	Res_img       string
	Group_id      string
	Res_up_id     string
	Res_open_type string
	New_iframe    string
	Inner_flag    string
	Res_attr      string
}

func (this HomePageMenusModel) Get(id, typeId, useId string) ([]byte, error) {

	// 首先获取用户主题信息
	theme, err := this.mut.Get(useId)
	if err != nil || len(theme) != 1 {
		logger.Error(err)
		return nil, err
	}

	// 获取这个主题的所有资源信息
	theme_resource, err := this.mts.Get(theme[0].Theme_id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var mres = make(map[string]ResData)

	// 如果是超级管理员用户，
	// 获取系统中所有的菜单信息
	if useId == "admin" {
		resdata, err := this.ms.GetChildren(id)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range resdata {
			if val.Res_type == typeId {
				mres[val.Res_id] = val
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
			mres[val.Res_id] = val
		}
	}

	// 获取角色拥有的资源信息
	var rst []homePageMenuData
	for _, t_res := range theme_resource {
		if val, ok := mres[t_res.Res_id]; ok {
			var one homePageMenuData
			one.Res_id = t_res.Res_id
			one.Res_up_id = val.Res_up_id
			one.Res_name = val.Res_name
			one.Group_id = t_res.Group_id
			one.Res_bg_color = t_res.Res_bg_color
			one.Res_class = t_res.Res_class
			one.Res_img = t_res.Res_img
			one.Res_url = t_res.Res_url
			one.Res_open_type = t_res.Res_type
			one.New_iframe = t_res.New_iframe
			one.Inner_flag = val.Inner_flag
			rst = append(rst, one)
		}
	}
	return json.Marshal(rst)
}

func (this HomePageMenusModel) GetUrl(user_id, id string) (string, error) {
	row := dbobj.QueryRow(sys_rdbms_011, user_id, id)
	var url string
	err := row.Scan(&url)
	if err != nil {
		logger.Error("Get Special Page failure, user_id is :", user_id, "menu id is :", id, "details error is:", err)
		return redirect, err
	}
	return url, nil
}

func (this HomePageMenusModel) GetAllMenusExceptButton(userId string) ([]byte, error) {
	// 首先获取用户主题信息
	theme, err := this.mut.Get(userId)
	if err != nil || len(theme) != 1 {
		logger.Error(err)
		return nil, err
	}

	// 获取这个主题的所有资源信息
	theme_resource, err := this.mts.Get(theme[0].Theme_id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var mres = make(map[string]ResData)

	// 如果是超级管理员用户，
	// 获取系统中所有的菜单信息
	if userId == "admin" {
		// 获取所有的下级资源
		resdata, err := this.ms.GetChildren("-1")
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range resdata {
			if val.Res_type != "2" {
				mres[val.Res_id] = val
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
		role_resource, err := this.mrs.Gets(role_list, "-1", "0")
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range role_resource {
			mres[val.Res_id] = val
		}
	}

	themeResourceMap := make(map[string]themeData)
	for _, t_res := range theme_resource {
		themeResourceMap[t_res.Res_id] = t_res
	}

	// 获取角色拥有的资源信息
	var rst []homePageMenuData
	for _, val := range mres {
		if t_res, ok := themeResourceMap[val.Res_id]; ok {
			var one homePageMenuData
			one.Res_id = t_res.Res_id
			one.Res_up_id = val.Res_up_id
			one.Res_name = val.Res_name
			one.Group_id = t_res.Group_id
			one.Res_bg_color = t_res.Res_bg_color
			one.Res_class = t_res.Res_class
			one.Res_img = t_res.Res_img
			one.Res_url = t_res.Res_url
			one.Res_open_type = t_res.Res_type
			one.New_iframe = t_res.New_iframe
			one.Inner_flag = val.Inner_flag
			one.Res_attr = val.Res_attr
			rst = append(rst, one)
		} else {
			var one homePageMenuData
			one.Res_id = val.Res_id
			one.Res_up_id = val.Res_up_id
			one.Res_name = val.Res_name
			one.Inner_flag = val.Inner_flag
			one.Res_attr = val.Res_attr
			rst = append(rst, one)
		}
	}
	return json.Marshal(rst)
}
