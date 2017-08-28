package controllers

import (
	"github.com/asofdate/auth-core/groupcache"
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/router"
)

type swaggerController struct {
}

var SwaggerCtl = &swaggerController{}

// swagger:operation GET /v1/auth/swagger/page StaticFiles swaggerController
//
// API文档页面
//
// 返回API信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: success
func (this swaggerController) Page(ctx router.Context) {
	if !service.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := groupcache.GetStaticFile("SwaggerPage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
		return
	}

	ctx.ResponseWriter.Write(rst)
}

func init() {
	groupcache.RegisterStaticFile("SwaggerPage", "./views/help/swagger_index.html")
}
