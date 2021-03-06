package controllers

import (
	"os"
	"path/filepath"

	"github.com/hzwy23/auth/groupcache"
	"github.com/hzwy23/auth/models"
	"github.com/hzwy23/auth/service"
	"github.com/hzwy23/panda/hret"
	"github.com/hzwy23/panda/i18n"
	"github.com/hzwy23/panda/logger"
	"github.com/tealeg/xlsx"
	"net/http"
)

type handleLogsController struct {
	model models.HandleLogMode
}

var HandleLogsCtl = &handleLogsController{}

// swagger:operation GET /v1/auth/HandleLogsPage StaticFiles handleLogsController
//
// 操作日志页面
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
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
//   '404':
//     description: page not found
func (this *handleLogsController) Page(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	rst, err := groupcache.GetStaticFile("AsofdateHandleLogPage")
	if err != nil {
		hret.Error(w, 404, i18n.PageNotFound(r))
		return
	}
	hz, err := service.ParseText(r, string(rst))
	if err != nil {
		hret.Error(w, 404, i18n.PageNotFound(r))
		return
	}
	hz.Execute(w, nil)
}

// swagger:operation GET /v1/auth/handle/logs/download handleLogsController handleLogsController
//
// 下载日志记录,返回excel格式数据
//
// API将会返回用户所属域中的所有操作记录信息.所以,在使用这个API时,必须登录系统.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// - application/vnd.ms-excel
// responses:
//   '200':
//     description: success
//   '403':
//     description: Insufficient permissions
//   '421':
//     description: query logs information failed.
func (this handleLogsController) Download(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	if !service.BasicAuth(r) {
		hret.Error(w, 403, i18n.NoAuth(r))
		return
	}
	w.Header().Set("Content-Type", "application/vnd.ms-excel")

	rst, err := this.model.Download()
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_handle_logs_get_failed"))
		return
	}

	file, err := xlsx.OpenFile(filepath.Join(os.Getenv("HBIGDATA_HOME"), "views", "uploadTemplate", "hauthHandleLogsTemplate.xlsx"))
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_handle_logs_open_error"), err)
		return
	}
	sheet, ok := file.Sheet["handle_logs"]
	if !ok {
		hret.Error(w, 421, i18n.Get(r, "error_handle_logs_sheet_error"))
		return
	}

	for _, v := range rst {
		row := sheet.AddRow()
		cell1 := row.AddCell()
		cell1.Value = v.UserId
		cell1.SetStyle(sheet.Rows[1].Cells[0].GetStyle())

		cell2 := row.AddCell()
		cell2.Value = v.HandleTime
		cell2.SetStyle(sheet.Rows[1].Cells[1].GetStyle())

		cell3 := row.AddCell()
		cell3.Value = v.ClientIP
		cell3.SetStyle(sheet.Rows[1].Cells[2].GetStyle())

		cell4 := row.AddCell()
		cell4.Value = v.Method
		cell4.SetStyle(sheet.Rows[1].Cells[3].GetStyle())

		cell5 := row.AddCell()
		cell5.Value = v.Url
		cell5.SetStyle(sheet.Rows[1].Cells[4].GetStyle())

		cell6 := row.AddCell()
		cell6.Value = v.StatusCode
		cell6.SetStyle(sheet.Rows[1].Cells[5].GetStyle())

		cell7 := row.AddCell()
		cell7.Value = v.Data
		cell7.SetStyle(sheet.Rows[1].Cells[6].GetStyle())
	}

	if len(sheet.Rows) >= 3 {
		sheet.Rows = append(sheet.Rows[0:1], sheet.Rows[2:]...)
	}

	file.Write(w)
}

// swagger:operation GET /v1/auth/handle/logs handleLogsController handleLogsController
//
// 查询用户所属域中的操作日志信息
//
// API只能查询用户所属域的操作日志信息, 数据处理中,采用了分页查询,所以,必须传入2个参数,分别是:
//
// offset: 起始行数
//
// limit : 最大行数
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: offset
//   in: query
//   description: 起始行数,必须是数字.
//   required: true
//   type: integer
//   format:
// - name: limit
//   in: query
//   description: 最大行数,必须是数字.
//   required: true
//   type: integer
//   format:
// responses:
//   '200':
//      description: success
//   '403':
//      description: Insufficient permissions
//   '421':
//      description: query logs information failed.
func (this handleLogsController) GetHandleLogs(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	// Get form data from client request.
	offset := r.FormValue("offset")
	limit := r.FormValue("limit")

	rst, total, err := this.model.Get(offset, limit)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_handle_logs_query_failed"))
		return
	}
	hret.BootstrapTable(w, total, rst)
}

// swagger:operation GET /v1/auth/handle/logs/search handleLogsController handleLogsController
//
// 返回满足用户搜索条件的日志信息
//
// API中会校验用户的权限,如果用户没有登录,将返回权限不足的提示信息
//
// 这个API需要提供3个参数,分别是:
//
// UserId    : 用户账号
//
// StartDate : 日志操作开始日期
//
// EndDate   : 日志操作结束日期
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: UserId
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: StartDate
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: EndDate
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
//   '403':
//     description: Insufficient permissions
//   '421':
//     description: query logs information failed.
func (this handleLogsController) SerachLogs(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	// Get form data from request.
	userid := r.FormValue("UserId")
	start := r.FormValue("StartDate")
	end := r.FormValue("EndDate")
	offset := r.FormValue("offset")
	limit := r.FormValue("limit")

	rst, cnt, err := this.model.Search(userid, start, end, offset, limit)
	if err != nil {
		logger.Error(err)
		hret.Error(w, 421, i18n.Get(r, "error_handle_logs_query_failed"))
		return
	}

	hret.BootstrapTable(w, cnt, rst)
}

func init() {
	groupcache.RegisterStaticFile("AsofdateHandleLogPage", "./views/hauth/handle_logs_page.tpl")
}
