/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package db

import (
	"log"

	"checkermakerweb/utils/database/sql"

	"github.com/astaxie/beego"
)

var Db sql.Database

func Init() (err error) {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Trying to connect DB")
	Db.Ip = beego.AppConfig.String("DBIP")
	Db.Port = beego.AppConfig.String("DBPort")
	Db.Type = beego.AppConfig.String("DBType")
	Db.Schema = beego.AppConfig.String("DBName")
	Db.Username = beego.AppConfig.String("DBUsername")
	Db.Password = beego.AppConfig.String("DBPassword")
	Db.LogLevel = beego.AppConfig.String("loglevel")

	err = Db.Connect()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", "DB Connect fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Info", "DB Connected successfully")
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return
}
