package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
)

type FuncRoute struct {
	Uuid        string
	ThemeId     string
	ResId       string
	ResName     string
	ResUrl      string
	ResOpenType string
	ServiceCd   string
	ResUpId     string
	NewIframe   string
}

// 删除功能服务
func (this *FuncRoute) Delete(rows []FuncRoute) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, val := range rows {
		_, err := tx.Exec(sys_rdbms_073, val.ResId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return err
		}
		_, err = tx.Exec(sys_rdbms_076, val.ResId)
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// 更新功能服务
func (this *FuncRoute) Update(row FuncRoute) error {
	innerFlag := "true"
	if len(row.ServiceCd) != 0 {
		innerFlag = "false"
	}

	if len(row.NewIframe) == 0 {
		row.NewIframe = "false"
	}

	tx, err := dbobj.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(sys_rdbms_104,
		row.ResName, row.ServiceCd,
		innerFlag, row.ResId,
	)
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sys_rdbms_105,
		row.ResUrl, row.NewIframe,
		row.ResOpenType, row.ThemeId, row.Uuid)
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// 新增功能服务
func (this *FuncRoute) Post(row FuncRoute) error {
	innerFlag := "true"
	if len(row.ServiceCd) != 0 {
		innerFlag = "false"
	}

	if len(row.NewIframe) == 0 {
		row.NewIframe = "false"
	}

	tx, err := dbobj.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(sys_rdbms_072,
		row.ResId,
		row.ResName, "1",
		row.ResUpId, "2",
		innerFlag,
		row.ServiceCd,
	)
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sys_rdbms_008,
		row.ThemeId, row.ResId, row.ResUrl,
		row.ResOpenType, "", "", "", "",
		0, row.NewIframe)
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (this *FuncRoute) AddTheme(row FuncRoute) error {
	if len(row.NewIframe) == 0 {
		row.NewIframe = "false"
	}

	_, err := dbobj.Exec(sys_rdbms_008,
		row.ThemeId, row.ResId, row.ResUrl,
		row.ResOpenType, "", "", "", "",
		0, row.NewIframe)
	return err
}

func (this *FuncRoute) IsExists(resId string, themeId string) bool {
	cnt := -1
	err := dbobj.QueryRow(sys_rdbms_006, themeId, resId).Scan(&cnt)
	if err != nil {
		return false
	}
	if cnt == 1 {
		return true
	}
	return false
}

// 查询某一个菜单页面下边所有的功能服务
func (this *FuncRoute) Get(resId string, themeId string) ([]FuncRoute, error) {
	var rst []FuncRoute
	err := dbobj.QueryForSlice(sys_rdbms_103, &rst, themeId)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var ret []FuncRoute
	this.dfs(rst, resId, &ret)
	return ret, nil
}

// 获取子资源信息
func (this *FuncRoute) dfs(all []FuncRoute, res_id string, rst *[]FuncRoute) {
	for _, val := range all {
		if val.ResUpId == res_id {
			*rst = append(*rst, val)
			if val.ResId == val.ResUpId {
				logger.Error("层级关系错误,不允许上级菜单域当前菜单编码一致,当前菜单编码:", val.ResId, "上级菜单编码:", val.ResUpId)
				return
			}
			this.dfs(all, val.ResId, rst)
		}
	}
}
