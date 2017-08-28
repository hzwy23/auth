package models

import (
	"errors"

	"github.com/asofdate/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/validator"
)

type RoleModel struct {
}

// 查询某一个角色的具体信息
func (this RoleModel) GetRow(role_id string) (entity.RoleInfo, error) {
	var rst entity.RoleInfo
	ret, err := this.Get()
	if err != nil {
		logger.Error(err)
		return rst, err
	}

	for _, val := range ret {
		if val.RoleId == role_id {
			return val, nil
		}
	}
	return rst, errors.New("no found")
}

func (RoleModel) Get() ([]entity.RoleInfo, error) {
	rows, err := dbobj.Query(sys_rdbms_028)
	defer rows.Close()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var rst []entity.RoleInfo
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (RoleModel) Post(arg entity.RoleInfo) (string, error) {
	//校验
	if !validator.IsAlnum(arg.RoleId) {
		return "error_role_id_format", errors.New("error_role_id_format")
	}
	//
	if validator.IsEmpty(arg.RoleName) {
		return "error_role_desc_empty", errors.New("error_role_desc_empty")
	}

	if !validator.IsIn(arg.RoleStatus, "0", "1") {
		return "error_role_status", errors.New("error_role_status")
	}

	_, err := dbobj.Exec(sys_rdbms_026, arg.RoleId, arg.RoleName, arg.RoleOwner, arg.RoleStatus, arg.RoleOwner)
	if err != nil {
		logger.Error(err)
		return "error_role_add_failed", err
	}
	return "success", nil
}

func (RoleModel) Delete(allrole []entity.RoleInfo) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return "error_sql_begin", err
	}

	for _, val := range allrole {
		_, err := tx.Exec(sys_rdbms_027, val.RoleId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return "error_role_delete_failed", err
		}
		logger.Info("delete role info successfully. role id is :", val.RoleId)
	}
	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return "error_role_delete_failed", err
	}
	return "success", nil
}

func (RoleModel) Update(arg entity.RoleInfo) (string, error) {

	if !validator.IsWord(arg.RoleId) {
		return "error_role_id_format", errors.New("error_role_id_format")
	}

	if validator.IsEmpty(arg.RoleName) {
		return "error_role_desc_empty", errors.New("error_role_desc_empty")
	}

	if !validator.IsIn(arg.RoleStatus, "0", "1") {
		return "error_role_status", errors.New("error_role_status")
	}

	_, err := dbobj.Exec(sys_rdbms_050, arg.RoleName, arg.RoleStatus, arg.RoleMaintanceUser, arg.RoleId)
	if err != nil {
		logger.Error(err)
		return "error_role_update_failed", errors.New("error_role_update_failed")
	}
	return "success", nil
}
