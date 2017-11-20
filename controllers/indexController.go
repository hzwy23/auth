package controllers

import (
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/logger"
	"net/http"
)

// swagger:operation GET / StaticFiles IndexPage
//
// 系统首页页面
//
// API将会返回系统首页页面给客户端
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
func IndexPage(w http.ResponseWriter, r *http.Request) {
	rst, err := groupcache.GetStaticFile("AsofdateIndexPage")
	if err != nil {
		logger.Error(err)
		hret.Error(w, 404, i18n.PageNotFound(r))
		return
	}
	w.Write(rst)
}

func init() {
	groupcache.RegisterStaticFile("AsofdateIndexPage", "./views/hauth/login.tpl")
}
