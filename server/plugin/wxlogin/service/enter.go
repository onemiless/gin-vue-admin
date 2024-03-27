package service

type ServiceGroup struct {
	WxLoginService
}

var ServiceGroupApp = new(ServiceGroup)
