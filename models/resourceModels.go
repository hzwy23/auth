package models

import (
	"errors"

	"github.com/asofdate/auth-core/dto"
	"github.com/asofdate/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/validator"
)

type ResourceModel struct {
	Mtheme ThemeResourceModel
}

// 查询所有的资源信息
func (this *ResourceModel) Get() ([]entity.ResData, error) {
	rows, err := dbobj.Query(sys_rdbms_071)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var rst []entity.ResData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (this *ResourceModel) GetSubsystem(userId string) ([]dto.SubsystemDTO, error) {
	var rst []dto.SubsystemDTO
	err := dbobj.QueryForSlice(sys_rdbms_085, &rst, userId)
	return rst, err
}

func (this *ResourceModel) GetInnerFlag(resId string) (string, error) {
	innerFlag := "true"
	err := dbobj.QueryForObject(sys_rdbms_079, dbobj.PackArgs(resId), &innerFlag)
	return innerFlag, err
}

func (this *ResourceModel) GetServiceCd(resId string) (string, error) {
	serviceCd := ""
	err := dbobj.QueryForObject(sys_rdbms_048, dbobj.PackArgs(resId), &serviceCd)
	return serviceCd, err
}

func (this *ResourceModel) GetChildren(res_id string) ([]entity.ResData, error) {
	rst, err := this.Get()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var ret []entity.ResData
	this.dfs(rst, res_id, &ret)
	return ret, nil
}

func (this *ResourceModel) GetChildExceptButton(resId string) ([]entity.ResData, error) {
	var rst []entity.ResData
	err := dbobj.QueryForSlice(sys_rdbms_047, &rst)

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var ret []entity.ResData
	this.dfs(rst, resId, &ret)
	return ret, nil
}

// 所有指定资源的详细信息
func (this *ResourceModel) Query(res_id string) ([]entity.ResData, error) {
	rows, err := dbobj.Query(sys_rdbms_089, res_id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var rst []entity.ResData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 新增菜单资源
func (this *ResourceModel) Post(data entity.ResData) (string, error) {
	// 如果所属系统非空，表示是内部菜单
	innnerFlag := "false"
	if len(data.ServiceCd) == 0 {
		innnerFlag = "true"
	}

	// 1 表示叶子
	// 0 表示结点
	res_attr := "1"
	if data.Restype == "0" || data.Restype == "4" {
		res_attr = "0"
	}

	// 如果是首页子系统菜单，设置上级编码为-1
	if data.Restype == "0" {
		data.ResUpid = "-1"
	}

	if !validator.IsWord(data.ResId) {
		logger.Error("资源编码必须由1,30位字母或数字组成")
		return "error_resource_res_id", errors.New("error_resource_res_id")
	}

	if validator.IsEmpty(data.ResName) {
		logger.Error("菜单名称不能为空")
		return "error_resource_desc_empty", errors.New("error_resource_desc_empty")
	}

	if validator.IsEmpty(data.Restype) {
		logger.Error("菜单类别不能为空")
		return "error_resource_type", errors.New("error_resource_type")
	}

	if validator.IsEmpty(data.ResUpid) {
		logger.Error("菜单上级编码不能为空")
		return "error_resource_up_id", errors.New("error_resource_up_id")
	}

	// add sys_resource_info
	_, err := dbobj.Exec(sys_rdbms_072,
		data.ResId, data.ResName, res_attr, data.ResUpid,
		data.Restype, innnerFlag, data.ServiceCd)

	if err != nil {
		logger.Error(err)
		return "error_resource_add", err
	}
	return "success", nil
}

// 删除指定的资源
func (this *ResourceModel) Delete(res_id string) (string, error) {
	var rst []entity.ResData

	all, err := this.Get()
	if err != nil {
		logger.Error(err)
		return "error_resource_query", err
	}

	this.dfs(all, res_id, &rst)

	// add res_id
	for _, val := range all {
		if val.ResId == res_id {
			rst = append(rst, val)
			break
		}
	}

	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return "error_resource_begin", err
	}

	for _, val := range rst {

		if val.SysFlag == "0" {
			tx.Rollback()
			return "error_resource_forbid_system_resource", errors.New("error_resource_forbid_system_resource")
		}

		_, err = tx.Exec(sys_rdbms_075, val.ResId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return "error_resource_role_relation", err
		}

		_, err = tx.Exec(sys_rdbms_076, val.ResId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return "error_resource_theme_relation", err
		}

		_, err = tx.Exec(sys_rdbms_077, val.ResId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return "error_resource_delete", err
		}
	}
	return "error_resource_commit", tx.Commit()
}

func (this *ResourceModel) Update(arg entity.ResData) (string, error) {

	if validator.IsEmpty(arg.ResName) {
		return "error_resource_desc_empty", errors.New("error_resource_desc_empty")
	}

	if arg.ResId == arg.ResUpid {
		return "error_resource_update_same", errors.New("error_resource_update_same")
	}

	//获取当前菜单所有子菜单列表
	childList, err := this.GetChildren(arg.ResId)
	if err != nil {
		logger.Error(err)
		return "error_resource_update", errors.New("error_resource_update")
	}

	for _, val := range childList {
		if val.ResId == arg.ResUpid {
			return "error_resource_update", errors.New("error_resource_update")
		}
	}

	_, err = dbobj.Exec(sys_rdbms_005,
		arg.ResName,
		arg.ResUpid,
		arg.ServiceCd,
		arg.ResId)

	if err != nil {
		logger.Error(err)
		return "error_resource_update", err
	}
	return "success", nil
}

func (this *ResourceModel) GetNodes() ([]dto.ResNodeData, error) {
	var rst []dto.ResNodeData
	err := dbobj.QueryForSlice(sys_rdbms_046, &rst)
	return rst, err
}

// 获取子资源信息
func (this *ResourceModel) dfs(all []entity.ResData, res_id string, rst *[]entity.ResData) {
	for _, val := range all {
		if val.ResUpid == res_id {
			*rst = append(*rst, val)
			if val.ResId == val.ResUpid {
				logger.Error("层级关系错误,不允许上级菜单域当前菜单编码一致,当前菜单编码:", val.ResId, "上级菜单编码:", val.ResUpid)
				return
			}
			this.dfs(all, val.ResId, rst)
		}
	}
}
