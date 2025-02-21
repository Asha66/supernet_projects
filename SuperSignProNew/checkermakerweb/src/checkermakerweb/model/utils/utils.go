/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/astaxie/beego/context"

	log "github.com/sirupsen/logrus"

	"github.com/astaxie/beego"

	"checkermakerweb/utils/encoding/base64"
	p "checkermakerweb/utils/util/password"
	"checkermakerweb/utils/util/pbkdf2"
	//"/*proyava.com/encoding/base64"
	//"proyava.com/util/log"
	//p "proyava.com/util/password"
	//"proyava.com/util/pbkdf2"/
)

type MenusStruct struct {
	Menus []string
}

func GeneratePassword() (login_pass, encrypted_pwd string, err error) {
	login_pass, _ = p.Numeric(6)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Login Password", login_pass)

	b := make([]byte, 32)
	_, err = rand.Read(b)
	var pbkdf pbkdf2.Pbkdf2
	pbkdf.Itr = 32
	pbkdf.KeyLen = 32
	pbkdf.Plain = []byte(login_pass)
	pbkdf.Salt = b
	err = pbkdf.Encrypt()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to create password")
		return
	}
	var tmp []byte
	tmp = append(tmp, pbkdf.Salt...)
	tmp = append(tmp, pbkdf.Cipher...)

	out, err := base64.Encode(tmp)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to create password")
		return
	}
	encrypted_pwd = string(out)
	return
}

func IsAuthorized(role, page string) (result bool, err error) {

	result = false

	var menusjson string

	menusjson = ""

	if role == "admin@gmail.com" {
		menusjson = beego.AppConfig.String("ADMIN_MENU_ARRAY")
	} else if role == "admin1@gmail.com" {
		menusjson = beego.AppConfig.String("ADMIN_MENU_ARRAY")
	} else if role == "REPORT" {
		menusjson = beego.AppConfig.String("REPORT_MENU_ARRAY")
	}

	if role == "" {
		err = errors.New("Role invalid")
		return
	}

	var menus MenusStruct
	err = json.Unmarshal([]byte(menusjson), &menus)
	if err != nil {
		return
	}

	for _, men := range menus.Menus {
		if strings.EqualFold(men, page) {
			result = true
		}
	}
	return
}

func SendEmail(host string, port int, userName string, password string, to []string, subject string, message string) (err error) {
	parameters := struct {
		From    string
		To      string
		Subject string
		Message string
	}{
		userName,
		strings.Join([]string(to), ","),
		subject,
		message,
	}

	buffer := new(bytes.Buffer)

	template := template.Must(template.New("emailTemplate").Parse(emailScript()))
	template.Execute(buffer, &parameters)

	auth := smtp.PlainAuth("", userName, password, host)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		userName,
		to,
		buffer.Bytes())

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	}
	return err
}
func emailScript() (script string) {
	return `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"

{{.Message}}`
}

func SetHTTPHeader(Ctx *context.Context) {
	Ctx.Output.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	Ctx.Output.Header("Pragma", "no-cache")
	Ctx.Output.Header("Expires", "0")
	Ctx.Output.Header("X-Content-Type-Options", "nosniff")
	Ctx.Output.Header("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	Ctx.Output.Header("X-Frame-Options", "SAMEORIGIN")
	Ctx.Output.Header("X-XSS-Protection", "1; mode=block")
	Ctx.Output.Header("X-Content-Security-Policy", "default-src 'self'")
	//Ctx.Output.Header("X-WebKit-CSP", "default-src 'self'")
}
