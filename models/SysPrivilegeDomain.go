package models

import (
	"github.com/hzwy23/auth-core/dto"
	"github.com/hzwy23/auth-core/entity"
	"github.com/hzwy23/dbobj"
)

type SysPrivilegeDomain struct {
}

// 查询权限代码域域的关联关系
func (this *SysPrivilegeDomain) Get(privilegeId string) ([]entity.PrivilegeDomain, error) {
	var rst []entity.PrivilegeDomain
	err := dbobj.QueryForSlice(sys_rdbms_060, &rst, privilegeId)
	return rst, err
}

func (this *SysPrivilegeDomain) GetUmmapDomain(privilegeId string) ([]dto.AccessDomainDTO, error) {
	var rst []dto.AccessDomainDTO
	err := dbobj.QueryForSlice(sys_rdbms_107, &rst, privilegeId)
	return rst, err
}

// 新增权限与域的关联关系
func (this *SysPrivilegeDomain) Post(rows []entity.PrivilegeDomain) error {
	tx, err := dbobj.Begin()
	if err != nil {
		return err
	}
	for _, row := range rows {
		_, err := dbobj.Exec(sys_rdbms_062,
			row.Uuid, row.PrivilegeId,
			row.DomainId, row.Permission,
			row.CreateUser, row.CreateTime,
			row.ModifyUser, row.ModifyTime)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// 更新权限与域的关联关系
func (this *SysPrivilegeDomain) Put(row entity.PrivilegeDomain) error {
	_, err := dbobj.Exec(sys_rdbms_064,
		row.DomainId, row.Permission,
		row.ModifyUser, row.ModifyTime, row.Uuid)
	return err
}

// 删除权限与域的关联关系
func (this *SysPrivilegeDomain) Delete(rows []entity.PrivilegeDomain) error {
	tx, err := dbobj.Begin()
	if err != nil {
		return err
	}
	for _, val := range rows {
		_, err := tx.Exec(sys_rdbms_063, val.Uuid)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
