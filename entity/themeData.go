package entity

type ThemeData struct {
	ThemeId     string `json:"theme_id"`
	ThemeDesc   string `json:"theme_desc"`
	ResId       string `json:"res_id"`
	ResUrl      string `json:"res_url"`
	ResOpenType string `json:"res_open_type"`
	ResBgColor  string `json:"res_bg_color"`
	ResClass    string `json:"res_class"`
	GroupId     string `json:"group_id"`
	ResImg      string `json:"res_img"`
	SortId      string `json:"sort_id"`
	NewIframe   string `json:"new_iframe"`
}
