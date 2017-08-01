package sso_jwt_auth

import (
	"sync"

	"github.com/asofdate/sso-jwt-auth/filter"
	"github.com/asofdate/sso-jwt-auth/service"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/astaxie/beego"
)

type RegisterFunc func()

// key 应用名称
// value 注册路由方法
var regApp = make(map[string]RegisterFunc)
var regLock = new(sync.RWMutex)

func AppRegister(name string, registerFunc RegisterFunc) {
	regLock.Lock()
	defer regLock.Unlock()
	if _, ok := regApp[name]; ok {
		panic("应用已经被注册，无法再次注册")
	} else {
		regApp[name] = registerFunc
	}
}

func Bootstrap() {
	// 开启消息，
	// 将80端口的请求，重定向到443上
	go service.RedictToHtpps()
	for key, fc := range regApp {
		logger.Info("register App, name is:", key)
		fc()
	}
	filter.LoggerFilter()
	// 启动beego服务
	beego.Run()
}
