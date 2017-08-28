package models

import (
	"errors"

	"github.com/asofdate/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
)

type PasswdModels struct {
}

func (r PasswdModels) UpdateMyPasswd(newPd, User_id, oriEn string) (string, error) {
	flag, _, _, _ := r.CheckPasswd(User_id, oriEn)
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

func (r PasswdModels) GetDetails(userId string) (entity.SysUserSec, error) {
	var sec entity.SysUserSec
	err := dbobj.QueryForStruct(sys_rdbms_002, &sec, userId)
	return sec, err
}

func (r PasswdModels) UpdateContinueErrorCnt(cnt int, user_id string) {
	dbobj.Exec(sys_rdbms_004, cnt, user_id)
}

func (r PasswdModels) ForbidUsers(user_id string) {
	dbobj.Exec(sys_rdbms_022, user_id)
}

// check user's passwd is right.
func (r PasswdModels) CheckPasswd(user_id, user_passwd string) (bool, int, int, string) {
	var sec entity.SysUserSec
	sec, err := r.GetDetails(user_id)
	if err != nil {
		return false, 402, 0, "用户不存在"
	}

	if sec.UserStatus != 0 {
		return false, 406, sec.UserStatus, "用户状态被锁定，请联系管理员解锁"
	}

	if sec.ErrorCnt > 6 {
		r.ForbidUsers(user_id)
		return false, 403, sec.ErrorCnt, "用户已被锁定，请联系管理员解锁"
	}

	if sec.UserId == user_id && sec.UserPasswd == user_passwd {
		r.UpdateContinueErrorCnt(0, user_id)
		return true, 200, 0, "success"
	} else {
		r.UpdateContinueErrorCnt(sec.ErrorCnt+1, user_id)
		return false, 405, sec.ErrorCnt + 1, "用户密码错误"
	}
}
