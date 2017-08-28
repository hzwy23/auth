package entity

type PrivilegeDomain struct {
	Uuid          string `json:"uuid"`
	PrivilegeId   string `json:"privilegeId"`
	PrivilegeDesc string `json:"privilegeDesc"`
	DomainId      string `json:"domainId"`
	DomainName    string `json:"domainName"`
	Permission    string `json:"permission"`
	CreateUser    string `json:"createUser"`
	CreateTime    string `json:"createTime"`
	ModifyUser    string `json:"modifyUser"`
	ModifyTime    string `json:"modifyTime"`
}
