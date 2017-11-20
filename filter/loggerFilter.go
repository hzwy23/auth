package filter

import (
	"github.com/hzwy23/auth/service"
	"net/http"
)

type LoggerFilter struct {
}

func (this *LoggerFilter) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
	go service.WriteHandleLogs(w, r)
}

func init() {
	// 开启操作日志监听
	go service.LogSync()
}
