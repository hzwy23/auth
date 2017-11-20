package controllers

import (
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
	"github.com/hzwy23/panda/route"
	"net/http"
)

type UserDomainPrivilegeController struct {
	UserDomainPrivilegeModel models.UserDomainPrivilege
	route.Controller
}

func (this *UserDomainPrivilegeController) Get(w http.ResponseWriter, r *http.Request) {
	claim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, err.Error())
		return
	}
	rst, err := this.UserDomainPrivilegeModel.GetByUserId(claim.UserId)
	if err != nil {
		logger.Error(w, 422, err.Error())
		return
	}
	hret.Json(w, rst)
}
