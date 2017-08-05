package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/asofdate/sso-jwt-auth/utils"
	"github.com/asofdate/sso-jwt-auth/utils/logger"
	"github.com/asofdate/sso-jwt-auth/utils/validator"
	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	*jwt.StandardClaims
	UserId      string
	DomainId    string
	OrgUnitId   string
	Authorities string `json:"authorities"`
	ClientIp    string
}

var (
	key []byte = []byte("hzwy23@163.com-jwt")
)

func GenToken(user_id, domain_id, org_id string, dt int64, clientIp string) string {
	claims := JwtClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + dt,
			Issuer:    "hzwy23",
		},
		UserId:      user_id,
		DomainId:    domain_id,
		OrgUnitId:   org_id,
		Authorities: "ROLE_ADMIN,AUTH_WRITE,ACTUATOR",
		ClientIp:    clientIp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return ss
}

func CheckToken(req *http.Request) bool {
	cookie, err := req.Cookie("Authorization")
	token := ""
	if err != nil || validator.IsEmpty(cookie.Value) {
		token = req.FormValue("Authorization")
	} else {
		token = cookie.Value
	}

	jclaim, err := ParseJwt(token)
	if err != nil {
		logs.Error(err)
		return false
	}
	return jclaim.ClientIp == utils.GetRequestIP(req)
}

func ParseJwt(token string) (*JwtClaims, error) {
	var jclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(token, jclaim, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err, token)
		return nil, errors.New("parase with claims failed.")
	}
	return jclaim, nil
}

func GetJwtClaims(request *http.Request) (*JwtClaims, error) {
	cookie, err := request.Cookie("Authorization")
	if err != nil || cookie == nil || validator.IsEmpty(cookie.Value) {
		jwt := request.Header.Get("Authorization")
		return ParseJwt(jwt)
	}
	return ParseJwt(cookie.Value)
}
