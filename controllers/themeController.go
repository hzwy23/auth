package controllers

import (
	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/models"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
	"github.com/hzwy23/utils/validator"
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
func (this *themeController) Post(ctx router.Context) {
	ctx.Request.ParseForm()
	theme_id := ctx.Request.FormValue("theme_id")

	// get user connection info from cookes.
	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// 提交更新数据库请求.
	// 更新当前连接用户的主题信息
	err = this.muser.Put(jclaim.UserId, theme_id)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_update"), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
func (this themeController) Put(ctx router.Context) {

	var row entity.ThemeData
	err := utils.ParseForm(ctx.Request, &row)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 423, "参数解析失败，请联系管理员")
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
				hret.Error(ctx.ResponseWriter, 421, err.Error())
				return
			}
			hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
			return
		} else if flag > 0 {
			// 更新主题信息
			err := this.mres.Update(row)
			if err != nil {
				logger.Error(err)
				hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_update"), err)
				return
			}
			hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
			return
		} else {
			hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_update"))
			return
		}
	} else {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_virtual_forbid"))
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
func (this themeController) QueryTheme(ctx router.Context) {
	ctx.Request.ParseForm()
	form := ctx.Request.Form

	res_id := form.Get("res_id")
	theme_id := form.Get("theme_id")

	rst, err := this.mres.GetDetails(res_id, theme_id)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_resource_query_theme"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}
