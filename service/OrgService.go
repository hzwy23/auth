package service

import "github.com/asofdate/auth-core/models"

var OrgService = &OrgServiceImpl{}

type OrgServiceImpl struct {
	user models.UserModel
}

func (this *OrgServiceImpl) GetOrgUnitId(userId string) (string, error) {
	return this.user.GetOrgId(userId)
}
