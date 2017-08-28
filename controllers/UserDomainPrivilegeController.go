package controllers

import (
	"github.com/asofdate/auth-core/models"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
)

type UserDomainPrivilegeController struct {
	UserDomainPrivilegeModel models.UserDomainPrivilege
	router.Controller
}

func (this *UserDomainPrivilegeController) Get() {
	claim, err := jwt.GetJwtClaims(this.Ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	rst, err := this.UserDomainPrivilegeModel.GetByUserId(claim.UserId)
	if err != nil {
		logger.Error(this.Ctx.ResponseWriter, 422, err.Error())
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}
