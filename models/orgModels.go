package models

import (
	"errors"

	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/panda/logger"
	"github.com/hzwy23/panda/validator"
)

type OrgModel struct {
}

//获取域下边所有机构号
func (this *OrgModel) Get() ([]entity.SysOrgInfo, error) {
	var rst []entity.SysOrgInfo
	rows, err := dbobj.Query(sys_rdbms_041)
	if err != nil {
		return nil, err
	}

	err = dbobj.Scan(rows, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 查询某个机构的详细信息
func (this *OrgModel) GetDetails(orgUnitId string) (entity.SysOrgInfo, error) {
	var row entity.SysOrgInfo
	err := dbobj.QueryForStruct(sys_rdbms_054, &row, orgUnitId)
	return row, err
}

func (this *OrgModel) Delete(orgUnitId string) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return "error_sql_begin", errors.New("error_sql_begin")
	}

	// 获取这个机构的所有下属机构信息
	sublist, err := this.GetSubOrgInfo(orgUnitId)
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return "error_org_sub_query", errors.New("error_org_sub_query")
	}

	for _, org := range sublist {
		_, err := tx.Exec(sys_rdbms_044, org.OrgUnitId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return "error_org_delete", errors.New("error_org_delete")
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return "error_org_delete_commit", errors.New("error_org_delete_commit")
	}
	return "success", nil
}

func (this *OrgModel) Update(row entity.SysOrgInfo, user_id string) (string, error) {

	if !validator.IsWord(row.OrgUnitId) {
		return "error_org_id_format", errors.New("error_org_id_format")
	}

	// 校验输入信息
	if validator.IsEmpty(row.OrgUnitDesc) {
		return "error_org_id_desc_empty", errors.New("error_org_id_desc_empty")
	}

	if !validator.IsWord(row.UpOrgId) {
		return "error_org_up_id_empty", errors.New("error_org_up_id_empty")
	}

	check, err := this.GetSubOrgInfo(row.OrgUnitId)
	if err != nil {
		logger.Error(err)
		return "error_org_sub_query", errors.New("error_org_sub_query")
	}

	for _, val := range check {
		if val.OrgUnitId == row.UpOrgId {
			return "error_org_up_id_complex", errors.New("error_org_up_id_complex")
		}
	}

	_, err = dbobj.Exec(sys_rdbms_069, row.OrgUnitDesc, row.UpOrgId, user_id, row.OrgUnitId)
	if err != nil {
		logger.Error(err)
		return "error_org_modify", err
	}
	return "success", nil
}

func (this *OrgModel) Post(arg entity.SysOrgInfo, user_id string) (string, error) {

	if !validator.IsAlnum(arg.OrgUnitId) {
		return "error_org_id_format", errors.New("error_org_id_format")
	}

	if validator.IsEmpty(arg.OrgUnitDesc) {
		return "error_org_id_desc_empty", errors.New("error_org_id_desc_empty")
	}

	if !validator.IsWord(arg.UpOrgId) {
		return "error_org_up_id_empty", errors.New("error_org_up_id_empty")
	}

	_, err := dbobj.Exec(sys_rdbms_043, arg.OrgUnitDesc, arg.UpOrgId, user_id, user_id, arg.OrgUnitId)
	if err != nil {
		logger.Error(err)
		return "error_org_add", errors.New("error_org_add")
	}
	return "success", nil
}

func (this *OrgModel) GetSubOrgInfo(org_id string) ([]entity.SysOrgInfo, error) {
	var rst []entity.SysOrgInfo

	all, err := this.Get()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// 将自身机构加入到结果中.
	for _, val := range all {
		if val.OrgUnitId == org_id {
			rst = append(rst, val)
			break
		}
	}

	this.dfs(all, org_id, &rst)

	return rst, nil
}

func (this *OrgModel) dfs(node []entity.SysOrgInfo, org_id string, rst *[]entity.SysOrgInfo) {
	for _, val := range node {
		if val.UpOrgId == org_id {
			*rst = append(*rst, val)
			if val.OrgUnitId == val.UpOrgId {
				logger.Error("当前机构与上级机构编码一致,逻辑错误,退出递归")
				return
			}
			this.dfs(node, val.OrgUnitId, rst)
		}
	}
}

func (this *OrgModel) Upload(data []entity.SysOrgInfo) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return "error_sql_begin", errors.New("error_sql_begin")
	}

	for _, val := range data {
		if !validator.IsAlnum(val.OrgUnitId) {
			tx.Rollback()
			return "error_org_id_format", errors.New("机构编码必须由1-30位字母,数字组成")
		}

		if validator.IsEmpty(val.OrgUnitDesc) {
			tx.Rollback()
			return "error_org_id_desc_empty", errors.New("error_org_id_desc_empty")
		}

		if validator.IsEmpty(val.UpOrgId) {
			tx.Rollback()
			return "error_org_up_id_empty", errors.New("error_org_up_id_empty")
		}

		_, err = tx.Exec(sys_rdbms_043, val.OrgUnitDesc, val.UpOrgId, val.CreateUser, val.CreateUser, val.OrgUnitId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return "error_org_upload", errors.New("上传机构信息失败,机构号是:" + val.OrgUnitId + ",机构名称是:" + val.OrgUnitDesc)
		}
	}
	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return "error_org_submit", errors.New("error_org_submit")
	}
	return "success", nil
}
