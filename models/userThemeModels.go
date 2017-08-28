package models

import (
	"github.com/asofdate/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
)

type UserThemeModel struct {
}

func (this UserThemeModel) Get(user_id string) ([]entity.UserThemeData, error) {
	var rst []entity.UserThemeData
	rows, err := dbobj.Query(sys_rdbms_102, user_id)
	if err != nil {
		logger.Error(err)
		return rst, err
	}
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (this UserThemeModel) Put(user_id, theme_id string) error {
	_, err := dbobj.Exec(sys_rdbms_024, theme_id, user_id)
	return err
}
