package entity

type HandleLogBuf struct {
	User_id    string `json:"user_id"`
	Client_ip  string `json:"client_ip"`
	Ret_status string `json:"ret_status"`
	Req_method string `json:"req_method"`
	Req_url    string `json:"req_url"`
	Domain_id  string `json:"domain_id"`
	Req_body   string `json:"req_body"`
}
