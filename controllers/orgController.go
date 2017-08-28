package controllers

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/asofdate/auth-core/entity"
	"github.com/asofdate/auth-core/groupcache"
	"github.com/asofdate/auth-core/models"
	"github.com/asofdate/auth-core/service"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
	"github.com/tealeg/xlsx"
)

type orgController struct {
	models *models.OrgModel
	upload chan int
}

var OrgCtl = &orgController{
	models: new(models.OrgModel),
	upload: make(chan int, 1),
}

// swagger:operation GET /v1/auth/resource/org/page StaticFiles orgController
//
// 机构信息配置管理页面
//
// 首先系统会检查用户的连接信息,如果用户被授权访问这个页面,将会返回机构配置管理页面内容,否则返回响应的错误住状态.
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
func (orgController) Page(ctx router.Context) {
	ctx.Request.ParseForm()
	if !service.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := groupcache.GetStaticFile("AsofdateOrgPage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}
	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/resource/org/get orgController orgController
//
// 查询机构信息
//
// API将会返回指定域中的机构信息,用户在请求这个API时,需要传入domain_id这个字段值
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
func (this orgController) Get(ctx router.Context) {
	ctx.Request.ParseForm()

	rst, err := this.models.Get()
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 417, i18n.Get(ctx.Request, "error_query_org_info"))
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/resource/org/delete orgController orgController
//
// 删除机构信息
//
// 首先系统会校验用户的权限,如果用户拥有删除机构的权限,系统将会根据用户请求的参数,删除响应的机构信息.
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
// - name: JSON
//   in: query
//   description: json format
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) Delete(ctx router.Context) {
	ctx.Request.ParseForm()

	orgUnitId := ctx.Request.FormValue("orgUnitId")

	msg, err := this.models.Delete(orgUnitId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 418, i18n.Get(ctx.Request, msg), err)
		return
	}

	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation PUT /v1/auth/resource/org/update orgController orgController
//
// 更新机构信息
//
// 系统将会更具用户传入的参数,修改指定机构信息.
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
func (this orgController) Update(ctx router.Context) {
	ctx.Request.ParseForm()
	var arg entity.SysOrgInfo
	err := utils.ParseForm(ctx.Request, &arg)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, err.Error())
		return
	}

	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Update(arg, jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Get(ctx.Request, "success"))
}

// swagger:operation POST /v1/auth/resource/org/post orgController orgController
//
// 新增机构信息
//
// 想指定域中新增机构信息
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
// - name: Org_unit_id
//   in: query
//   description: org code number
//   required: true
//   type: string
//   format:
// - name: Org_unit_desc
//   in: query
//   description: org desc
//   required: true
//   type: string
//   format:
// - name: Up_org_id
//   in: query
//   description: up org id
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) Post(ctx router.Context) {
	ctx.Request.ParseForm()
	var arg entity.SysOrgInfo
	err := utils.ParseForm(ctx.Request, &arg)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 423, err.Error())
		return
	}
	jclaim, err := jwt.GetJwtClaims(ctx.Request)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Post(arg, jclaim.UserId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Get(ctx.Request, "success"))
}

// swagger:operation GET /v1/auth/relation/domain/org orgController orgController
//
// 返回某个机构的所有下级机构信息
//
// 根据客户端请求时指定的机构id,获取这个id所有的下属机构信息
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
// - name: org_unit_id
//   in: query
//   description: org code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) GetSubOrgInfo(ctx router.Context) {
	ctx.Request.ParseForm()

	org_unit_id := ctx.Request.FormValue("org_unit_id")

	rst, err := this.models.GetSubOrgInfo(org_unit_id)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_org_sub_query"))
		return
	}

	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/resource/org/download orgController orgController
//
// 下载机构信息
//
// 下载某个指定域的所有机构信息. 只能下载用户有权限访问的域中的机构
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
func (this orgController) Download(ctx router.Context) {
	ctx.Request.ParseForm()
	if !service.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/vnd.ms-excel")

	rst, err := this.models.Get()
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 417, i18n.Get(ctx.Request, "error_query_org_info"))
		return
	}

	var sheet *xlsx.Sheet
	file, err := xlsx.OpenFile(filepath.Join(os.Getenv("HBIGDATA_HOME"), "views", "uploadTemplate", "hauthOrgExportTemplate.xlsx"))
	if err != nil {
		file = xlsx.NewFile()
		sheet, err = file.AddSheet("机构信息")
		if err != nil {
			logger.Error(err)
			hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_org_sheet"))
			return
		}

		{
			row := sheet.AddRow()
			cell1 := row.AddCell()
			cell1.Value = "机构编码"
			cell2 := row.AddCell()
			cell2.Value = "机构名称"
			cell3 := row.AddCell()
			cell3.Value = "上级编码"
			cell9 := row.AddCell()
			cell9.Value = "所属域"

			cell5 := row.AddCell()
			cell5.Value = "创建日期"
			cell6 := row.AddCell()
			cell6.Value = "创建人"
			cell7 := row.AddCell()
			cell7.Value = "维护日期"
			cell8 := row.AddCell()
			cell8.Value = "维护人"

		}
	} else {
		sheet = file.Sheet["机构信息"]
		if sheet == nil {
			hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_org_sheet"))
			return
		}
	}
	for _, v := range rst {
		row := sheet.AddRow()
		cell1 := row.AddCell()
		cell1.Value = v.OrgUnitId
		cell1.SetStyle(sheet.Rows[1].Cells[0].GetStyle())

		cell2 := row.AddCell()
		cell2.Value = v.OrgUnitDesc
		cell2.SetStyle(sheet.Rows[1].Cells[1].GetStyle())

		cell3 := row.AddCell()
		cell3.Value, _ = utils.SplitCode(v.UpOrgId)
		cell3.SetStyle(sheet.Rows[1].Cells[2].GetStyle())

		cell5 := row.AddCell()
		cell5.Value = v.CreateDate
		cell5.SetStyle(sheet.Rows[1].Cells[4].GetStyle())

		cell6 := row.AddCell()
		cell6.Value = v.CreateUser
		cell6.SetStyle(sheet.Rows[1].Cells[5].GetStyle())

		cell7 := row.AddCell()
		cell7.Value = v.MaintanceDate
		cell7.SetStyle(sheet.Rows[1].Cells[6].GetStyle())

		cell8 := row.AddCell()
		cell8.Value = v.MaintanceUser
		cell8.SetStyle(sheet.Rows[1].Cells[7].GetStyle())

	}

	if len(sheet.Rows) >= 3 {
		sheet.Rows = append(sheet.Rows[0:1], sheet.Rows[2:]...)
	}

	file.Write(ctx.ResponseWriter)
}

// swagger:operation GET /v1/auth/resource/org/upload orgController orgController
//
// 上传机构信息
//
// 根据客户端导入的excel格式的数据,将机构信息写入到数据库中.
//
// 这个上传过程是:增量删除, 一旦出现重复的机构,将会中断上传过程,且数据库会立刻回滚.
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
func (this orgController) Upload(ctx router.Context) {
	if len(this.upload) != 0 {
		hret.Success(ctx.ResponseWriter, i18n.Get(ctx.Request, "error_org_upload_wait"))
		return
	}

	// 从cookies中获取用户连接信息
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// 同一个时间,只能有一个导入任务
	this.upload <- 1
	defer func() { <-this.upload }()

	ctx.Request.ParseForm()
	fd, _, err := ctx.Request.FormFile("file")
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_org_read_upload_file"))
		return
	}

	result, err := ioutil.ReadAll(fd)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_org_read_upload_file"))
		return
	}

	// 读取上传过来的文件信息
	// 转换成二进制数据流
	file, err := xlsx.OpenBinary(result)
	sheet, ok := file.Sheet["机构信息"]
	if !ok {
		logger.Error("没有找到'机构信息'这个sheet页")
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_org_sheet"))
		return
	}

	var data []entity.SysOrgInfo
	for index, val := range sheet.Rows {
		if index > 0 {
			var one entity.SysOrgInfo
			one.OrgUnitId = val.Cells[0].Value
			one.OrgUnitDesc = val.Cells[1].Value
			one.UpOrgId = val.Cells[2].Value
			one.CreateUser = jclaim.UserId

			if one.OrgUnitId == one.UpOrgId {
				hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "as_of_date_up_org_equal_org_id"))
				return
			}

			data = append(data, one)
		}
	}

	msg, err := this.models.Upload(data)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// 查询某个机构的详细信息
func (this orgController) GetDetails(ctx router.Context) {
	ctx.Request.ParseForm()

	orgUnitId := ctx.Request.FormValue("orgUnitId")
	if len(orgUnitId) == 0 {
		logger.Error("机构号为空")
		hret.Error(ctx.ResponseWriter, 421, "机构号为空")
		return
	}
	row, err := this.models.GetDetails(orgUnitId)
	if err != nil {
		logger.Error(err)
		hret.Error(ctx.ResponseWriter, 423, err.Error())
		return
	}
	hret.Json(ctx.ResponseWriter, row)
}

func init() {
	groupcache.RegisterStaticFile("AsofdateOrgPage", "./views/hauth/org_page.tpl")
}
