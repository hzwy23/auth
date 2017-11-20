package controllers

import (
	"net/http"
	"time"

	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
)

type SysPrivilegeController struct {
	sysPrivilege models.SysPrivilege
}

func GetSysPrivilegePage(w http.ResponseWriter, r *http.Request) {
	rst, err := groupcache.GetStaticFile("SysPrivilegePage")
	if err != nil {
		logger.Error(err)
		hret.Error(w, 404, i18n.Get(r, "as_of_date_page_not_exist"))
		return
	}

	hz, err := service.ParseText(r, string(rst))
	if err != nil {
		logger.Error(err)
		hret.Error(w, 404, i18n.Get(r, "as_of_date_page_not_exist"))
		return
	}
	hz.Execute(w, nil)
}

// 查询权限定义信息
func (this *SysPrivilegeController) Get(w http.ResponseWriter, r *http.Request) {
	rst, err := this.sysPrivilege.Get()
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	hret.Json(w, rst)
}

// 新增权限代码
func (this *SysPrivilegeController) Post(w http.ResponseWriter, r *http.Request) {
	var row entity.SysPrivilegeEntity

	err := panda.ParseForm(r, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, err.Error())
		return
	}

	claim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 424, err.Error())
		return
	}

	curTime := time.Now().Format("2006-01-02 15:04:05")

	row.ModifyTime = curTime
	row.ModifyUser = claim.UserId
	row.CreateTime = curTime
	row.CreateUser = claim.UserId

	err = this.sysPrivilege.Post(row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	hret.Success(w, "success")
}

// 更新权限代码
func (this *SysPrivilegeController) Put(w http.ResponseWriter, r *http.Request) {
	var row entity.SysPrivilegeEntity

	err := panda.ParseForm(r, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, err.Error())
		return
	}

	claim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 424, err.Error())
		return
	}

	curTime := time.Now().Format("2006-01-02 15:04:05")

	row.ModifyTime = curTime
	row.ModifyUser = claim.UserId
	row.CreateTime = curTime
	row.CreateUser = claim.UserId

	err = this.sysPrivilege.Put(row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	hret.Success(w, "success")
}

// 删除权限代码
func (this *SysPrivilegeController) Delete(w http.ResponseWriter, r *http.Request) {
	var rows []entity.SysPrivilegeEntity
	err := panda.ParseForm(r, &rows, "JSON")
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	err = this.sysPrivilege.Delete(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, err.Error())
		return
	}
	hret.Success(w, "success")
}

func init() {
	groupcache.RegisterStaticFile("SysPrivilegePage", "./views/hauth/sysPrivilege.tpl")
}
