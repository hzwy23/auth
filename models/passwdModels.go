package models

import (
	"errors"

	"github.com/asofdate/sso-jwt-auth/hrpc"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/hzwy23/dbobj"
)

type PasswdModels struct {
}

func (r PasswdModels) UpdateMyPasswd(newPd, User_id, oriEn string) (string, error) {
	flag, _, _, _ := hrpc.CheckPasswd(User_id, oriEn)
	if !flag {
		return "error_old_passwd", errors.New("error_old_passwd")
	}
	_, err := dbobj.Exec(sys_rdbms_014, newPd, User_id, oriEn)
	if err != nil {
		logger.Error(err)
		return "error_passwd_modify", err
	}
	return "success", nil
}

func (r PasswdModels) UpdateUserPasswd(newPd, userid string) error {
	_, err := dbobj.Exec(sys_rdbms_015, newPd, userid)
	return err
}
