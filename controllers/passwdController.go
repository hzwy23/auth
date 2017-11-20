package controllers

import (
	"net/http"
	"strings"

	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/panda/crypto/aes"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
)

type passwdController struct {
	p *models.PasswdModels
}

var PasswdController = &passwdController{
	p: &models.PasswdModels{},
}

// swagger:operation POST /v1/auth/passwd/update passwdController passwdController
//
// 修改用户自己的密码信息
//
// API提供了修改用户自己密码的服务,这个服务,不能删除其他用户的密码
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: orapasswd
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: newpasswd
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: surepasswd
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: all domain information
func (this passwdController) PostModifyPasswd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	oriPasswd := r.FormValue("orapasswd")
	newPasswd := r.FormValue("newpasswd")
	surePasswd := r.FormValue("surepasswd")

	if oriPasswd == newPasswd {
		hret.Error(w, 421, i18n.Get(r, "error_passwd_same"))
		return
	}

	if newPasswd != surePasswd {
		logger.Error("new passwd confirm failed. please check your new password and confirm password")
		hret.Error(w, 421, i18n.Get(r, "error_passwd_confirm_failed"))
		return
	}

	if len(strings.TrimSpace(newPasswd)) != len(newPasswd) {
		hret.Error(w, 421, i18n.Get(r, "error_passwd_blank"))
		return
	}

	if len(strings.TrimSpace(newPasswd)) < 6 || len(strings.TrimSpace(newPasswd)) > 30 {
		logger.Error("新密码长度不能小于6位,且不能大于30位")
		hret.Error(w, 421, i18n.Get(r, "error_passwd_short"))
		return
	}

	oriEn, err := aes.Encrypt(oriPasswd)
	if err != nil {
		hret.Error(w, 421, i18n.Get(r, "error_password_encrpty"))
		return
	}

	newPd, err := aes.Encrypt(newPasswd)
	if err != nil {
		hret.Error(w, 421, i18n.Get(r, "error_password_encrpty"))
		return
	}

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	err_msg, err := this.p.UpdateMyPasswd(newPd, jclaim.UserId, oriEn)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, err_msg), err)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1})
	hret.Success(w, i18n.Success(r))
}
