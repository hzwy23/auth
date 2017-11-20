package service

import (
	"html/template"
	"io/ioutil"

	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/panda"
	"github.com/hzwy23/panda/jwt"
	"net/http"
)

var roleResourceModel = models.RoleAndResourceModel{}

func ParseText(r *http.Request, content string) (*template.Template, error) {
	claim, err := jwt.ParseHttp(r)
	if err != nil {
		return nil, err
	}
	return template.New("template").Funcs(template.FuncMap{"checkResIDAuth": func(args ...string) bool {
		if len(args) < 2 {
			return false
		}
		if panda.IsAdmin(claim.UserId) {
			return true
		}
		if args[0] == "2" {
			return roleResourceModel.CheckResIDAuth(claim.UserId, args[1])
		} else {
			return false
		}
	}}).Parse(content)
}

func ParseFile(r *http.Request, filePath string) (*template.Template, error) {
	rst, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return ParseText(r, string(rst))
}
