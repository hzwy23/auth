package models

import (
	"github.com/hzwy23/auth/dto"
	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/dbobj"
)

type SysPrivilegeRole struct {
}

// 查询权限代码域域的关联关系
func (this *SysPrivilegeRole) Get(privilegeId string) ([]entity.PrivilegeRole, error) {
	var rst []entity.PrivilegeRole
	err := dbobj.QueryForSlice(sys_rdbms_061, &rst, privilegeId)
	return rst, err
}

func (this *SysPrivilegeRole) GetUnMapRole(privilegeId string) ([]dto.RoleDTO, error) {
	var rst []dto.RoleDTO
	err := dbobj.QueryForSlice(sys_rdbms_068, &rst, privilegeId)
	return rst, err
}

func (this *SysPrivilegeRole) Post(rows []entity.PrivilegeRole) error {
	tx, err := dbobj.Begin()
	if err != nil {
		return err
	}
	for _, row := range rows {
		_, err := tx.Exec(sys_rdbms_065,
			row.Uuid, row.PrivilegeId, row.RoleId, row.CreateUser, row.CreateTime, row.ModifyUser, row.ModifyTime)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (this *SysPrivilegeRole) Put(row entity.PrivilegeRole) error {
	_, err := dbobj.Exec(sys_rdbms_067,
		row.RoleId, row.ModifyUser, row.ModifyTime, row.Uuid)
	return err
}

func (this *SysPrivilegeRole) Delete(rows []entity.PrivilegeRole) error {
	tx, err := dbobj.Begin()
	if err != nil {
		return err
	}
	for _, val := range rows {
		_, err := tx.Exec(sys_rdbms_066, val.Uuid)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
