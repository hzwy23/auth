package service

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
)

type BasicAuthFilter struct {
	// filterAuthUrl 存放不需要校验用户权限的路由信息
	filterAuthUrl map[string]bool
	alock         *sync.RWMutex

	// filterConnUrl 存放不需要校验连接的路由信息
	filterConnUrl map[string]bool
	clock         *sync.RWMutex
}

var BAFilter = NewBasicAuthFilter()

func BasicAuth(r *http.Request) bool {
	return BAFilter.basicAuth(r)
}

func AddConnUrl(url string) {
	BAFilter.AddConnUrl(url)
}

func AddAuthUrl(url string) {
	BAFilter.AddAuthUrl(url)
}

func Identify(ctx router.Context) {
	BAFilter.Identify(ctx)
}

func NewBasicAuthFilter() *BasicAuthFilter {
	return &BasicAuthFilter{
		filterAuthUrl: make(map[string]bool),
		alock:         new(sync.RWMutex),
		filterConnUrl: make(map[string]bool),
		clock:         new(sync.RWMutex),
	}
}

func (this *BasicAuthFilter) AddAuthUrl(url string) {
	this.alock.Lock()
	defer this.alock.Unlock()
	this.filterAuthUrl[url] = true
}

func (this *BasicAuthFilter) AddConnUrl(url string) {
	this.clock.Lock()
	defer this.clock.Unlock()
	this.filterConnUrl[url] = true
}

func (this *BasicAuthFilter) basicAuth(r *http.Request) bool {
	jclaim, err := jwt.GetJwtClaims(r)
	if err != nil {
		logger.Error(err)
		return false
	}
	if utils.IsAdmin(jclaim.UserId) {
		return true
	}

	method := r.Method
	if method == http.MethodPost && r.FormValue("_method") == http.MethodDelete {
		method = http.MethodDelete
	}
	logger.Debug("basicAuth,method is:", method, ",path is:", r.URL.Path, ",user is:", jclaim.UserId)
	return RouteService.CheckUrlAuth(jclaim.UserId, r.URL.Path, method)
}

// 校验用户权限信息
func (this *BasicAuthFilter) Identify(ctx router.Context) {
	// 校验连接信息
	this.clock.RLock()
	if _, yes := this.filterConnUrl[ctx.Request.URL.Path]; !yes {
		if !jwt.CheckToken(ctx.Request) {
			this.clock.RUnlock()
			hz, _ := template.ParseFiles("./views/hauth/disconnect.tpl")
			hz.Execute(ctx.ResponseWriter, nil)
			ctx.ResponseWriter.Started = true
			return
		}
	} else {
		this.clock.RUnlock()
		return
	}
	this.clock.RUnlock()

	// 校验授权信息
	this.alock.RLock()
	if _, yes := this.filterAuthUrl[ctx.Request.URL.Path]; !yes {
		flag := this.basicAuth(ctx.Request)
		if !flag {
			logger.Info("您没有被授权访问：", ctx.Request.URL)
			hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
			ctx.ResponseWriter.Started = true
		}
	}
	this.alock.RUnlock()
	return
}
