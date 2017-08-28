package filter

import (
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils/router"
)

func LoggerFilter() {

	// 开启操作日志监听
	go service.LogSync()

	// 拦截http请求，写入请求记录到日志缓存队列
	router.InsertFilter("/*", router.FinishRouter, func(ctx router.Context) {
		go service.WriteHandleLogs(ctx)
	}, false)
}
