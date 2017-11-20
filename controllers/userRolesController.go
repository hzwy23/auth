package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
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
func (this *userRolesController) Page(w http.ResponseWriter, r *http.Request) {
	// According to the key get the value from the groupCache system
	rst, err := groupcache.GetStaticFile("AuthorityPage")
	if err != nil {
		hret.Error(w, 404, i18n.Get(r, "as_of_date_page_not_exist"))
		return
	}

	hz, err := service.ParseText(r, string(rst))
	if err != nil {
		hret.Error(w, 404, i18n.Get(r, "as_of_date_page_not_exist"))
		return
	}
	hz.Execute(w, nil)
}

func (this *userRolesController) UserPage(w http.ResponseWriter, r *http.Request) {
	// According to the key get the value from the groupCache system

	r.ParseForm()

	var role_id = r.FormValue("role_id")

	rst, err := this.roleModel.GetRow(role_id)

	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, "error_role_resource_query"))
		return
	}
	file, _ := service.ParseFile(r, "./views/hauth/role_user.tpl")

	file.Execute(w, rst)

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
func (this userRolesController) GetRolesByUserId(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user_id := r.FormValue("user_id")

	rst, err := this.models.GetRolesByUser(user_id)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, "error_user_role_query"), err)
		return
	}
	hret.Json(w, rst)
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
func (this userRolesController) GetOtherRoles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user_id := r.FormValue("user_id")

	if len(user_id) == 0 {
		hret.Error(w, 419, i18n.Get(r, "error_user_role_no_user"))
		return
	}

	rst, err := this.models.GetOtherRoles(user_id)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, "error_user_role_un_auth"), err)
		return
	}
	hret.Json(w, rst)
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
func (this userRolesController) Auth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var rst []models.UserRolesModel
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_unmarsh_json"), err)
		return
	}

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	msg, err := this.models.Auth(rst, jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
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
func (this userRolesController) Revoke(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.FormValue("JSON")
	var rst []models.UserRolesModel

	err := json.Unmarshal([]byte(form), &rst)
	if err != nil {
		logger.Error("解析json格式数据失败，请联系管理员")
		hret.Error(w, 421, i18n.Get(r, "error_unmarsh_json"))
		return
	}

	msg, err := this.models.Revoke(rst)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, msg), err)
		return
	}
	hret.Success(w, i18n.Success(r))
}

func (this userRolesController) GetOtherUserListByRoleId(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	roleId := r.FormValue("roleId")
	if len(roleId) == 0 {
		logger.Error("role id is empty")
		hret.Error(w, 423, "查询失败，角色编码为空")
		return
	}
	rst, err := this.models.GetOtherUsersByRoleId(roleId)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	hret.Json(w, rst)
}

func (this userRolesController) GetUserListByRoleId(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	roleId := r.FormValue("roleId")
	if len(roleId) == 0 {
		logger.Error("role id is empty")
		hret.Error(w, 423, "查询失败，角色编码为空")
		return
	}
	rst, err := this.models.GetRelationUsersByRoleId(roleId)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	hret.Json(w, rst)
}

func init() {
	// Registered in the static page to the groupCache system
	groupcache.RegisterStaticFile("AuthorityPage", "./views/hauth/sys_batch_page.tpl")
}
