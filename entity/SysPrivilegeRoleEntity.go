package entity

type PrivilegeRole struct {
	Uuid          string `json:"uuid"`
	PrivilegeId   string `json:"privilegeId"`
	PrivilegeDesc string `json:"privilegeDesc"`
	RoleId        string `json:"roleId"`
	RoleName      string `json:"roleName"`
	CreateUser    string `json:"createUser"`
	CreateTime    string `json:"createTime"`
	ModifyUser    string `json:"modifyUser"`
	ModifyTime    string `json:"modifyTime"`
}
