package service

import (
	"html/template"
	"io/ioutil"

	"github.com/hzwy23/auth-core/models"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/router"
)

var roleResourceModel = models.RoleAndResourceModel{}

func ParseText(ctx router.Context, content string) (*template.Template, error) {
	claim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		return nil, err
	}
	return template.New("template").Funcs(template.FuncMap{"checkResIDAuth": func(args ...string) bool {
		if len(args) < 2 {
			return false
		}
		if utils.IsAdmin(claim.UserId) {
			return true
		}
		if args[0] == "2" {
			return roleResourceModel.CheckResIDAuth(claim.UserId, args[1])
		} else {
			return false
		}
	}}).Parse(content)
}

func ParseFile(ctx router.Context, filePath string) (*template.Template, error) {
	rst, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return ParseText(ctx, string(rst))
}
