package entity

type RoleInfo struct {
	RoleName          string `json:"role_name"`
	RoleOwner         string `json:"create_user"`
	RoleCreateDate    string `json:"create_date"`
	RoleStatusDesc    string `json:"role_status_desc"`
	RoleStatus        string `json:"role_status_code"`
	RoleMaintanceDate string `json:"modify_date"`
	RoleMaintanceUser string `json:"modify_user"`
	RoleId            string `json:"role_id"`
}
