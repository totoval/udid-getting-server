package groups

import (
	"github.com/totoval/framework/route"
	"totoval/app/http/controllers"
)

type UdidGroup struct {
	UdidController    controllers.Udid
}

func (ug *UdidGroup) Group(group route.Grouper) {
	group.GET("/mobileconfig", ug.UdidController.MobileConfig)
	group.POST("/retrive", ug.UdidController.Retrive)
	group.GET("/show/:udid", ug.UdidController.Show)
}
