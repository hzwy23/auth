package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/uuid"
	"github.com/hzwy23/auth-core/entity"
)

type ThemeResourceModel struct {
}

// 查询主题对应的资源配置信息
func (this *ThemeResourceModel) Get(theme_id string) ([]entity.ThemeData, error) {
	var rst []entity.ThemeData
	rows, err := dbobj.Query(sys_rdbms_101, theme_id)
	if err != nil {
		logger.Error(err)
		return rst, err
	}
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 查询指定资源,指定主题的详细信息
func (this *ThemeResourceModel) GetDetails(res_id string, theme_id string) ([]entity.ThemeData, error) {
	rows, err := dbobj.Query(sys_rdbms_070, theme_id, res_id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var rst []entity.ThemeData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 更新资源主题信息
func (this *ThemeResourceModel) Update(row entity.ThemeData) error {
	_, err := dbobj.Exec(sys_rdbms_009,
		row.ResUrl, row.ResBgColor, row.ResClass, row.ResImg, row.GroupId,
		row.SortId, row.ResOpenType, row.NewIframe, row.ThemeId, row.ResId)
	return err
}

// 新建资源主题信息
func (this *ThemeResourceModel) Post(row entity.ThemeData) error {

	_, err := dbobj.Exec(sys_rdbms_008, uuid.GenUUID(),
		row.ThemeId, row.ResId, row.ResUrl, row.ResOpenType,
		row.ResBgColor, row.ResClass, row.GroupId,
		row.ResImg, row.SortId, row.NewIframe)
	return err
}

// 检查资源是否已经配置主题信息
func (this *ThemeResourceModel) CheckThemeExists(theme_id string, res_id string) (int, string) {
	cnt := -1
	err := dbobj.QueryRow(sys_rdbms_006, theme_id, res_id).Scan(&cnt)
	if err != nil {
		return -1, ""
	}
	res_type := "4"
	err = dbobj.QueryRow(sys_rdbms_013, res_id).Scan(&res_type)
	if err != nil {
		return -1, ""
	}
	return cnt, res_type
}
