package service

import (
	"encoding/json"
	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/route"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var log_buf = make(chan entity.HandleLogBuf, 40960)
var handleModel = &models.HandleLogMode{}

func WriteHandleLogs(w http.ResponseWriter, r *http.Request) {
	defer hret.RecvPanic()

	if nw, ok := w.(*route.Response); ok {
		var one entity.HandleLogBuf
		one.Ret_status = strconv.Itoa(nw.Status)
		one.Req_url = r.URL.Path
		one.Req_body = formencode(r.Form)
		one.Client_ip = route.RequestIP(r)
		one.Req_method = r.Method

		jclaim, err := jwt.ParseHttp(r)
		if err != nil {
			one.User_id = one.Client_ip
			one.Domain_id = one.Client_ip
		} else {
			one.User_id = jclaim.UserId
			one.Domain_id = jclaim.OrgUnitId
		}
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
