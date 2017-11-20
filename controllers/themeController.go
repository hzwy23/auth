package controllers

import (
	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/panda"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/jwt"
	"github.com/hzwy23/panda/logger"
	"github.com/hzwy23/panda/validator"
	"net/http"
)

type themeController struct {
	muser *models.UserThemeModel
	mres  *models.ThemeResourceModel
}

var ThemeCtl = &themeController{
	new(models.UserThemeModel),
	new(models.ThemeResourceModel),
}

// swagger:operation POST /v1/auth/theme/update themeController themeController
//
// 更新用户主题信息
//
// 更新用户主题信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: theme_id
//   in: query
//   description: theme code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this *themeController) Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	theme_id := r.FormValue("theme_id")

	// get user connection info from cookes.
	jclaim, err := jwt.ParseHttp(r)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 403, i18n.Disconnect(r))
		return
	}

	// 提交更新数据库请求.
	// 更新当前连接用户的主题信息
	err = this.muser.Put(jclaim.UserId, theme_id)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_theme_update"), err)
		return
	}
	hret.Success(w, i18n.Success(r))
}

// swagger:operation PUT /v1/auth/resource/config/theme themeController themeController
//
// 更新某个资源的主题信息
//
// 更新菜单资源的主题信息,如果这个资源没有主题信息,则新增,否则更新.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: success
func (this themeController) Put(w http.ResponseWriter, r *http.Request) {
	var row entity.ThemeData
	err := panda.ParseForm(r, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 423, "参数解析失败，请联系管理员")
		return
	}
	if validator.IsNull(row.SortId) {
		row.SortId = "0"
	}

	flag, res_type := this.mres.CheckThemeExists(row.ThemeId, row.ResId)
	if validator.IsIn(res_type, "0", "1", "2") {
		if flag == 0 {
			// 没有这个主题的配置信息,新增主题信息
			err := this.mres.Post(row)
			if err != nil {
				hret.Error(w, 421, err.Error())
				return
			}
			hret.Success(w, i18n.Success(r))
			return
		} else if flag > 0 {
			// 更新主题信息
			err := this.mres.Update(row)
			if err != nil {
				logger.Error(err)
				hret.Error(w, 421, i18n.Get(r, "error_theme_update"), err)
				return
			}
			hret.Success(w, i18n.Success(r))
			return
		} else {
			hret.Error(w, 421, i18n.Get(r, "error_theme_update"))
			return
		}
	} else {
		hret.Error(w, 421, i18n.Get(r, "error_theme_virtual_forbid"))
		return
	}
}

// swagger:operation GET /v1/auth/resource/queryTheme themeController themeController
//
// 查询主题信息
//
// 查询某个菜单资源,某个主题的详细信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this themeController) QueryTheme(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form

	res_id := form.Get("res_id")
	theme_id := form.Get("theme_id")

	rst, err := this.mres.GetDetails(res_id, theme_id)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 419, i18n.Get(r, "error_resource_query_theme"), err)
		return
	}
	hret.Json(w, rst)
}
