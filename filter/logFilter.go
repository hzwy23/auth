package filter

import (
	"github.com/asofdate/sso-jwt-auth/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func LoggerFilter() {
	beego.InsertFilter("/*", beego.FinishRouter, func(ctx *context.Context) {
		go service.WriteHandleLogs(ctx)
	}, false)
}
