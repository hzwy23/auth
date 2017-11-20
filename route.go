package auth

import (
	"github.com/hzwy23/auth/controllers"
	"github.com/hzwy23/panda/route"
)

func Register() {

	route.Handler("GET", "/v1/auth/HomePage", controllers.HomePage)
	route.Handler("POST", "/v1/auth/login", controllers.LoginSystem)
	route.Handler("GET", "/v1/auth/logout", controllers.LogoutSystem)
	route.Handler("GET", "/", controllers.IndexPage)

	route.Handler("POST", "/v1/auth/theme/update", controllers.ThemeCtl.Post)
	route.Handler("PUT", "/v1/auth/resource/config/theme", controllers.ThemeCtl.Put)
	route.Handler("GET", "/v1/auth/resource/queryTheme", controllers.ThemeCtl.QueryTheme)

	route.Handler("GET", "/v1/auth/index/entry", controllers.SubSystemEntry)
	route.Handler("GET", "/v1/auth/main/menu", controllers.IndexMenus)
	route.Handler("GET", "/v1/auth/menu/all/except/button", controllers.AllMenusExceptButton)
	route.Handler("POST", "/v1/auth/passwd/update", controllers.PasswdController.PostModifyPasswd)

	//domain_info
	route.Handler("GET", "/v1/auth/domain/get", controllers.DomainCtl.Get)
	route.Handler("POST", "/v1/auth/domain/post", controllers.DomainCtl.Post)
	route.Handler("POST", "/v1/auth/domain/delete", controllers.DomainCtl.Delete)
	route.Handler("PUT", "/v1/auth/domain/update", controllers.DomainCtl.Put)
	route.Handler("GET", "/v1/auth/domain/details", controllers.DomainCtl.GetDetails)

	//handle_logs
	route.Handler("GET", "/v1/auth/handle/logs/search", controllers.HandleLogsCtl.SerachLogs)
	route.Handler("GET", "/v1/auth/handle/logs", controllers.HandleLogsCtl.GetHandleLogs)
	route.Handler("GET", "/v1/auth/handle/logs/download", controllers.HandleLogsCtl.Download)

	//org_info
	route.Handler("GET", "/v1/auth/org/get", controllers.OrgCtl.Get)
	route.Handler("POST", "/v1/auth/org/insert", controllers.OrgCtl.Post)
	route.Handler("PUT", "/v1/auth/org/update", controllers.OrgCtl.Update)
	route.Handler("POST", "/v1/auth/org/delete", controllers.OrgCtl.Delete)
	route.Handler("GET", "/v1/auth/org/download", controllers.OrgCtl.Download)
	route.Handler("POST", "/v1/auth/org/upload", controllers.OrgCtl.Upload)
	route.Handler("GET", "/v1/auth/org/sub", controllers.OrgCtl.GetSubOrgInfo)
	route.Handler("GET", "/v1/auth/org/details", controllers.OrgCtl.GetDetails)

	//resource_info
	route.Handler("POST", "/v1/auth/resource/delete", controllers.ResourceCtl.Delete)
	route.Handler("POST", "/v1/auth/resource/post", controllers.ResourceCtl.Post)
	route.Handler("PUT", "/v1/auth/resource/update", controllers.ResourceCtl.Update)
	route.Handler("GET", "/v1/auth/resource/get", controllers.ResourceCtl.Get)
	route.Handler("GET", "/v1/auth/resource/query", controllers.ResourceCtl.Query)
	route.Handler("GET", "/v1/auth/resource/node", controllers.ResourceCtl.GetNodes)
	route.Handler("GET", "/v1/auth/resource/subsystem", controllers.ResourceCtl.GetSubsystemList)

	// 功能服务路由注册
	// 主要完成非菜单类型的路由的配置管理
	route.RESTful("/v1/auth/resource/func", &controllers.FuncSrvController{})
	route.RESTful("/v1/auth/privilege/user/domain", &controllers.UserDomainPrivilegeController{})
	route.RESTful("/v1/auth/privilege", &controllers.SysPrivilegeController{})
	route.RESTful("/v1/auth/privilege/domain", &controllers.SysPrivilegeDomainController{})
	route.RESTful("/v1/auth/privilege/role", &controllers.SysPrivilegeRoleController{})

	//role_resource_info
	route.Handler("GET", "/v1/auth/role/resource/get", controllers.RoleAndResourceCtl.GetResource)
	route.Handler("POST", "/v1/auth/role/resource/rights", controllers.RoleAndResourceCtl.HandleResource)
	route.Handler("GET", "/v1/auth/role/resource/details", controllers.RoleAndResourceCtl.ResourcePage)

	//role_info
	route.Handler("GET", "/v1/auth/role/get", controllers.RoleCtl.Get)
	route.Handler("POST", "/v1/auth/role/post", controllers.RoleCtl.Post)
	route.Handler("PUT", "/v1/auth/role/update", controllers.RoleCtl.Update)
	route.Handler("POST", "/v1/auth/role/delete", controllers.RoleCtl.Delete)

	//sys_auth_info
	route.Handler("GET", "/v1/auth/user/roles/get", controllers.UserRolesCtl.GetRolesByUserId)
	route.Handler("GET", "/v1/auth/user/search", controllers.UserCtl.Search)
	route.Handler("GET", "/v1/auth/user/roles/other", controllers.UserRolesCtl.GetOtherRoles)
	route.Handler("POST", "/v1/auth/user/roles/auth", controllers.UserRolesCtl.Auth)
	route.Handler("POST", "/v1/auth/user/roles/revoke", controllers.UserRolesCtl.Revoke)
	route.Handler("GET", "/v1/auth/role/query/user", controllers.UserRolesCtl.GetUserListByRoleId)
	route.Handler("GET", "/v1/auth/role/query/user/other", controllers.UserRolesCtl.GetOtherUserListByRoleId)

	//user_info
	route.Handler("GET", "/v1/auth/user/get", controllers.UserCtl.Get)
	route.Handler("POST", "/v1/auth/user/post", controllers.UserCtl.Post)
	route.Handler("PUT", "/v1/auth/user/put", controllers.UserCtl.Put)
	route.Handler("PUT", "/v1/auth/user/modify/passwd", controllers.UserCtl.ModifyPasswd)
	route.Handler("PUT", "/v1/auth/user/modify/status", controllers.UserCtl.ModifyStatus)
	route.Handler("POST", "/v1/auth/user/delete", controllers.UserCtl.Delete)
	route.Handler("GET", "/v1/auth/user/query", controllers.UserCtl.GetUserDetails)

	// help
	route.Handler("GET", "/v1/help/system/help", controllers.HelpCtl.Page)

	// pages
	route.Handler("GET", "/v1/auth/HandleLogsPage", controllers.HandleLogsCtl.Page)
	route.Handler("GET", "/v1/auth/domain/page", controllers.DomainCtl.Page)
	route.Handler("GET", "/v1/auth/batch/page", controllers.UserRolesCtl.Page)
	route.Handler("GET", "/v1/auth/resource/org/page", controllers.OrgCtl.Page)
	route.Handler("GET", "/v1/auth/resource/page", controllers.ResourceCtl.Page)
	route.Handler("GET", "/v1/auth/user/page", controllers.UserCtl.Page)
	route.Handler("GET", "/v1/auth/role/page", controllers.RoleCtl.Page)
	route.Handler("GET", "/v1/auth/swagger/page", controllers.SwaggerCtl.Page)
	route.Handler("GET", "/v1/auth/resource/service", controllers.GetServiceManagePage)
	route.Handler("GET", "/v1/auth/role/user/page", controllers.UserRolesCtl.UserPage)
	route.Handler("GET", "/v1/auth/privilege/page", controllers.GetSysPrivilegePage)
	route.Handler("GET", "/v1/auth/privilege/domain/page", controllers.GetPrivilegeDomainPage)
	route.Handler("GET", "/v1/auth/privilege/role/page", controllers.GetPrivilegeRolePage)
}
