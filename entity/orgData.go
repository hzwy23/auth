package entity

type SysOrgInfo struct {
	OrgUnitId     string `json:"org_id"`
	OrgUnitDesc   string `json:"org_desc"`
	UpOrgId       string `json:"up_org_id"`
	CreateDate    string `json:"create_date"`
	MaintanceDate string `json:"modify_date"`
	CreateUser    string `json:"create_user"`
	MaintanceUser string `json:"modify_user"`
	OrgDept       string `json:"org_dept,omitempty"`
}
