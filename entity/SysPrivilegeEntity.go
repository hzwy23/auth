package entity

type SysPrivilegeEntity struct {
	PrivilegeId   string `json:"privilegeId"`
	PrivilegeDesc string `json:"privilegeDesc"`
	CreateUser    string `json:"createUser"`
	CreateTime    string `json:"createTime"`
	ModifyUser    string `json:"modifyUser"`
	ModifyTime    string `json:"modifyTime"`
}
