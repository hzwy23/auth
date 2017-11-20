package controllers

import (
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"net/http"
)

type helpController struct {
}

var HelpCtl = &helpController{}

// swagger:operation GET /v1/help/system/help StaticFiles helpController
//
// 系统帮助页面
//
// 将会返回系统帮助首页,其中包含了系统管理操作文档,API文档
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: all domain information
func (this helpController) Page(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	if !service.BasicAuth(r) {
		hret.Error(w, 403, i18n.NoAuth(r))
		return
	}

	rst, err := groupcache.GetStaticFile("AsofdateHelpPage")
	if err != nil {
		hret.Error(w, 404, i18n.PageNotFound(r))
		return
	}
	w.Write(rst)
}

func init() {
	groupcache.RegisterStaticFile("AsofdateHelpPage", "./views/help/auth_help.tpl")
}
