package config

type WxGlobanConfig struct {
	AppID       string // 微信AppID
	Appsecret   string // 微信secret
	AccessToken string // 会自动刷新计算 不用管
	Token       string // 自定义的安全码
	AuthorityID uint   // 微信扫码注册用户的ID
}
