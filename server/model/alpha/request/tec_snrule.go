package request

type TecSNRuleSearch struct {
	//StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	//EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name       string `json:"name" form:"name" `
	Branch     string `json:"branch" form:"branch"`         //分部
	IsCustomer string `json:"isCustomer" form:"isCustomer"` //是否客户首字母
	Customer   string `json:"customer" form:"customer"`     //客户id
	//request.PageInfo
}
