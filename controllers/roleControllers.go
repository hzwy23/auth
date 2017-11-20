package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
)

type roleController struct {
	models models.RoleModel
}

var RoleCtl = &roleController{
	models.RoleModel{},
}

// swagger:operation GET /v1/auth/role/page StaticFiles roleController
//
// 角色管理页面
//
// 如果用户被授权访问角色管理页面,则系统返回角色管理页面内容,否则返回404错误
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
func (roleController) Page(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !service.BasicAuth(r) {
		hret.Error(w, 403, i18n.NoAuth(r))
		return
	}

	rst, err := groupcache.GetStaticFile("AsofdateRolePage")
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

// swagger:operation GET /v1/auth/role/get roleController roleController
//
// 查询角色信息
//
// 查询指定域中的角色信息
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
func (this roleController) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	rst, err := this.models.Get()

	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_role_query"), err)
		return
	}

	hret.Json(w, rst)
}

// swagger:operation POST /v1/auth/role/post roleController roleController
//
// 新增角色信息
//
// 在某个指定的域中,新增角色信息
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
func (this roleController) Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var arg entity.RoleInfo
	err := panda.ParseForm(r, &arg)
	if err != nil {
		logger.Error(w, 423, err.Error())
		return
	}

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	arg.RoleOwner = jclaim.UserId
	msg, err := this.models.Post(arg)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
}

// swagger:operation POST /v1/auth/role/delete roleController roleController
//
// 删除角色信息
//
// 删除某个指定域中的角色信息
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
func (this roleController) Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !service.BasicAuth(r) {
		hret.Error(w, 403, i18n.NoAuth(r))
		return
	}

	var allrole []entity.RoleInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &allrole)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_role_json_failed"), err)
		return
	}

	msg, err := this.models.Delete(allrole)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 418, i18n.Get(r, msg))
		return
	}
	hret.Success(w, i18n.Success(r))
}

// swagger:operation PUT /v1/auth/role/put roleController roleController
//
// 更新角色信息
//
// 更新某个域中的角色信息,角色编码不能更新
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
func (this roleController) Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var arg entity.RoleInfo
	err := panda.ParseForm(r, &arg)
	if err != nil {
		logger.Error(w, 423, err.Error())
		return
	}

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}
	arg.RoleMaintanceUser = jclaim.UserId

	msg, err := this.models.Update(arg)
	if err != nil {
		logger.Error(err.Error())
		hret.Error(w, 421, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
}

func init() {
	groupcache.RegisterStaticFile("AsofdateRolePage", "./views/hauth/role_info_page.tpl")
}
