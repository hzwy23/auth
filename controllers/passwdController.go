package controllers

import (
	"net/http"
	"strings"

	"github.com/asofdate/auth-core/models"
	"github.com/hzwy23/utils/crypto/haes"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
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
func (this passwdController) PostModifyPasswd(ctx router.Context) {
	ctx.Request.ParseForm()

	oriPasswd := ctx.Request.FormValue("orapasswd")
	newPasswd := ctx.Request.FormValue("newpasswd")
	surePasswd := ctx.Request.FormValue("surepasswd")

	if oriPasswd == newPasswd {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_same"))
		return
	}

	if newPasswd != surePasswd {
		logger.Error("new passwd confirm failed. please check your new password and confirm password")
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_confirm_failed"))
		return
	}

	if len(strings.TrimSpace(newPasswd)) != len(newPasswd) {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_blank"))
		return
	}

	if len(strings.TrimSpace(newPasswd)) < 6 || len(strings.TrimSpace(newPasswd)) > 30 {
		logger.Error("新密码长度不能小于6位,且不能大于30位")
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_short"))
		return
	}

	oriEn, err := haes.Encrypt(oriPasswd)
	if err != nil {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_password_encrpty"))
		return
	}

	newPd, err := haes.Encrypt(newPasswd)
	if err != nil {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_password_encrpty"))
		return
	}
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	err_msg, err := this.p.UpdateMyPasswd(newPd, jclaim.UserId, oriEn)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, err_msg), err)
		return
	}
	http.SetCookie(ctx.ResponseWriter, &http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1})
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}
