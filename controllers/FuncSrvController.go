package controllers

import (
	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/logger"
	"net/http"
	"github.com/hzwy23/panda"
)

type FuncSrvController struct {
	funcRoute models.FuncRoute
	}

func GetServiceManagePage(w http.ResponseWriter,r *http.Request) {
	rst, err := groupcache.GetStaticFile("ServiceManage")
	if err != nil {
		hret.Error(w, 404, i18n.Get(r, "as_of_date_page_not_exist"))
		return
	}

	hz, err := service.ParseText(r, string(rst))
	if err != nil {
		logger.Error(err)
		hret.Error(w, 404, i18n.Get(r, "as_of_date_page_not_exist"))
		return
	}
	hz.Execute(w, nil)
}

// 查询功能服务
// @param resId   资源编码
func (this *FuncSrvController) Get(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	resId := r.FormValue("resId")

	rst, err := this.funcRoute.Get(resId)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, "查询失败，请联系管理员")
		return
	}
	hret.Json(w, rst)
}

// 删除功能服务配置信息
func (this *FuncSrvController) Delete(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	//js := this.Ctx.Request.FormValue("JSON")
	var rows []models.FuncRoute
	panda.ParseForm(r,&rows)

	//err := json.Unmarshal([]byte(js), &rows)
	//if err != nil {
	//	logger.Error(err)
	//	hret.Error(this.Ctx.ResponseWriter, 421, "解析json数据失败，请联系管理员")
	//	return
	//}
	err := this.funcRoute.Delete(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 422, err.Error())
		return
	}
	hret.Success(w, "success")
}

// 更新功能服务配置信息
func (this *FuncSrvController) Put(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	form := r.Form
	var row models.FuncRoute
	row.ResId = form.Get("res_id")
	row.ResName = form.Get("res_name")
	row.ResUpId = form.Get("res_up_id")
	row.ResUrl = form.Get("res_url")
	row.Method = form.Get("method")
	row.ResOpenType = form.Get("res_type")
	row.NewIframe = form.Get("new_iframe")
	row.Uuid = form.Get("uuid")
	row.ThemeId = "funcs"

	var err error
	if this.funcRoute.IsExists(row.ResId) {
		err = this.funcRoute.Update(row)
	} else {
		err = this.funcRoute.AddTheme(row)
	}
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, err.Error())
		return
	}
	hret.Success(w, "success")
}

// 新建功能服务配置信息
func (this *FuncSrvController) Post(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	form := r.Form
	var row models.FuncRoute
	row.ResId = form.Get("res_id")
	row.ResName = form.Get("res_name")
	row.ResUpId = form.Get("res_up_id")
	row.ResUrl = form.Get("res_url")
	row.Method = form.Get("method")
	row.ResOpenType = form.Get("res_type")
	row.NewIframe = form.Get("new_iframe")
	row.ThemeId = "funcs"

	err := this.funcRoute.Post(row)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, err.Error())
		return
	}
	hret.Success(w, "success")
}

func init() {
	groupcache.RegisterStaticFile("ServiceManage", "./views/hauth/service.tpl")
}
