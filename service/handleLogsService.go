package service

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/models"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
)

var log_buf = make(chan entity.HandleLogBuf, 40960)
var handleModel = &models.HandleLogMode{}

func WriteHandleLogs(ctx router.Context) {
	defer hret.HttpPanic()

	if strings.HasPrefix(ctx.Request.URL.Path, "/") {
		var one entity.HandleLogBuf
		status := ctx.ResponseWriter.Status
		if status == 0 {
			status = 200
		}
		one.Ret_status = strconv.Itoa(status)
		one.Req_url = ctx.Request.URL.Path
		one.Req_body = formencode(ctx.Request.Form)
		one.Client_ip = ctx.Input.IP()
		one.Req_method = ctx.Request.Method

		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := jwt.ParseJwt(cookie.Value)
		if err != nil {
			one.User_id = one.Client_ip
			one.Domain_id = one.Client_ip
		} else {
			one.User_id = jclaim.UserId
			one.Domain_id = jclaim.OrgUnitId
		}
		logger.Infow("http request:", "user_id", one.User_id, "client_up", one.Client_ip, "ret_status", one.Ret_status, "req_method", one.Req_method, "req_url", one.Req_url, "domain_id", one.Domain_id, "req_body", one.Req_body)
		log_buf <- one
	}
}

func formencode(form url.Values) string {
	rst := make(map[string]string)
	for key, val := range form {
		if key == "_" {
			continue
		}
		rst[key] = val[0]
	}

	str, _ := json.Marshal(rst)
	return string(str)
}

func savelogs(log_buf []entity.HandleLogBuf) {
	handleModel.Post(log_buf)
}

func LogSync() {
	var buf []entity.HandleLogBuf
	for {
		select {
		case <-time.After(time.Second * 10):
			// sync handle logs to database per 5 second.
			if len(buf) == 0 {
				continue
			}
			go savelogs(buf)
			buf = make([]entity.HandleLogBuf, 0)
		case val, ok := <-log_buf:
			if ok {
				buf = append(buf, val)
				if len(buf) > 1000 {
					go savelogs(buf)
					buf = make([]entity.HandleLogBuf, 0)
				}
			}
		}
	}
}
