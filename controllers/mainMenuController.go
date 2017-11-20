package controllers

import (
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
	"net/http"
	"github.com/hzwy23/panda/crypto"
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
func SubSystemEntry(w http.ResponseWriter,r *http.Request) {
	defer hret.RecvPanic()

	r.ParseForm()
	id := r.FormValue("Id")

	// get user connection information from cookie.
	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	// get url of the id number.
	url, err := homePageMenusModel.GetUrl(jclaim.UserId, id)
	if err != nil {
		logger.Error(err)
		w.Write([]byte(url))
		return
	}

	key := crypto.Sha1(id, jclaim.UserId, url)

	if !groupcache.FileIsExist(key) {
		groupcache.RegisterStaticFile(key, url)
	}

	tpl, err := groupcache.GetStaticFile(key)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 404, i18n.PageNotFound(r))
		return
	}
	w.Write(tpl)
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
func IndexMenus(w http.ResponseWriter,r *http.Request) {
	defer hret.RecvPanic()
	r.ParseForm()

	form := r.Form
	id := form.Get("Id")
	typeId := form.Get("TypeId")

	// get user connection information from cookie

	claim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	ojs, err := homePageMenusModel.Get(id, typeId, claim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_query_menu"))
		return
	}
	w.Write(ojs)
}

func AllMenusExceptButton(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	meunId := r.FormValue("resId")

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error("Get user connect info failed. please login again, error info is:", err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	menus, err := homePageMenusModel.GetAllMenusExceptButton(jclaim.UserId, meunId)
	if err != nil {
		logger.Error("Get meuns failed. error info is :", err)
		hret.Error(w, 403, err.Error())
		return
	}
	w.Write(menus)
}
