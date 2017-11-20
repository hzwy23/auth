package filter

import (
	"strings"
	"github.com/hzwy23/auth/service"
	"net/http"
	"github.com/hzwy23/panda/route"
)

type AuthFilter struct {
}

func (this *AuthFilter)staticFile(url string)bool{
	if strings.HasPrefix(url,"/static"){
		return true
	}
	return url == "/favicon.ico"
}

func (this *AuthFilter) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if this.staticFile(r.URL.Path) || service.Identify(w,r) {
		nw:=route.NewResponse(w)
		next(nw, r)
	}
}

func init() {
	// 设置白名单，免认证请求
	service.AddConnUrl("/")
	service.AddConnUrl("/v1/auth/login")

	/// 设置白名单，免授权请求
	service.AddAuthUrl("/v1/auth/logout")
	service.AddAuthUrl("/v1/auth/theme/update")
	service.AddAuthUrl("/v1/auth/user/query")
	service.AddAuthUrl("/v1/auth/HomePage")
	service.AddAuthUrl("/v1/auth/main/menu")
	service.AddAuthUrl("/v1/auth/index/entry")
	service.AddAuthUrl("/v1/auth/privilege/user/domain")
	service.AddAuthUrl("/v1/auth/menu/all/except/button")
}
