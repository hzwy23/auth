package controllers

import (
	"github.com/asofdate/auth-core/groupcache"
	"github.com/asofdate/auth-core/models"
	"github.com/hzwy23/utils/crypto/sha1"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
)

var homePageMenusModel = new(models.HomePageMenusModel)

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
func SubSystemEntry(ctx router.Context) {
	defer hret.HttpPanic()

	ctx.Request.ParseForm()
	id := ctx.Request.FormValue("Id")

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
func IndexMenus(ctx router.Context) {
	defer hret.HttpPanic()
	ctx.Request.ParseForm()
	form := ctx.Request.Form

	id := form.Get("Id")
	typeId := form.Get("TypeId")

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

func AllMenusExceptButton(ctx router.Context) {
	ctx.Request.ParseForm()

	meunId := ctx.Request.FormValue("resId")

	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		logger.Error("Get user connect info failed. please login again, error info is:", err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	menus, err := homePageMenusModel.GetAllMenusExceptButton(jclaim.UserId, meunId)
	if err != nil {
		logger.Error("Get meuns failed. error info is :", err)
		hret.Error(ctx.ResponseWriter, 403, err.Error())
		return
	}
	ctx.ResponseWriter.Write(menus)
}
