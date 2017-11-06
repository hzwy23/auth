package auth_core

import (
	"sync"

	"github.com/hzwy23/auth-core/filter"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
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

	// 开启路由追踪
	filter.LoggerFilter()

	// 用户连接与权限校验
	filter.AuthFilter()

	// 启动beego服务
	router.Run()
}
