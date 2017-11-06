package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/hzwy23/auth-core/dto"
	"github.com/hzwy23/auth-core/models"
	"github.com/hzwy23/auth-core/service"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/crypto/haes"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
	"github.com/hzwy23/utils/validator"
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
func HomePage(ctx router.Context) {
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

func LogoutSystem(ctx router.Context) {
	cookie := http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(ctx.ResponseWriter, &cookie)
	ctx.ResponseWriter.Header().Set("Authorization", "")
	hret.Success(ctx.ResponseWriter, "logout successfully")
	return
}

// 登录系统
func LoginSystem(ctx router.Context) {
	ctx.Request.ParseForm()
	form := ctx.Request.Form
	username := form.Get("username")
	password := form.Get("password")
	duration := form.Get("duration")

	if validator.IsEmpty(username) {
		rdto := dto.AuthDto{
			Username: username,
			RetCode:  422,
			RetMsg:   "鉴权失败，账号为空",
		}
		result(ctx.ResponseWriter, rdto)
		return
	}

	if validator.IsEmpty(password) {
		rdto := dto.AuthDto{
			Username: username,
			RetCode:  423,
			RetMsg:   "鉴权失败，密码为空",
		}
		result(ctx.ResponseWriter, rdto)
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
		//domainId, err := service.DomainService.GetDomainByUserid(username)
		//if err != nil {
		//	logger.Error(username, " 用户没有指定的域", err)
		//	retDto.RetCode = 426
		//	retDto.RetMsg = "获取用户域信息失败"
		//	result(ctx.ResponseWriter, retDto)
		//	return
		//}

		orgid, err := service.OrgService.GetOrgUnitId(username)
		if err != nil {
			logger.Error(username, " 用户没有指定机构", err)
			retDto.RetCode = 427
			retDto.RetMsg = "获取用户所在机构失败"
			result(ctx.ResponseWriter, retDto)
			return
		}

		et, err := strconv.ParseInt(duration, 10, 64)
		if err != nil || validator.IsEmpty(duration) {
			et = 17280
		}
		reqIP := utils.GetRequestIP(ctx.Request)
		token := jwt.GenToken(username, orgid, et, reqIP)
		cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: int(et)}
		http.SetCookie(ctx.ResponseWriter, &cookie)
		retDto.RetMsg = token
		result(ctx.ResponseWriter, retDto)
		return
	}
	result(ctx.ResponseWriter, retDto)
	return
}

func checkUserPassword(cdto dto.AuthDto) dto.AuthDto {
	// 加密用户信息
	pd, err := haes.Encrypt(cdto.Password)
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
