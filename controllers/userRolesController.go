package controllers

import (
	"encoding/json"

	"github.com/hzwy23/auth-core/groupcache"
	"github.com/hzwy23/auth-core/models"
	"github.com/hzwy23/auth-core/service"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
)

type userRolesController struct {
	models    *models.UserRolesModel
	roleModel *models.RoleModel
}

var UserRolesCtl = &userRolesController{
	models:    new(models.UserRolesModel),
	roleModel: new(models.RoleModel),
}

// swagger:operation GET /v1/auth/batch/page StaticFiles domainShareControll
//
// If the request is successful,
// will return authorization page information to the client
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// If the user is authorized to visit, the return authorization information page
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: request success.
//   '404':
//     description: page not found.
func (this *userRolesController) Page(ctx router.Context) {
	// According to the key get the value from the groupCache system
	rst, err := groupcache.GetStaticFile("AuthorityPage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
		return
	}

	hz, err := service.ParseText(ctx, string(rst))
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
		return
	}
	hz.Execute(ctx.ResponseWriter, nil)
}

func (this *userRolesController) UserPage(ctx router.Context) {
	// According to the key get the value from the groupCache system

	ctx.Request.ParseForm()

	var role_id = ctx.Request.FormValue("role_id")

	rst, err := this.roleModel.GetRow(role_id)

	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_role_resource_query"))
		return
	}
	file, _ := service.ParseFile(ctx, "./views/hauth/role_user.tpl")

	file.Execute(ctx.ResponseWriter, rst)

}

// swagger:operation GET /v1/auth/user/roles/get userRolesController userRolesController
//
// 通过user_id用户账号，来查询这个用户拥有的角色信息
//
// 查询角色信息
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
func (this userRolesController) GetRolesByUserId(ctx router.Context) {
	ctx.Request.ParseForm()
	user_id := ctx.Request.FormValue("user_id")

	rst, err := this.models.GetRolesByUser(user_id)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_role_query"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/user/roles/other userRolesController userRolesController
//
// 通过user_id账号，查询这个用户能够访问，但是又没有获取到的角色信息
//
// 查询用户没有获取的角色
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
//     description: all domain information
func (this userRolesController) GetOtherRoles(ctx router.Context) {
	ctx.Request.ParseForm()
	user_id := ctx.Request.FormValue("user_id")

	if len(user_id) == 0 {
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_role_no_user"))
		return
	}

	rst, err := this.models.GetOtherRoles(user_id)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_role_un_auth"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/user/roles/auth userRolesController userRolesController
//
// 给指定的用户授予角色
//
// 给指定的用户授予角色
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
func (this userRolesController) Auth(ctx router.Context) {
	ctx.Request.ParseForm()
	var rst []models.UserRolesModel
	err := json.Unmarshal([]byte(ctx.Request.FormValue("JSON")), &rst)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_unmarsh_json"), err)
		return
	}

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Auth(rst, jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation POST /v1/auth/user/roles/revoke userRolesController userRolesController
//
// Delete user has been granted the roles
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// If the user is authorized to visit, the system will delete the roles that client request specified.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: user_id
//   in: query
//   description: Removed the role of the user
//   required: true
//   type: string
//   format:
// - name: role_id
//   in: query
//   description: The role of ready to delete
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userRolesController) Revoke(ctx router.Context) {
	ctx.Request.ParseForm()
	form := ctx.Request.FormValue("JSON")
	var rst []models.UserRolesModel

	err := json.Unmarshal([]byte(form), &rst)
	if err != nil {
		logger.Error("解析json格式数据失败，请联系管理员")
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_unmarsh_json"))
		return
	}

	msg, err := this.models.Revoke(rst)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

func (this userRolesController) GetOtherUserListByRoleId(ctx router.Context) {
	ctx.Request.ParseForm()
	roleId := ctx.Request.FormValue("roleId")
	if len(roleId) == 0 {
		logger.Error("role id is empty")
		hret.Error(ctx.ResponseWriter, 423, "查询失败，角色编码为空")
		return
	}
	rst, err := this.models.GetOtherUsersByRoleId(roleId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

func (this userRolesController) GetUserListByRoleId(ctx router.Context) {
	ctx.Request.ParseForm()

	roleId := ctx.Request.FormValue("roleId")
	if len(roleId) == 0 {
		logger.Error("role id is empty")
		hret.Error(ctx.ResponseWriter, 423, "查询失败，角色编码为空")
		return
	}
	rst, err := this.models.GetRelationUsersByRoleId(roleId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

func init() {
	// Registered in the static page to the groupCache system
	groupcache.RegisterStaticFile("AuthorityPage", "./views/hauth/sys_batch_page.tpl")
}
