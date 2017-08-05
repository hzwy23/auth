package controllers

import (
	"html/template"

	"github.com/asofdate/sso-jwt-auth/models"
	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/i18n"
	"github.com/asofdate/sso-jwt-auth/utils/jwt"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/astaxie/beego/context"
)

var indexModels = new(models.LoginModels)

// swagger:operation GET /HomePage StaticFiles IndexPage
//
// 返回用户登录后的主菜单页面
//
// 用户登录成功后,将会根据用户主题,返回用户的主菜单页面.
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
func HomePage(ctx *context.Context) {
	defer hret.HttpPanic(func() {
		logger.Error("Get Home Page Failure.")
		ctx.Redirect(302, "/")
	})

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		logger.Error(err)
		ctx.Redirect(302, "/")
		return
	}

	url := indexModels.GetDefaultPage(jclaim.UserId)
	h, err := template.ParseFiles(url)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_get_login_page"), err)
		return
	}
	h.Execute(ctx.ResponseWriter, jclaim.UserId)
}
