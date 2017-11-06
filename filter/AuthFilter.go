package filter

import (
	"github.com/hzwy23/utils/router"
	"github.com/hzwy23/auth-core/service"
)

func AuthFilter() {
	// 处理系统内部路由
	router.InsertFilter("/*", router.BeforeExec, func(ctx router.Context) {
		service.Identify(ctx)
	}, true)
}

func init() {
	// 设置白名单，免认证请求
	service.AddConnUrl("/")
	service.AddConnUrl("/v1/auth/login")

	/// 设置白名单，免授权请求
	service.AddAuthUrl("/v1/auth/logout")
	service.AddAuthUrl("/v1/auth/theme/update")
	service.AddAuthUrl("/v1/auth/user/query")
	service.AddAuthUrl("/v1/auth/HomePage")
	service.AddAuthUrl("/v1/auth/main/menu")
	service.AddAuthUrl("/v1/auth/index/entry")
	service.AddAuthUrl("/v1/auth/privilege/user/domain")
	service.AddAuthUrl("/v1/auth/menu/all/except/button")
}
