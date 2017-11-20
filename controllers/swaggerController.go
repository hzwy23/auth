package controllers

import (
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"net/http"
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
func (this swaggerController) Page(w http.ResponseWriter, r *http.Request) {
	if !service.BasicAuth(r) {
		hret.Error(w, 403, i18n.NoAuth(r))
		return
	}

	rst, err := groupcache.GetStaticFile("SwaggerPage")
	if err != nil {
		hret.Error(w, 404, i18n.Get(r, "as_of_date_page_not_exist"))
		return
	}

	w.Write(rst)
}

func init() {
	groupcache.RegisterStaticFile("SwaggerPage", "./views/help/swagger_index.html")
}
