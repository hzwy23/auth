package models

import (
	"github.com/asofdate/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
)

type SysPrivilege struct {
}

// 获取权限定义信息
func (this *SysPrivilege) Get() ([]entity.SysPrivilegeEntity, error) {
	var rst []entity.SysPrivilegeEntity
	err := dbobj.QueryForSlice(sys_rdbms_055, &rst)
	return rst, err
}

// 查询某一个指定权限的相信信息
func (this *SysPrivilege) GetRow(privilegeId string) (entity.SysPrivilegeEntity, error) {
	var row entity.SysPrivilegeEntity
	err := dbobj.QueryForStruct(sys_rdbms_059, &row, privilegeId)
	return row, err
}

// 新建权限信息
func (this *SysPrivilege) Post(row entity.SysPrivilegeEntity) error {
	_, err := dbobj.Exec(sys_rdbms_056,
		row.PrivilegeId,
		row.PrivilegeDesc,
		row.CreateUser,
		row.CreateTime,
		row.ModifyUser,
		row.ModifyTime)
	return err
}

// 更新权限定义信息
func (this *SysPrivilege) Put(row entity.SysPrivilegeEntity) error {
	_, err := dbobj.Exec(sys_rdbms_057,
		row.PrivilegeDesc,
		row.ModifyUser,
		row.ModifyTime,
		row.PrivilegeId)
	return err
}

// 删除权限定义信息
func (this *SysPrivilege) Delete(rows []entity.SysPrivilegeEntity) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, val := range rows {
		_, err := tx.Exec(sys_rdbms_058, val.PrivilegeId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
