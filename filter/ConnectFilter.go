package filter

import (
	"github.com/asofdate/sso-jwt-auth/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func ConnectFilter() {
	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
		service.CheckConnection(ctx.ResponseWriter, ctx.Request)
	}, false)
}
