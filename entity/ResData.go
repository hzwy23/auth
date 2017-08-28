package entity

type ResData struct {
	ResId       string `json:"res_id"`
	ResName     string `json:"res_name"`
	ResAttr     string `json:"res_attr"`
	ResAttrDesc string `json:"res_attr_desc"`
	ResUpid     string `json:"res_up_id"`
	Restype     string `json:"res_type"`
	ResTypeDesc string `json:"res_type_desc"`
	SysFlag     string `json:"sys_flag"`
	InnerFlag   string `json:"inner_flag"`
	ServiceCd   string `json:"service_cd"`
}
