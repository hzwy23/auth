package entity

type HandleLogs struct {
	Uuid       string `json:"uuid"`
	UserId     string `json:"user_id"`
	HandleTime string `json:"handle_time" dateType:"YYYY-MM-DD HH24:MM:SS"`
	ClientIP   string `json:"client_ip"`
	StatusCode string `json:"status_code"`
	Method     string `json:"method"`
	Url        string `json:"url"`
	Data       string `json:"data"`
}
