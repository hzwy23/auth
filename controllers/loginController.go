package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/hzwy23/auth/dto"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda/crypto/aes"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
	"github.com/hzwy23/panda/validator"
	)

var indexModels = new(models.LoginModels)
var password = &models.PasswdModels{}

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
func HomePage(w http.ResponseWriter,r *http.Request) {
	defer hret.RecvPanic(func() {
		logger.Error("Get Home Page Failure.")
		http.Redirect(w,r,"/",302)
	})

	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		http.Redirect(w,r,"/",302)
		return
	}

	url := indexModels.GetDefaultPage(jclaim.UserId)
	h, err := template.ParseFiles(url)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_get_login_page"), err)
		return
	}
	h.Execute(w, jclaim.UserId)
}

func LogoutSystem(w http.ResponseWriter,r *http.Request) {
	cookie := http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	w.Header().Set("Authorization", "")
	hret.Success(w, "logout successfully")
	return
}

// 登录系统
func LoginSystem(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	form := r.Form
	username := form.Get("username")
	password := form.Get("password")
	duration := form.Get("duration")

	if validator.IsEmpty(username) {
		rdto := dto.AuthDto{
			Username: username,
			RetCode:  422,
			RetMsg:   "鉴权失败，账号为空",
		}
		result(w, rdto)
		return
	}

	if validator.IsEmpty(password) {
		rdto := dto.AuthDto{
			Username: username,
			RetCode:  423,
			RetMsg:   "鉴权失败，密码为空",
		}
		result(w, rdto)
		return
	}

	retDto := checkUserPassword(dto.AuthDto{
		Username: username,
		Password: password,
		Duration: duration,
		RetCode:  404,
		RetMsg:   "",
	})

	if retDto.RetCode == 200 {

		orgid, err := service.OrgService.GetOrgUnitId(username)
		if err != nil {
			logger.Error(username, " 用户没有指定机构", err)
			retDto.RetCode = 427
			retDto.RetMsg = "获取用户所在机构失败"
			result(w, retDto)
			return
		}

		et, err := strconv.ParseInt(duration, 10, 64)
		if err != nil || validator.IsEmpty(duration) {
			et = 17280
		}

		token,_ := jwt.GenToken(jwt.NewUserdata().SetUserId(username).SetOrgunitId(orgid))
		cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: int(et)}
		http.SetCookie(w, &cookie)
		retDto.RetMsg = token
		result(w, retDto)
		return
	}
	result(w, retDto)
	return
}

func checkUserPassword(cdto dto.AuthDto) dto.AuthDto {
	// 加密用户信息
	pd, err := aes.Encrypt(cdto.Password)
	if err != nil {
		logger.Error(err)
		cdto.RetCode = 434
		cdto.RetMsg = "加密用户信息失败"
		return cdto
	}
	_, code, _, msg := password.CheckPasswd(cdto.Username, pd)
	cdto.RetCode = code
	cdto.RetMsg = msg
	return cdto
}

func result(respone http.ResponseWriter, cdto dto.AuthDto) {
	msg, err := json.Marshal(cdto)
	if err != nil {
		respone.WriteHeader(http.StatusExpectationFailed)
		respone.Write([]byte(`{username:` + cdto.Username + `,RetCode:"431",retMsg:"format json type info failed."}`))
		return
	}
	respone.WriteHeader(cdto.RetCode)
	respone.Header().Set("Authorization", cdto.RetMsg)
	respone.Write(msg)
}
