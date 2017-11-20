package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/auth/entity"
)

type UserThemeModel struct {
}

func (this UserThemeModel) Get(user_id string) (entity.UserThemeData, error) {
	var rst entity.UserThemeData
	err := dbobj.QueryForStruct(sys_rdbms_102, &rst, user_id)
	return rst, err
}

func (this UserThemeModel) Put(user_id, theme_id string) error {
	_, err := dbobj.Exec(sys_rdbms_024, theme_id, user_id)
	return err
}
