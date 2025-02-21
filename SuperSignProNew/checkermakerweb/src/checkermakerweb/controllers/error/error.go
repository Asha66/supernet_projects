/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package error

import (
	"checkermakerweb/model/utils"
	"checkermakerweb/session"
	"log"

	"github.com/astaxie/beego"
)

type Error struct {
	beego.Controller
}

func (c *Error) Error404() {
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "404 Path", c.Ctx.Request.URL.Path)

	c.TplName = "error/error.html"
}

func (c *Error) Error501() {
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)
	utils.SetHTTPHeader(c.Ctx)
	sess, _ := session.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)

	log.Println(beego.AppConfig.String("loglevel"), "Info", "UserName Nil Found")
	uname := sess.Get("uname")
	if uname != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Info", "UserName Nil Found")
		session.SeTLogoutSession(uname.(string))
	}
	sess.SessionRelease(c.Ctx.ResponseWriter)
	session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
	c.Data["DisplayMessage"] = "501, server error"
	c.TplName = "error/error.html"
}
