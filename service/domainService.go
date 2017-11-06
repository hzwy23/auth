package service

import (
	"net/http"
	"strings"

	"github.com/hzwy23/auth-core/models"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/validator"
)

var DomainService = &DomainServiceImpl{}

type DomainServiceImpl struct {
	domain models.UserDomainPrivilege
}

func (this *DomainServiceImpl) GetAuthLevel(userId string, domainId string) int {
	row, err := this.domain.Get(userId, domainId)
	if err != nil {
		logger.Error("检验用户对域的访问权限失败", err)
		return -1
	}
	if row.RoleStatusId != 0 {
		logger.Error("角色【", row.RoleId, "】已经被禁止使用")
		return -1
	}
	return row.Permission
}

// 检查用户对指定的域的权限
// 第一个参数中,http.Request,包含了用户的连接信息,cookie中.
// 第二个参数中,domain_id,是用户想要访问的域
// 第三个参数是访问模式,r 表示 只读, w 表示 读写.
// 如果返回true,表示用户有权限
// 返回false,表示用户没有权限
func DomainAuth(req *http.Request, domain_id string, pattern string) bool {
	if validator.IsEmpty(domain_id) {
		return false
	}
	level := checkDomainAuthLevel(req, domain_id)
	switch strings.ToLower(pattern) {
	case "r":
		return level != -1
	case "w":
		return level == 2
	default:
		return false
	}
}

func checkDomainAuthLevel(req *http.Request, domain_id string) int {
	level := -1
	jclaim, err := jwt.GetJwtClaims(req)
	if err != nil {
		logger.Error(err)
		return level
	}
	// if the user is not admin, and user_id is not owner this domain_id
	// check share info. or not
	if utils.IsAdmin(jclaim.UserId) {
		return 2
	} else {
		return DomainService.GetAuthLevel(jclaim.UserId, domain_id)
	}
}
