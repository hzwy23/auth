package entity

type UserRoleData struct {
	UserId      string `json:"userId"`
	UserDesc    string `json:"userDesc"`
	OrgUnitId   string `json:"orgUnitId"`
	OrgUnitDesc string `json:"orgUnitDesc"`
	RoleId      string `json:"roleId"`
	CreateUser  string `json:"createUser"`
	CreateTime  string `json:"createTime"`
}
