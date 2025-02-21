/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package createStatus

import (
	"checkermakerweb/session"

	//	"io/ioutil"
	//"net/mail"
	"net/smtp"
	"runtime/debug"

	"strings"

	"checkermakerweb/model/db"
	"checkermakerweb/utils"
	"errors"

	//"github.com/scorredoira/email"

	"github.com/astaxie/beego"

	//	"checkermakerweb/services"

	"checkermakerweb/utils/database/sql"

	log "github.com/sirupsen/logrus"

	"fmt"

	//"proyava.com/util/password"

	"github.com/google/uuid"
)

type CreateStatus struct {
	beego.Controller
}

type unencryptedAuth struct {
	smtp.Auth
}

func (c *CreateStatus) Get() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Add Assets Start")
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)
	var err error
	var Autherr error
	sessErr := false
	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
			c.Data["DisplayMessage"] = "Something went wrong.Please Contact CustomerCare."
			c.TplName = "error/error.html"
		}
		if Autherr != nil {
			c.Data["DisplayMessage"] = Autherr.Error()
			c.TplName = "error/error.html"
			return
		}
		if err != nil {
			if sessErr == true {
				log.Println(beego.AppConfig.String("loglevel"), "Info", "Redirecting to login")
				c.Redirect(beego.AppConfig.String("LOGIN_PATH"), 302)

			} else {
				c.Data["DisplayMessage"] = err.Error()
			}
			c.TplName = "systemconfiguration/status/createStatus/createStatus.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Status Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "systemconfiguration/status/createStatus/createStatus.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Status Page Success")
		}
		return
	}()

	sess, err := session.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System is unable to process your request.Please contact customer care")
		sessErr = true
		return
	}

	if err = session.ValidateSession(sess); err != nil {
		sess.SessionRelease(c.Ctx.ResponseWriter)
		session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		sessErr = true
		return
	}
	defer func() {
		utils.EventLogs(c.Ctx, sess, c.Ctx.Input.Method(), c.Input(), c.Data, err)
		sess.SessionRelease(c.Ctx.ResponseWriter)
	}()

	utype := sess.Get("user_type").(string)
	c.Data["utype"] = utype
	//log.Println(beego.AppConfig.String("loglevel"), "Debug", "utype", utype)

	username := sess.Get("username").(string)
	username1 := strings.ToUpper(username)
	c.Data["username"] = username1

	user_type := sess.Get("user_type").(string)
	user_type1 := strings.ToUpper(user_type)
	c.Data["user_type"] = user_type1

	mobile := sess.Get("mobile").(string)
	mobile1 := strings.ToUpper(mobile)
	c.Data["mobile"] = mobile1

	//language

	language := sess.Get("language").(string)
	c.Data["language"] = language

	role := sess.Get("role").(string)
	c.Data["role"] = role
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "role :- ", role)
	menu := sess.Get("menu").(string)
	c.Data["menu"] = menu
	submenu := sess.Get("submenu").(string)
	c.Data["submenu"] = submenu
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "usertype :- ", user_type)
	auth, err := utils.IsAuthorized(role, "systemmanagement-menu", "searchstatus-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	return
}
func (c *CreateStatus) Post() {
	var statusmsg string

	log.Println(beego.AppConfig.String("loglevel"), "Info", "add asset post page")
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)
	var err error
	var Autherr error
	sessErr := false
	defer func() {
		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
			c.TplName = "error/error.html"
		}
		if Autherr != nil {
			c.Data["DisplayMessage"] = Autherr.Error()
			c.TplName = "error/error.html"
			return
		}
		if err != nil {
			if sessErr == true {
				log.Println(beego.AppConfig.String("loglevel"), "Info", "Redirecting to login")
				c.Redirect(beego.AppConfig.String("LOGIN_PATH"), 302)
			} else {
				c.Data["DisplayMessage"] = err.Error()
			}
			c.TplName = "systemconfiguration/status/createStatus/createStatus.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Status Page Fail")
		} else {
			c.Data["DisplayMessage"] = statusmsg
			c.TplName = "systemconfiguration/status/createStatus/createStatus.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Status  Page Success")
		}
		return
	}()
	utils.SetHTTPHeader(c.Ctx)
	sess, err := session.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System is unable to process your request.Please contact customer care")
		sessErr = true
		return
	}

	if err = session.ValidateSession(sess); err != nil {
		sess.SessionRelease(c.Ctx.ResponseWriter)
		session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		sessErr = true
		return
	}
	defer func() {
		utils.EventLogs(c.Ctx, sess, c.Ctx.Input.Method(), c.Input(), c.Data, err)
		sess.SessionRelease(c.Ctx.ResponseWriter)
	}()
	uid := sess.Get("uid").(string)
	c.Data["uid"] = uid

	username := sess.Get("username").(string)
	username1 := strings.ToUpper(username)
	c.Data["username"] = username1

	user_type := sess.Get("user_type").(string)
	user_type1 := strings.ToUpper(user_type)
	c.Data["user_type"] = user_type1

	mobile := sess.Get("mobile").(string)
	mobile1 := strings.ToUpper(mobile)
	c.Data["mobile"] = mobile1

	//language

	language := sess.Get("language").(string)
	c.Data["language"] = language

	role := sess.Get("role").(string)
	c.Data["role"] = role
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "role :- ", role)
	menu := sess.Get("menu").(string)
	c.Data["menu"] = menu
	submenu := sess.Get("submenu").(string)
	c.Data["submenu"] = submenu
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "usertype :- ", user_type)
	auth, err := utils.IsAuthorized(role, "systemmanagement-menu", "searchstatus-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	id := uuid.New()
	fmt.Println(id.String())

	input_full_name := c.Input().Get("input_full_name")
	input_desc := c.Input().Get("input_desc")

	err = CheckUserAlreadyExists(input_full_name)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Status already Exists")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_STATUS_ALREADY_EXISTS"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_STATUS_ALREADY_EXISTS"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	result, err := db.Db.Exec(`INSERT INTO status (id,
	    name,
		"desc",
		created_by,
		created_at)
		VALUES ($1, $2, $3,$4,now())`,
		id,
		input_full_name,
		input_desc,
		uid)
	if err != nil {
		//err = errors.New("Status creation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_STATUS_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_STATUS_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	i, err := result.RowsAffected()
	if err != nil || i == 0 {
		//err = errors.New("Status creation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_STATUS_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_STATUS_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	if language == "english" {

		statusmsg = beego.AppConfig.String("EN_STATUS_CREATESUCCESFULLY")

	} else if language == "french" {

		statusmsg = beego.AppConfig.String("FN_STATUS_CREATESUCCESFULLY")

	}

	return
}

func CheckUserAlreadyExists(name string) (err error) {

	err = nil

	row, err := db.Db.Query(`select count(*) from status WHERE name = $1`, name)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to fetch status")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to fetch status")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

	countlen := data[0][0]

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "countlen", countlen)

	if countlen != "0" {
		err = errors.New("Status already exists")
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	return

}
