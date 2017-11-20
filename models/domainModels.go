package models

import (
	"errors"

	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/panda/logger"
	"github.com/hzwy23/panda/validator"
)

type DomainMmodel struct {
	ProjectId      string `json:"domain_id"`
	ProjectName    string `json:"domain_desc"`
	ProjectStatus  string `json:"domain_status"`
	CreateTime     string `json:"maintance_date" dateType:"YYYY-MM-DD HH24:MM:SS"`
	CreateUser     string `json:"create_user_id"`
	ModifyTime     string `json:"domain_modify_date" dateType:"YYYY-MM-DD HH24:MM:SS"`
	ModifyUser     string `json:"domain_modify_user"`
	DomainStatusCd string `json:"domain_status_cd"`
}

func (this *DomainMmodel) Get() ([]DomainMmodel, error) {
	var rst []DomainMmodel
	rows, err := dbobj.Query(sys_rdbms_025)
	defer rows.Close()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	//	var oneLine DomainMmodel
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return rst, nil
}

// 获取指定域的详细信息
func (this *DomainMmodel) GetRow(domain_id string) (DomainMmodel, error) {
	var rst DomainMmodel
	err := dbobj.QueryForStruct(sys_rdbms_084, &rst, domain_id)
	return rst, err
}

// 新增域信息
// 并将新增的域授权个创建人
func (this *DomainMmodel) Post(arg entity.DomainData) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		return "error_sql_begin", err
	}

	// validator domain id format
	if !validator.IsAlnum(arg.DomainId) {
		return "as_of_date_domain_id_check", errors.New("as_of_date_domain_id_check")
	}

	// validator domain describe format. It does not allow null values
	if validator.IsEmpty(arg.DomainDesc) {
		return "as_of_date_domain_isempty", errors.New("as_of_date_domain_isempty")
	}

	// validator domain status format
	// It must be in the 0 and 1
	if !validator.IsIn(arg.DomainStatus, "0", "1") {
		return "as_of_date_domain_status_check", errors.New("as_of_date_domain_status_check")
	}

	logger.Debug("新增域信息是：", arg)
	_, err = tx.Exec(sys_rdbms_036, arg.DomainId, arg.DomainDesc, arg.DomainStatus, arg.ModifyUser, arg.ModifyUser)
	if err != nil {
		tx.Rollback()
		return "as_of_date_domain_add_failed", err
	}

	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return "as_of_date_domain_add_failed", err
	}
	return "success", nil
}

// 删除域信息
// 在controller中校验权限
func (this *DomainMmodel) Delete(js []DomainMmodel) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return err
	}
	for _, val := range js {
		_, err := tx.Exec(sys_rdbms_037, val.ProjectId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// 更新域信息
// 只能更新名称和状态
func (this *DomainMmodel) Update(arg entity.DomainData) (string, error) {

	// 校验域名称,不能为空
	if validator.IsEmpty(arg.DomainDesc) {
		return "as_of_date_domain_isempty", errors.New("as_of_date_domain_isempty")
	}

	// 校验域状态编码,必须是0或者1
	if !validator.IsIn(arg.DomainStatus, "0", "1") {
		return "as_of_date_domain_status_check", errors.New("as_of_date_domain_status_check")
	}

	_, err := dbobj.Exec(sys_rdbms_038, arg.DomainDesc, arg.DomainStatus, arg.ModifyUser, arg.DomainId)
	if err != nil {
		logger.Error(err)
		return "as_of_date_domain_update", err
	}
	return "success", nil
}
