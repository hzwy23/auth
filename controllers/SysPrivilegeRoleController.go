package controllers

import (
	"html/template"

	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/models"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
	"github.com/hzwy23/utils/uuid"
)

type SysPrivilegeRoleController struct {
	privilegeRole models.SysPrivilegeRole
	router.Controller
}

// 查询权限与域关联关系
func GetPrivilegeRolePage(ctx router.Context) {
	tpl, _ := template.ParseFiles("./views/hauth/sysPrivilegeRole.tpl")
	privilegeId := ctx.Request.FormValue("privilegeId")
	tpl.Execute(ctx.ResponseWriter, privilegeId)
}

func (this *SysPrivilegeRoleController) Get() {
	this.Ctx.Request.ParseForm()

	privilegeId := this.Ctx.Request.FormValue("privilegeId")
	typeCd := this.Ctx.Request.FormValue("typeCd")

	if len(privilegeId) == 0 {
		hret.Error(this.Ctx.ResponseWriter, 423, "参数为空")
		return
	}

	if len(typeCd) == 0 {
		rst, err := this.privilegeRole.Get(privilegeId)
		if err != nil {
			logger.Error(err)
			hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
			return
		}
		hret.Json(this.Ctx.ResponseWriter, rst)
	} else {
		rst, err := this.privilegeRole.GetUnMapRole(privilegeId)
		if err != nil {
			logger.Error(err)
			hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
			return
		}
		hret.Json(this.Ctx.ResponseWriter, rst)
	}
}

func (this *SysPrivilegeRoleController) Post() {
	var rows []entity.PrivilegeRole
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
		rows[index].CreateTime = utils.GetCurrentTime()
		rows[index].ModifyUser = claim.UserId
		rows[index].ModifyTime = rows[index].CreateTime
	}
	err = this.privilegeRole.Post(rows)
	if err != nil {
		logger.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}

	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *SysPrivilegeRoleController) Put() {
	var row entity.PrivilegeRole
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
	row.CreateTime = utils.GetCurrentTime()
	row.ModifyUser = claim.UserId
	row.ModifyTime = row.CreateTime

	err = this.privilegeRole.Put(row)
	if err != nil {
		logger.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}

	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}

func (this *SysPrivilegeRoleController) Delete() {
	var rows []entity.PrivilegeRole
	err := utils.ParseForm(this.Ctx.Request, &rows, "JSON")
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}

	err = this.privilegeRole.Delete(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, i18n.Success(this.Ctx.Request))
}
