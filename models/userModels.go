package models

import (
	"errors"
	"net/url"
	"strings"

	"github.com/hzwy23/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/crypto/haes"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/validator"
)

type UserModel struct {
	morg OrgModel
}

// 查询用户自己的详细信息
func (UserModel) GetOwnerDetails(user_id string) ([]entity.UserInfo, error) {
	var rst []entity.UserInfo
	row, err := dbobj.Query(sys_rdbms_023, user_id)
	defer row.Close()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	err = dbobj.Scan(row, &rst)
	return rst, err
}

// 查询域中所有的用户信息
func (UserModel) GetDefault() ([]entity.UserInfo, error) {
	row, err := dbobj.Query(sys_rdbms_017)
	defer row.Close()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var rst []entity.UserInfo
	err = dbobj.Scan(row, &rst)
	return rst, err
}

// 新增用户信息
func (UserModel) Post(arg entity.UserInfo, userPasswd string) (string, error) {

	if !validator.IsWord(arg.UserId) {
		return "error_user_id_check", errors.New("error_user_id_check")
	}
	//

	if validator.IsEmpty(arg.UserName) {
		return "error_user_name_check", errors.New("error_user_name_check")
	}

	//
	if !validator.IsEmail(arg.UserEmail) {
		return "error_user_email_check", errors.New("error_user_email_check")
	}

	if validator.IsEmpty(arg.OrgUnitId) {
		return "error_user_role_org", errors.New("error_user_role_org")
	}

	//
	if !validator.IsMobilePhone(arg.UserPhone) {
		return "error_user_phone_check", errors.New("error_user_phone_check")
	}

	tx, err := dbobj.Begin()
	// insert user details
	//
	_, err = tx.Exec(sys_rdbms_018, arg.UserId, arg.UserName, arg.UserOwner, arg.UserEmail, arg.UserPhone, arg.OrgUnitId, arg.UserOwner)
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return "error_user_post", err
	}

	// insert user passwd
	_, err = tx.Exec(sys_rdbms_019, arg.UserId, userPasswd, arg.UserStatusId)
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return "error_user_post", err
	}

	// insert theme info
	_, err = tx.Exec(sys_rdbms_045, arg.UserId, "1005")
	if err != nil {
		tx.Rollback()
		logger.Error(err.Error())
		return "error_user_post", err
	}

	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return "error_user_post", err
	}
	return "success", nil
}

// 删除用户信息
func (UserModel) Delete(data []entity.UserInfo) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		return "error_sql_begin", err
	}

	for _, val := range data {
		_, err = tx.Exec(sys_rdbms_007, val.UserId, val.OrgUnitId)
		if err != nil {
			tx.Rollback()
			logger.Error(err)
			return "error_user_exec", err
		}
	}
	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return "error_user_commit", err
	}

	return "success", nil
}

// 搜索用户信息
func (this UserModel) Search(org_id string, status_id string) ([]entity.UserInfo, error) {
	var rst []entity.UserInfo

	ret, err := this.GetDefault()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if !validator.IsEmpty(org_id) {

		orglist, err := this.morg.GetSubOrgInfo(org_id)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		var orgmap map[string]string = make(map[string]string)
		for _, val := range orglist {
			orgmap[val.OrgUnitId] = ""
		}

		for _, val := range ret {
			if _, ok := orgmap[val.OrgUnitId]; ok {
				if !validator.IsEmpty(status_id) {
					if val.UserStatusId == status_id {
						rst = append(rst, val)
					}
				} else {
					rst = append(rst, val)
				}
			}
		}
	} else {
		for _, val := range ret {
			if !validator.IsEmpty(status_id) {
				if val.UserStatusId == status_id {
					rst = append(rst, val)
				}
			} else {
				rst = append(rst, val)
			}
		}
	}
	return rst, nil
}

func (this UserModel) ModifyStatus(status_id, user_id string) (string, error) {
	if !validator.IsIn(status_id, "0", "1") {
		return "error_user_status_empty", errors.New("error_user_status_empty")
	}

	_, err := dbobj.Exec(sys_rdbms_016, status_id, user_id)
	return "error_user_modify_status", err
}

func (this UserModel) ModifyPasswd(data url.Values) (string, error) {
	user_id := data.Get("userid")
	user_password := data.Get("newpasswd")
	confirm_password := data.Get("surepasswd")
	if user_password != confirm_password {
		return "error_passwd_confirm_failed", errors.New("error_passwd_confirm_failed")
	}

	if len(strings.TrimSpace(confirm_password)) < 6 || len(strings.TrimSpace(confirm_password)) > 30 {
		return "error_passwd_short", errors.New("error_passwd_short")
	}

	encry_passwd, err := haes.Encrypt(user_password)
	if err != nil {
		logger.Error(err)
		return "error_password_encrpty", errors.New("error_password_encrpty")
	}

	_, err = dbobj.Exec(sys_rdbms_020, encry_passwd, user_id)
	if err != nil {
		logger.Error(err)
		return "error_user_modify_passwd", err
	}
	return "success", nil
}

// 查询机构编码
func (this UserModel) GetOrgId(userId string) (string, error) {
	var orgId string
	err := dbobj.QueryForObject(sys_rdbms_010, dbobj.PackArgs(userId), &orgId)
	return orgId, err
}

// 修改用户信息
func (this UserModel) Put(data url.Values, modify_user string) (string, error) {
	user_name := data.Get("userDesc")
	org_id := data.Get("orgId")
	phone := data.Get("userPhone")
	email := data.Get("userEmail")
	user_id := data.Get("userId")

	if !validator.IsWord(user_id) {
		return "error_user_id_empty", errors.New("error_user_id_empty")
	}

	if validator.IsEmpty(user_name) {
		return "error_user_desc_empty", errors.New("error_user_desc_empty")
	}

	if !validator.IsEmail(email) {
		return "error_user_email_format", errors.New("error_user_email_format")
	}

	if !validator.IsWord(org_id) {
		return "error_org_id_format", errors.New("error_org_id_format")
	}

	if !validator.IsMobilePhone(phone) {
		return "error_user_phone_format", errors.New("error_user_phone_format")
	}

	_, err := dbobj.Exec(sys_rdbms_021, user_name, phone, email, modify_user, org_id, user_id)
	if err != nil {
		logger.Error(err)
		return "error_user_modify_info", err
	}
	return "success", nil
}
