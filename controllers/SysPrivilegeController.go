package controllers

import (
	"time"

	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/groupcache"
	"github.com/asofdate/auth-core/models"
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
)

type SysPrivilegeController struct {
	sysPrivilege models.SysPrivilege
	router.Controller
}

func GetSysPrivilegePage(ctx router.Context) {
	rst, err := groupcache.GetStaticFile("SysPrivilegePage")
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
		return
	}

	hz, err := service.ParseText(ctx, string(rst))
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
		return
	}
	hz.Execute(ctx.ResponseWriter, nil)
}

// 查询权限定义信息
func (this *SysPrivilegeController) Get() {
	rst, err := this.sysPrivilege.Get()
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}

// 新增权限代码
func (this *SysPrivilegeController) Post() {
	var row entity.SysPrivilegeEntity

	err := utils.ParseForm(this.Ctx.Request, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}

	claim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 424, err.Error())
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
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, "success")
}

// 更新权限代码
func (this *SysPrivilegeController) Put() {
	var row entity.SysPrivilegeEntity

	err := utils.ParseForm(this.Ctx.Request, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}

	claim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 424, err.Error())
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
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, "success")
}

// 删除权限代码
func (this *SysPrivilegeController) Delete() {
	var rows []entity.SysPrivilegeEntity
	err := utils.ParseForm(this.Ctx.Request, &rows, "JSON")
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 423, err.Error())
		return
	}
	err = this.sysPrivilege.Delete(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, "success")
}

func init() {
	groupcache.RegisterStaticFile("SysPrivilegePage", "./views/hauth/sysPrivilege.tpl")
}
