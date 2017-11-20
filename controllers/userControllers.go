package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda"
	"github.com/hzwy23/panda/crypto/aes"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
	"github.com/hzwy23/panda/validator"
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
func (userController) Page(w http.ResponseWriter, r *http.Request) {

	rst, err := groupcache.GetStaticFile("AsofdasteUserPage")
	if err != nil {
		hret.Error(w, 404, i18n.PageNotFound(r))
		return
	}
	hz, err := service.ParseText(r, string(rst))
	if err != nil {
		hret.Error(w, 404, i18n.PageNotFound(r))
		return
	}
	hz.Execute(w, nil)
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
func (this userController) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	rst, err := this.models.GetDefault()
	if err != nil {
		logger.Error(err)
		hret.Error(w, 410, i18n.Get(r, "error_user_query"), err)
		return
	}
	hret.Json(w, rst)
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
func (this userController) Post(w http.ResponseWriter, r *http.Request) {
	var arg entity.UserInfo
	err := panda.ParseForm(r, &arg)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}
	arg.UserOwner = jclaim.UserId

	password := r.FormValue("user_passwd")
	surepassword := r.FormValue("user_passwd_confirm")

	if validator.IsEmpty(password) {
		hret.Error(w, 421, i18n.Get(r, "error_user_passwd_check"))
		return
	}

	if validator.IsEmpty(surepassword) {
		hret.Error(w, 421, i18n.Get(r, "error_passwd_empty"))
		return
	}

	if password != surepassword {
		hret.Error(w, 421, i18n.Get(r, "error_passwd_confirm_failed"))
		return
	}

	if len(strings.TrimSpace(password)) < 6 {
		hret.Error(w, 421, i18n.Get(r, "error_passwd_short"))
		return
	}

	userPasswd, err := aes.Encrypt(password)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_user_passwd_encrypt"))
		return
	}

	msg, err := this.models.Post(arg, userPasswd)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
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
func (this userController) Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var rst []entity.UserInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_user_json"))
		return
	}

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	for _, val := range rst {

		if panda.IsAdmin(val.UserId) {
			hret.Error(w, 403, i18n.Get(r, "error_user_forbid_delete_admin"))
			return
		}

		if val.UserId == jclaim.UserId {
			hret.Error(w, 403, i18n.Get(r, "error_user_forbid_yourself"))
			return
		}
	}

	msg, err := this.models.Delete(rst)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
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
func (this userController) Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var org_id = r.FormValue("org_id")
	var status_id = r.FormValue("status_id")

	logger.Debug("search user info,", org_id, status_id)

	rst, err := this.models.Search(org_id, status_id)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, "error_user_query"), err)
		return
	}
	hret.Json(w, rst)
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
func (this userController) Put(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	form := r.Form

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}
	msg, err := this.models.Put(form, jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
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
func (this userController) ModifyPasswd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	form := r.Form

	msg, err := this.models.ModifyPasswd(form)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))

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
func (this userController) ModifyStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user_id := r.FormValue("userId")
	status_id := r.FormValue("userStatus")

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	if jclaim.UserId == user_id {
		hret.Error(w, 403, i18n.Get(r, "error_user_modify_yourself"))
		return
	}

	msg, err := this.models.ModifyStatus(status_id, user_id)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
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
func (this userController) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 401, i18n.Disconnect(r))
		return
	}
	rst, err := this.models.GetOwnerDetails(jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, "error_user_query"))
		return
	}
	hret.Json(w, rst)
}

func init() {
	groupcache.RegisterStaticFile("AsofdasteUserPage", "./views/hauth/UserInfoPage.tpl")
}
