/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package main

import (
	_ "checkermakerweb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var RedirectHttp = func(ctx *context.Context) {
	if !ctx.Input.IsSecure() {
		// no need focToa+VI8Fovi/ECVJRXcpIgldlicRploflkqJz3NqeUeb4l3mH+AtN9Xha+y/R9Br an additional '/' between domain and uri
		url := "https://" + ctx.Input.Domain() + ":" + beego.AppConfig.String("HttpsPort") + ctx.Input.URI()
		ctx.Redirect(302, url)
	}
}

func main() {
	if beego.AppConfig.String("EnableHTTPS") == "true" {
		beego.InsertFilter("/", beego.BeforeRouter, RedirectHttp) // for http://mysite
		beego.InsertFilter("*", beego.BeforeRouter, RedirectHttp) // for http://mysite/*
	}

	beego.Run()
}
