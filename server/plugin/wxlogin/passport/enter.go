package passport

type PassportGroup struct {
	WxLoginPassport
}

var PassportGroupApp = new(WxLoginPassport)
