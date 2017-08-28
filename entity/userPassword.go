package entity

type SysUserSec struct {
	UserId     string `json:"user_id"`
	UserPasswd string `json:"user_passwd"`
	UserStatus int    `json:"user_status"`
	ErrorCnt   int
}
