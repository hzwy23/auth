package auth

import (
	"strings"
	"github.com/hzwy23/panda/config"
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

type WebApp struct{
	Port        string
	HttpsEnable bool
	HttpsCrt    string
	HttpsKey    string
}

var AppConfig = &WebApp{
	Port:":8080",
	HttpsEnable:false,
	HttpsCrt:"./conf/rui.crt",
	HttpsKey:"./conf/rui.key",
}

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
	// 从配置文件获取端口号
	c,err := config.Load("conf/app.conf",config.INI)
	if err != nil {
		logger.Warn("读取配置文件conf/app.conf文件失败，使用默认参数启动服务")
		http.ListenAndServe(AppConfig.Port,middle)
		return
	}

	// 从配置文件中读取端口号
	port,err := c.Get("ServerPort")
	if err != nil {
		logger.Warn("http.port 不存在，使用默认端口启动服务")
		http.ListenAndServe(AppConfig.Port,middle)
		return
	}
	AppConfig.Port = port

	// 检查是否开启https
	httpsEnable,err := c.Get("https.enable")
	if err != nil {
		logger.Warn("http.port 不存在，使用默认端口启动服务")
		http.ListenAndServe(AppConfig.Port,middle)
		return
	}
	AppConfig.HttpsEnable = (strings.ToLower(httpsEnable)=="true")
	// 使用http启动web服务
	if !AppConfig.HttpsEnable {
		logger.Info("start web server with http.")
		http.ListenAndServe(":"+AppConfig.Port, middle)
		return
	}

	// 使用https启动服务，读取crt文件地址
	httpsCrt,err := c.Get("https.crt")
	if err != nil {
		logger.Warn("http.port 不存在，使用默认端口启动服务")
	} else {
		AppConfig.HttpsCrt = httpsCrt
	}

	// 使用https启动服务，读取key文件地址
	httpsKey,err := c.Get("https.key")
	if err != nil {
		logger.Warn("http.port 不存在，使用默认端口启动服务")
	} else {
		AppConfig.HttpsKey = httpsKey
	}

	http.ListenAndServeTLS(":"+AppConfig.Port,AppConfig.HttpsCrt,AppConfig.HttpsKey,middle)
	
}
