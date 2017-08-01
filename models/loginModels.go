package models

import (
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/hzwy23/dbobj"
)

type LoginModels struct {
}

func (this LoginModels) GetDefaultPage(user_id string) string {
	row := dbobj.QueryRow(sys_rdbms_078, user_id)
	var url = "./views/hauth/theme/default/index.tpl"
	err := row.Scan(&url)
	if err != nil {
		logger.Debug("get default theme.")
		url = "./views/hauth/theme/default/index.tpl"
	}
	return url
}

func (this LoginModels) GetDefaultOrgId(user_id string) (org_id string, err error) {
	err = dbobj.QueryRow(sys_rdbms_080, user_id).Scan(&org_id)
	return
}
