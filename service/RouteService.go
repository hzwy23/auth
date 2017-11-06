package service

import "github.com/hzwy23/auth-core/models"

var RouteService = &RouteServiceImpl{}

type RouteServiceImpl struct {
	role models.RoleAndResourceModel
}

func (this *RouteServiceImpl) CheckUrlAuth(userId string, url, method string) bool {
	return this.role.CheckUrlAuth(userId, url, method)
}

func (this *RouteServiceImpl) CheckResIDAuth(userId string, resId string) bool {
	return this.role.CheckResIDAuth(userId, resId)
}
