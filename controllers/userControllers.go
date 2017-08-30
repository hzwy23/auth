package controllers

import (
	"encoding/json"
	"strings"

	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/groupcache"
	"github.com/asofdate/auth-core/models"
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/crypto/haes"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
	"github.com/hzwy23/utils/validator"
)

type userController struct {
	models models.UserModel
}

var UserCtl = &userController{}

// swagger:operation GET /v1/auth/user/page StaticFiles userController
//
// 获取用户管理子页面
//
// 返回用户管理页面
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: success
func (userController) Page(ctx router.Context) {

	rst, err := groupcache.GetStaticFile("AsofdasteUserPage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}
	hz, err := service.ParseText(ctx, string(rst))
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}
	hz.Execute(ctx.ResponseWriter, nil)
}

// swagger:operation GET /v1/auth/user/get userController userController
//
// 获取指定域中用户信息
//
// 返回指定域中的用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Get(ctx router.Context) {
	ctx.Request.ParseForm()

	rst, err := this.models.GetDefault()
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 410, i18n.Get(ctx.Request, "error_user_query"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/user/post userController userController
//
// 新增用户信息
//
// 在某个域中新增用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Post(ctx router.Context) {
	var arg entity.UserInfo
	err := utils.ParseForm(ctx.Request, &arg)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 423, err.Error())
		return
	}

	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}
	arg.UserOwner = jclaim.UserId

	password := ctx.Request.FormValue("user_passwd")
	surepassword := ctx.Request.FormValue("user_passwd_confirm")

	if validator.IsEmpty(password) {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_user_passwd_check"))
		return
	}

	if validator.IsEmpty(surepassword) {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_empty"))
		return
	}

	if password != surepassword {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_confirm_failed"))
		return
	}

	if len(strings.TrimSpace(password)) < 6 {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_short"))
		return
	}

	userPasswd, err := haes.Encrypt(password)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_user_passwd_encrypt"))
		return
	}

	msg, err := this.models.Post(arg, userPasswd)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation POST /v1/auth/user/delete userController userController
//
// 删除用户信息
//
// 删除某个域中的用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: success
func (this userController) Delete(ctx router.Context) {
	ctx.Request.ParseForm()
	var rst []entity.UserInfo
	err := json.Unmarshal([]byte(ctx.Request.FormValue("JSON")), &rst)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_user_json"))
		return
	}

	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	for _, val := range rst {

		if utils.IsAdmin(val.UserId) {
			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_forbid_delete_admin"))
			return
		}

		if val.UserId == jclaim.UserId {
			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_forbid_yourself"))
			return
		}
	}

	msg, err := this.models.Delete(rst)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation GET /v1/auth/user/search userController userController
//
// 搜索用户信息
//
// 客户端发起请求时,必须带如下几个参数;
//
// org_id 机构编码
// status_id 机构状态
// domain_id 所属域
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: org_id
//   in: query
//   description: org code number
//   required: true
//   type: string
//   format:
// - name: status_id
//   in: query
//   description: status code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Search(ctx router.Context) {
	ctx.Request.ParseForm()
	var org_id = ctx.Request.FormValue("org_id")
	var status_id = ctx.Request.FormValue("status_id")

	logger.Debug("search user info,", org_id, status_id)

	rst, err := this.models.Search(org_id, status_id)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_query"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation PUT /v1/auth/user/put userController userController
//
// 修改用户信息
//
// 修改用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Put(ctx router.Context) {
	ctx.Request.ParseForm()

	form := ctx.Request.Form

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}
	msg, err := this.models.Put(form, jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation PUT /v1/auth/user/modify/passwd userController userController
//
// 修改用户密码
//
// 修改用户密码
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) ModifyPasswd(ctx router.Context) {
	ctx.Request.ParseForm()

	form := ctx.Request.Form

	msg, err := this.models.ModifyPasswd(form)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))

}

// swagger:operation PUT /v1/auth/user/modify/status userController userController
//
// 修改用户锁状态
//
// 修改用户锁状态
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) ModifyStatus(ctx router.Context) {
	ctx.Request.ParseForm()
	user_id := ctx.Request.FormValue("userId")
	status_id := ctx.Request.FormValue("userStatus")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if jclaim.UserId == user_id {
		hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_modify_yourself"))
		return
	}

	msg, err := this.models.ModifyStatus(status_id, user_id)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation GET /v1/auth/user/query userController userController
//
// 查询用户自身信息
//
// 查询用户自身信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) GetUserDetails(ctx router.Context) {
	ctx.Request.ParseForm()
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 401, i18n.Disconnect(ctx.Request))
		return
	}
	rst, err := this.models.GetOwnerDetails(jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_query"))
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

func init() {
	groupcache.RegisterStaticFile("AsofdasteUserPage", "./views/hauth/UserInfoPage.tpl")
}
