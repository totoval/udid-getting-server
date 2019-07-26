package controllers

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/totoval/framework/config"
	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/framework/request"
	"net/http"
	"strings"

	"github.com/totoval/framework/http/controller"
	"howett.net/plist"
)

type Udid struct {
	controller.BaseController
}

func (u *Udid) MobileConfig(c *request.Context) {
	uu, err := uuid.NewGen().NewV1()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": "uuid generate error"})
		return
	}

	type payloadContent struct {
		URL string `plist:"URL"`
		DeviceAttributes []string `plist:"DeviceAttributes"`
	}
	type mobileConfig struct {
		PayloadContent payloadContent `plist:"PayloadContent"`
		PayloadOrganization string `plist:"PayloadOrganization"`
		PayloadDisplayName string `plist:"PayloadDisplayName"`
		PayloadVersion int `plist:"PayloadVersion"`
		PayloadUUID string `plist:"PayloadUUID"`
		PayloadIdentifier string `plist:"PayloadIdentifier"`
		PayloadDescription string `plist:"PayloadDescription"`
		PayloadType string `plist:"PayloadType"`
	}

	mc := mobileConfig{
		PayloadContent: payloadContent{
			URL: "https://"+config.GetString("app.domain")+"/v1/udid/retrive",
			DeviceAttributes: []string{"UDID", "IMEI", "ICCID", "VERSION", "PRODUCT"},
		},
		PayloadOrganization: "totoval.com",
		PayloadDisplayName: "Getting Device UDID (查询设备UDID)",
		PayloadVersion: 1,
		PayloadUUID: uu.String(),
		PayloadIdentifier: "totoval.com.profile-service",
		PayloadDescription: "This file is only using for retriving device ID (本文件仅用来获取设备ID)",
		PayloadType: "Profile Service",
	}

	c.Writer.Header().Set("Content-Type", "application/x-apple-aspen-config")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=\"totoval-udid-getting-server.mobileconfig\"")

	p, err := plist.MarshalIndent(mc, plist.XMLFormat, "\t")
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, string(p))
	return
}

func (u *Udid) Retrive (c *request.Context){
	requestData, err := c.GetRawData()
	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": "uuid retrive error"})
		return
	}

	type retriveData struct {
		IMEI string `plist:"IMEI"`
		PRODUCT string `plist:"PRODUCT"`
		UDID string `plist:"UDID"`
		VERSION string `plist:"VERSION"`
	}

	var rd retriveData
	start := strings.Index(string(requestData[:]), "<?xml")
	end := strings.Index(string(requestData[:]), "</plist>")
	if start >= end {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": errors.New("cannot locate UDID field").Error()})
		return
	}
	end = end + len("</plist>")

	if _, err := plist.Unmarshal(requestData[start:end], &rd); err != nil{
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "https://"+config.GetString("app.domain")+"/v1/udid/show/"+rd.UDID)
	return
}

func (u *Udid) Show(c *request.Context){
	udid := c.Param("udid")
	c.JSON(http.StatusOK, toto.V{"data": toto.V{
		"udid": udid,
	}})
	return
}
