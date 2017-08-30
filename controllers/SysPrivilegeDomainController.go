package controllers

import (
	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/models"
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
	"github.com/hzwy23/utils/uuid"
)

type SysPrivilegeDomainController struct {
	privilegeDomain models.SysPrivilegeDomain
	router.Controller
}

// 查询权限与域关联关系
func GetPrivilegeDomainPage(ctx router.Context) {
	tpl, _ := service.ParseFile(ctx, "./views/hauth/sysPrivilegeDomain.tpl")
	privilegeId := ctx.Request.FormValue("privilegeId")
	tpl.Execute(ctx.ResponseWriter, privilegeId)
}

func (this *SysPrivilegeDomainController) Get() {
	this.Ctx.Request.ParseForm()

	privilegeId := this.Ctx.Request.FormValue("privilegeId")
	typeCd := this.Ctx.Request.FormValue("typeCd")

	if len(privilegeId) == 0 {
		hret.Error(this.Ctx.ResponseWriter, 423, "参数为空")
		return
	}

	if len(typeCd) == 0 {
		rst, err := this.privilegeDomain.Get(privilegeId)
		if err != nil {
			logger.Error(err)
			hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
			return
		}
		hret.Json(this.Ctx.ResponseWriter, rst)
	} else {
		rst, err := this.privilegeDomain.GetUmmapDomain(privilegeId)
		if err != nil {
			logger.Error(err)
			hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
			return
		}
		hret.Json(this.Ctx.ResponseWriter, rst)
	}
}

func (this *SysPrivilegeDomainController) Post() {
	var rows []entity.PrivilegeDomain
	err := utils.ParseForm(this.Ctx.Request, &rows, "JSON")
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}

	claim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 424, err.Error())
		return
	}

	for index, _ := range rows {
		rows[index].Uuid = uuid.GenUUID()
		rows[index].CreateUser = claim.UserId
		rows[index].ModifyUser = claim.UserId
		rows[index].CreateTime = utils.GetCurrentTime()
		rows[index].ModifyTime = rows[index].CreateTime
	}

	err = this.privilegeDomain.Post(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *SysPrivilegeDomainController) Put() {
	var row entity.PrivilegeDomain
	err := utils.ParseForm(this.Ctx.Request, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}

	claim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 424, err.Error())
		return
	}

	row.CreateUser = claim.UserId
	row.ModifyUser = claim.UserId
	row.CreateTime = utils.GetCurrentTime()
	row.ModifyTime = row.CreateTime

	err = this.privilegeDomain.Put(row)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *SysPrivilegeDomainController) Delete() {
	var rows []entity.PrivilegeDomain
	err := utils.ParseForm(this.Ctx.Request, &rows, "JSON")
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	err = this.privilegeDomain.Delete(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 424, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}
