package controllers

import (
	"encoding/json"

	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/groupcache"
	"github.com/asofdate/auth-core/models"
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
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
func (roleController) Page(ctx router.Context) {
	ctx.Request.ParseForm()
	if !service.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := groupcache.GetStaticFile("AsofdateRolePage")
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
func (this roleController) Get(ctx router.Context) {
	ctx.Request.ParseForm()

	rst, err := this.models.Get()

	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_role_query"), err)
		return
	}

	hret.Json(ctx.ResponseWriter, rst)
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
func (this roleController) Post(ctx router.Context) {
	ctx.Request.ParseForm()

	var arg entity.RoleInfo
	err := utils.ParseForm(ctx.Request, &arg)
	if err != nil {
		logger.Error(ctx.ResponseWriter, 423, err.Error())
		return
	}

	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	arg.RoleOwner = jclaim.UserId
	msg, err := this.models.Post(arg)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
func (this roleController) Delete(ctx router.Context) {
	ctx.Request.ParseForm()
	if !service.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	var allrole []entity.RoleInfo
	err := json.Unmarshal([]byte(ctx.Request.FormValue("JSON")), &allrole)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_role_json_failed"), err)
		return
	}

	msg, err := this.models.Delete(allrole)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 418, i18n.Get(ctx.Request, msg))
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
func (this roleController) Update(ctx router.Context) {
	ctx.Request.ParseForm()
	var arg entity.RoleInfo
	err := utils.ParseForm(ctx.Request, &arg)
	if err != nil {
		logger.Error(ctx.ResponseWriter, 423, err.Error())
		return
	}

	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}
	arg.RoleMaintanceUser = jclaim.UserId

	msg, err := this.models.Update(arg)
	if err != nil {
		logger.Error(err.Error())
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

func init() {
	groupcache.RegisterStaticFile("AsofdateRolePage", "./views/hauth/role_info_page.tpl")
}
