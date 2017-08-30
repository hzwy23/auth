package service

import "github.com/asofdate/auth-core/models"

var RouteService = &RouteServiceImpl{}

type RouteServiceImpl struct {
	role models.RoleAndResourceModel
}

func (this *RouteServiceImpl) CheckUrlAuth(userId string, url string) bool {
	return this.role.CheckUrlAuth(userId, url)
}

func (this *RouteServiceImpl) CheckResIDAuth(userId string, resId string) bool {
	return this.role.CheckResIDAuth(userId, resId)
}
