package filter

import (
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils/router"
)

func AuthFilter() {
	// 处理系统内部路由
	router.InsertFilter("/*", router.BeforeExec, func(ctx router.Context) {
		service.Identify(ctx)
	}, true)
}

func init() {
	// 设置白名单
	service.AddConnUrl("/")
	service.AddConnUrl("/login")
	service.AddAuthUrl("/HomePage")

	/// 设置白名单
	service.AddAuthUrl("/v1/auth/main/menu")
	service.AddAuthUrl("/v1/auth/menu/all/except/button")
}
