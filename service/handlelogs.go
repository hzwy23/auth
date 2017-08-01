package service

import (
	"strings"

	"strconv"

	"time"

	"net/url"

	"encoding/json"

	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/jwt"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/dbobj"
)

var log_buf = make(chan handleLogBuf, 40960)

type handleLogBuf struct {
	User_id    string `json:"user_id"`
	Client_ip  string `json:"client_ip"`
	Ret_status string `json:"ret_status"`
	Req_method string `json:"req_method"`
	Req_url    string `json:"req_url"`
	Domain_id  string `json:"domain_id"`
	Req_body   string `json:"req_body"`
}

func WriteHandleLogs(ctx *context.Context) {
	defer hret.HttpPanic()

	if strings.HasPrefix(ctx.Request.URL.Path, "/") {
		var one handleLogBuf
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
			one.Domain_id = jclaim.DomainId
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

func savelogs(log_buf []handleLogBuf) {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
	}

	for _, val := range log_buf {
		_, err := tx.Exec(hauth_service_001, val.User_id, val.Client_ip, val.Ret_status, val.Req_method, val.Req_url, val.Domain_id, val.Req_body)
		if err != nil {
			tx.Rollback()
			logger.Error("同步日志信息到数据库失败")
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		logger.Error("同步日志信息到数据库失败")
	}
}

func LogSync() {
	var buf []handleLogBuf
	for {
		select {
		case <-time.After(time.Second * 5):
			// sync handle logs to database per 5 second.
			if len(buf) == 0 {
				continue
			}
			go savelogs(buf)
			buf = make([]handleLogBuf, 0)
		case val := <-log_buf:
			buf = append(buf, val)
			if len(buf) > 1000 {
				go savelogs(buf)
				buf = make([]handleLogBuf, 0)
			}
		}
	}
}

func init() {
	go LogSync()
}
