package entity

type DomainData struct {
	DomainId     string `json:"domainId"`
	DomainDesc   string `json:"domainDesc"`
	DomainStatus string `json:"domainStatus"`
	ModifyUser   string `json:"modifyUser"`
}
