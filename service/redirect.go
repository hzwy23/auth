package service

import (
	"fmt"
	"net/http"

	"os"
	"path/filepath"

	"github.com/asofdate/sso-jwt-auth/utils/config"
	"github.com/asofdate/sso-jwt-auth/utils/hret"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
)

// 将80端口,重定向到443端口
func RedictToHtpps() {

	defer hret.HttpPanic()

	red, err := config.GetConfig(filepath.Join(os.Getenv("HBIGDATA_HOME"), "conf", "asofdate.conf"))
	if err != nil {
		logger.Error(err)
		return
	}

	portMap, err := red.Get("PortMap")
	if err != nil {
		logger.Error(err)
		return
	}

	targetHost, err := red.Get("TargetHost")
	if err != nil {
		logger.Error(err)
		return
	}

	if portMap == "true" {

		var redirectHandle = http.NewServeMux()

		redirectHandle.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			http.Redirect(w, r, targetHost, http.StatusTemporaryRedirect)

		})

		fmt.Println("即将开启服务，将http默认的80端口请求消息转发到https默认端口443上...")

		err := http.ListenAndServe(":80", redirectHandle)

		if err != nil {

			fmt.Println("您的用户没有权限开启80端口，或80端口已经被占用，无法将http默认80端口消息转发到https默认端口443上，请使用https协议访问系统", err)

		}
	}
}
