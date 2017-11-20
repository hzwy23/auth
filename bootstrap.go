package auth

import (
	"github.com/hzwy23/panda/route"
	"net/http"
	"sync"

	"github.com/hzwy23/auth/filter"
	"github.com/hzwy23/panda/logger"
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

	for key, fc := range regApp {
		logger.Info("register App, name is:", key)
		fc()
	}

	// 开启权限过滤
	authFilter := &filter.AuthFilter{}
	// 开启日志过滤
	loggerFilter := &filter.LoggerFilter{}

	// 注册路由
	Register()

	// 注册静态路由
	route.ServeFiles("/static", http.Dir("./static"))

	// 创建中间件
	middle := route.NewMiddleware(authFilter, loggerFilter, route.Wrap(route.DefaultRouter()))

	// 启动服务
	http.ListenAndServe(":8080", middle)
}
