package controllers

import (
	"github.com/asofdate/sso-jwt-auth/groupcache"
	"github.com/asofdate/sso-jwt-auth/models"
	"github.com/asofdate/sso-jwt-auth/utils/crypto/sha1"
	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/i18n"
	"github.com/asofdate/sso-jwt-auth/utils/jwt"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/asofdate/sso-core/utils"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

var homePageMenusModel = new(models.HomePageMenusModel)
var resourceModel = new(models.ResourceModel)

// swagger:operation GET /v1/auth/index/entry StaticFiles SubSystemEntry
//
// According to the ID number, return subsystem information page
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: Id
//   in: query
//   description: subsystem id number.
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
//   '403':
//     description: disconnect, please login.
//   '404':
//     description: page not found
func SubSystemEntry(ctx *context.Context) {
	defer hret.HttpPanic()

	ctx.Request.ParseForm()
	id := ctx.Request.FormValue("Id")
	innerFlag := ctx.Request.FormValue("innerFlag")

	if innerFlag == "false" {
		// 使用反向代理，获取子系统首页信息
		utils.SysIndexReverProxy(ctx)
		return
	}

	// get user connection information from cookie.
	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// get url of the id number.
	url, err := homePageMenusModel.GetUrl(jclaim.UserId, id)
	if err != nil {
		logger.Error(err)
		ctx.WriteString(url)
		return
	}

	key := sha1.GenSha1Key(id, jclaim.UserId, url)

	if !groupcache.FileIsExist(key) {
		groupcache.RegisterStaticFile(key, url)
	}

	tpl, err := groupcache.GetStaticFile(key)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}
	ctx.ResponseWriter.Write(tpl)
}

// swagger:operation GET /v1/auth/main/menu HomePageMenus HomePageMenus
//
// If the request is successful, will return the user to be able to access the menu information
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: TypeId
//   in: query
//   description: The menu type, 1 means home page ,2 means subsystem page
//   required: true
//   type: string
//   format:
// - name: Id
//   in: query
//   description: This up menu id , the response will return the lower menu information of the up menu id
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
//   '403':
//     description: disconnect
//   '421':
//     description: get menu information failed.
func IndexMenus(ctx *context.Context) {
	defer hret.HttpPanic()
	ctx.Request.ParseForm()
	form := ctx.Request.Form

	id := form.Get("Id")
	typeId := form.Get("TypeId")

	// typeId == "1" 表示获取子系统菜单
	// typeId == "0" 表示获取首页菜单，否也菜单不允许使用方向代理获取
	// 首页菜单信息，必须配置到SSO系统中
	if typeId == "1" {
		// 获取子系统菜单信息
		// 子系统中菜单交个子系统自行实现
		// 跳转到子系统中，获取子系统中的菜单信息
		innerFlag, err := resourceModel.GetInnerFlag(id)
		if err != nil {
			logs.Error(err)
			hret.Error(ctx.ResponseWriter, 404, "查询菜单资源属性失败")
			return
		}

		if innerFlag == "false" {
			// 非内部系统，反向代理获取子系统菜单信息
			// 子系统菜单格式必须按照规定时限，否则sso系统中的菜单构建函数将会失效
			utils.SsoSubsystemMenuReverProxy(ctx)
			return
		}
	}

	// get user connection information from cookie
	cookie, _ := ctx.Request.Cookie("Authorization")
	claim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	ojs, err := homePageMenusModel.Get(id, typeId, claim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_query_menu"))
		return
	}
	ctx.ResponseWriter.Write(ojs)
}
