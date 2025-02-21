/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package sql

import (
	"log"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	var obj Database
	obj.Username = "postgres"
	obj.Password = "hangover"
	obj.Ip = "localhost"
	obj.Port = "5432"
	obj.Schema = "BillingEngine"
	obj.Type = "postgres"

	err := obj.Connect()
	if err != nil {
		t.Error(err)
		return
	}
	log.Println("Connect Done")
	time.Sleep(10 * time.Second)
	row, err := obj.Query("select * from logs.transactions_copy")
	defer row.Close()
	if err != nil {
		t.Error(err)
		return

	}
	log.Println("Row Done")
	time.Sleep(10 * time.Second)
	_, data, err := Scan(row)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println("Scan Done")
	time.Sleep(20 * time.Second)
	log.Println(data[0][0])
	time.Sleep(20 * time.Second)
}
