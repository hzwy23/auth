package dto

type AuthDto struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Duration string `json:"-"`
	RetCode  int    `json:"retCode"`
	RetMsg   string `json:"retMsg"`
}
