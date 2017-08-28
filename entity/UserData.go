package entity

type UserInfo struct {
	UserId            string `json:"user_id"`
	UserName          string `json:"user_name"`
	UserStatusDesc    string `json:"status_desc"`
	UserCreateDate    string `json:"create_date"`
	UserOwner         string `json:"create_user"`
	UserEmail         string `json:"user_email"`
	UserPhone         string `json:"user_phone"`
	OrgUnitId         string `json:"org_unit_id"`
	OrgUnitDesc       string `json:"org_unit_desc"`
	UserMaintanceDate string `json:"modify_date"`
	UserMaintanceUser string `json:"modify_user"`
	UserStatusId      string `json:"status_cd"`
}
