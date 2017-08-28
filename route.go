package auth_core

import (
	"github.com/asofdate/auth-core/controllers"
	"github.com/hzwy23/utils/router"
)

func Register() {

	router.Get("/HomePage", controllers.HomePage)
	router.Post("/login", controllers.LoginSystem)
	router.Any("/logout", controllers.LogoutSystem)
	router.Get("/", controllers.IndexPage)

	router.Post("/v1/auth/theme/update", controllers.ThemeCtl.Post)
	router.Put("/v1/auth/resource/config/theme", controllers.ThemeCtl.Put)
	router.Get("/v1/auth/resource/queryTheme", controllers.ThemeCtl.QueryTheme)

	router.Get("/v1/auth/index/entry", controllers.SubSystemEntry)
	router.Get("/v1/auth/main/menu", controllers.IndexMenus)
	router.Get("/v1/auth/menu/all/except/button", controllers.AllMenusExceptButton)
	router.Post("/v1/auth/passwd/update", controllers.PasswdController.PostModifyPasswd)

	//domain_info
	router.Get("/v1/auth/domain/get", controllers.DomainCtl.Get)
	router.Post("/v1/auth/domain/post", controllers.DomainCtl.Post)
	router.Post("/v1/auth/domain/delete", controllers.DomainCtl.Delete)
	router.Put("/v1/auth/domain/update", controllers.DomainCtl.Put)
	router.Get("/v1/auth/domain/details", controllers.DomainCtl.GetDetails)

	//handle_logs
	router.Get("/v1/auth/handle/logs/search", controllers.HandleLogsCtl.SerachLogs)
	router.Get("/v1/auth/handle/logs", controllers.HandleLogsCtl.GetHandleLogs)
	router.Get("/v1/auth/handle/logs/download", controllers.HandleLogsCtl.Download)

	//org_info
	router.Get("/v1/auth/org/get", controllers.OrgCtl.Get)
	router.Post("/v1/auth/org/insert", controllers.OrgCtl.Post)
	router.Put("/v1/auth/org/update", controllers.OrgCtl.Update)
	router.Post("/v1/auth/org/delete", controllers.OrgCtl.Delete)
	router.Get("/v1/auth/org/download", controllers.OrgCtl.Download)
	router.Post("/v1/auth/org/upload", controllers.OrgCtl.Upload)
	router.Get("/v1/auth/org/sub", controllers.OrgCtl.GetSubOrgInfo)
	router.Get("/v1/auth/org/details", controllers.OrgCtl.GetDetails)

	//resource_info
	router.Post("/v1/auth/resource/delete", controllers.ResourceCtl.Delete)
	router.Post("/v1/auth/resource/post", controllers.ResourceCtl.Post)
	router.Put("/v1/auth/resource/update", controllers.ResourceCtl.Update)
	router.Get("/v1/auth/resource/get", controllers.ResourceCtl.Get)
	router.Get("/v1/auth/resource/query", controllers.ResourceCtl.Query)
	router.Get("/v1/auth/resource/node", controllers.ResourceCtl.GetNodes)
	router.Get("/v1/auth/resource/subsystem", controllers.ResourceCtl.GetSubsystemList)

	// 功能服务路由注册
	// 主要完成非菜单类型的路由的配置管理
	router.RESTful("/v1/auth/resource/func", &controllers.FuncSrvController{})
	router.RESTful("/v1/auth/privilege/user/domain", &controllers.UserDomainPrivilegeController{})
	router.RESTful("/v1/auth/privilege", &controllers.SysPrivilegeController{})
	router.RESTful("/v1/auth/privilege/domain", &controllers.SysPrivilegeDomainController{})
	router.RESTful("/v1/auth/privilege/role", &controllers.SysPrivilegeRoleController{})

	//role_resource_info
	router.Get("/v1/auth/role/resource/get", controllers.RoleAndResourceCtl.GetResource)
	router.Post("/v1/auth/role/resource/rights", controllers.RoleAndResourceCtl.HandleResource)
	router.Get("/v1/auth/role/resource/details", controllers.RoleAndResourceCtl.ResourcePage)

	//role_info
	router.Get("/v1/auth/role/get", controllers.RoleCtl.Get)
	router.Post("/v1/auth/role/post", controllers.RoleCtl.Post)
	router.Put("/v1/auth/role/update", controllers.RoleCtl.Update)
	router.Post("/v1/auth/role/delete", controllers.RoleCtl.Delete)

	//sys_auth_info
	router.Get("/v1/auth/user/roles/get", controllers.UserRolesCtl.GetRolesByUserId)
	router.Get("/v1/auth/user/search", controllers.UserCtl.Search)
	router.Get("/v1/auth/user/roles/other", controllers.UserRolesCtl.GetOtherRoles)
	router.Post("/v1/auth/user/roles/auth", controllers.UserRolesCtl.Auth)
	router.Post("/v1/auth/user/roles/revoke", controllers.UserRolesCtl.Revoke)
	router.Get("/v1/auth/role/query/user", controllers.UserRolesCtl.GetUserListByRoleId)
	router.Get("/v1/auth/role/query/user/other", controllers.UserRolesCtl.GetOtherUserListByRoleId)

	//user_info
	router.Get("/v1/auth/user/get", controllers.UserCtl.Get)
	router.Post("/v1/auth/user/post", controllers.UserCtl.Post)
	router.Put("/v1/auth/user/put", controllers.UserCtl.Put)
	router.Put("/v1/auth/user/modify/passwd", controllers.UserCtl.ModifyPasswd)
	router.Put("/v1/auth/user/modify/status", controllers.UserCtl.ModifyStatus)
	router.Post("/v1/auth/user/delete", controllers.UserCtl.Delete)
	router.Get("/v1/auth/user/query", controllers.UserCtl.GetUserDetails)

	// help
	router.Get("/v1/help/system/help", controllers.HelpCtl.Page)

	// pages
	router.Get("/v1/auth/HandleLogsPage", controllers.HandleLogsCtl.Page)
	router.Get("/v1/auth/domain/page", controllers.DomainCtl.Page)
	router.Get("/v1/auth/batch/page", controllers.UserRolesCtl.Page)
	router.Get("/v1/auth/resource/org/page", controllers.OrgCtl.Page)
	router.Get("/v1/auth/resource/page", controllers.ResourceCtl.Page)
	router.Get("/v1/auth/user/page", controllers.UserCtl.Page)
	router.Get("/v1/auth/role/page", controllers.RoleCtl.Page)
	router.Get("/v1/auth/swagger/page", controllers.SwaggerCtl.Page)
	router.Get("/v1/auth/resource/service", controllers.GetServiceManagePage)
	router.Get("/v1/auth/role/user/page", controllers.UserRolesCtl.UserPage)
	router.Get("/v1/auth/privilege/page", controllers.GetSysPrivilegePage)
	router.Get("/v1/auth/privilege/domain/page", controllers.GetPrivilegeDomainPage)
	router.Get("/v1/auth/privilege/role/page", controllers.GetPrivilegeRolePage)
}
