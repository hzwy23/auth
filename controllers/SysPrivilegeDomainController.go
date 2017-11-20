package controllers

import (
	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
	"github.com/hzwy23/panda/uuid"
	"net/http"
	"time"
)

type SysPrivilegeDomainController struct {
	privilegeDomain models.SysPrivilegeDomain
}

// 查询权限与域关联关系
func GetPrivilegeDomainPage(w http.ResponseWriter, r *http.Request) {
	tpl, _ := service.ParseFile(r, "./views/hauth/sysPrivilegeDomain.tpl")
	privilegeId := r.FormValue("privilegeId")
	tpl.Execute(w, privilegeId)
}

func (this *SysPrivilegeDomainController) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	privilegeId := r.FormValue("privilegeId")
	typeCd := r.FormValue("typeCd")

	if len(privilegeId) == 0 {
		hret.Error(w, 423, "参数为空")
		return
	}

	if len(typeCd) == 0 {
		rst, err := this.privilegeDomain.Get(privilegeId)
		if err != nil {
			logger.Error(err)
			hret.Error(w, 421, err.Error())
			return
		}
		hret.Json(w, rst)
	} else {
		rst, err := this.privilegeDomain.GetUmmapDomain(privilegeId)
		if err != nil {
			logger.Error(err)
			hret.Error(w, 421, err.Error())
			return
		}
		hret.Json(w, rst)
	}
}

func (this *SysPrivilegeDomainController) Post(w http.ResponseWriter, r *http.Request) {
	var rows []entity.PrivilegeDomain
	err := panda.ParseForm(r, &rows, "JSON")
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}

	claim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 424, err.Error())
		return
	}

	for index, _ := range rows {
		rows[index].Uuid = uuid.Random()
		rows[index].CreateUser = claim.UserId
		rows[index].ModifyUser = claim.UserId
		rows[index].CreateTime = time.Now().Format("2006-01-02 15:04:05")
		rows[index].ModifyTime = rows[index].CreateTime
	}

	err = this.privilegeDomain.Post(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	hret.Success(w, i18n.Success(r))
}

func (this *SysPrivilegeDomainController) Put(w http.ResponseWriter, r *http.Request) {
	var row entity.PrivilegeDomain
	err := panda.ParseForm(r, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}

	claim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 424, err.Error())
		return
	}

	row.CreateUser = claim.UserId
	row.ModifyUser = claim.UserId
	row.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	row.ModifyTime = row.CreateTime

	err = this.privilegeDomain.Put(row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	hret.Success(w, i18n.Success(r))
}

func (this *SysPrivilegeDomainController) Delete(w http.ResponseWriter, r *http.Request) {
	var rows []entity.PrivilegeDomain
	err := panda.ParseForm(r, &rows, "JSON")
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, err.Error())
		return
	}
	err = this.privilegeDomain.Delete(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 424, err.Error())
		return
	}
	hret.Success(w, i18n.Success(r))
}
