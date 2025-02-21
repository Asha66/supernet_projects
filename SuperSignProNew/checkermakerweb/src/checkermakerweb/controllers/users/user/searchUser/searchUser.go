/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package searchUser

import (
	"checkermakerweb/session"

	"runtime/debug"

	"checkermakerweb/utils"
	"errors"

	// "checkermakerweb/model/db"
	"strconv"

	"strings"

	"checkermakerweb/services"

	"github.com/astaxie/beego"

	// "proyava.com/database/sql"

	log "github.com/sirupsen/logrus"
)

type Row struct {
	Id         string
	FullName   string
	Mobile     string
	Email      string
	Address1   string
	Address2   string
	PartialId  string
	Status     string
	Timestamp  string
	Timestamp1 string
	UserType   string
}

type SearchUser struct {
	beego.Controller
}

func (c *SearchUser) Get() {
	Utype := c.Ctx.Input.Param(":Utype")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Utype", Utype)
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Search Customer Start")
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
			c.TplName = "users/user/searchUser/searchUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search User Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "users/user/searchUser/searchUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search User  Page Success")
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

	// utype := sess.Get("user_type").(string)
	// c.Data["utype"] = utype
	// //log.Println(beego.AppConfig.String("loglevel"), "Debug", "utype", utype)

	// language := sess.Get("language").(string)
	// c.Data["language"] = language
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "language", language)

	// fullname := sess.Get("fullname").(string)
	// fullname1 := strings.ToUpper(fullname)
	// c.Data["fullname1"] = fullname1

	// auth, err := utils.IsAuthorized(utype, "UserManagment")
	// if !auth {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
	// 	Autherr = errors.New("UnAuthorized")
	// 	return
	// }
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

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
	auth, err := utils.IsAuthorized(role, "userrmanagement-menu", "searchuser-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	data, err := services.SearchUsers()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Users fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_USER_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_USER_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	var result []Row
	var ts, tdate, ttime string

	for i := range data {
		var r Row
		r.Id = data[i][0]
		r.PartialId = r.Id[0:8]
		r.FullName = data[i][1]
		r.Mobile = data[i][2]
		r.Email = data[i][3]
		r.Address1 = data[i][4]
		s1 := data[i][5]
		b1, _ := strconv.ParseBool(s1)

		if b1 == true {

			r.Status = "ACTIVE"

		} else {

			r.Status = "INACTIVE"

		}
		r.Timestamp1 = data[i][6]
		ts = data[i][6]
		tdate = ts[0:10]
		ttime = ts[11:19]
		r.Timestamp = tdate + " " + ttime
		r.UserType = data[i][7]
		r.Address2 = data[i][4] + data[i][8]
		result = append(result, r)

	}
	c.Data["CustomerData"] = result

	return
}

func (c *SearchUser) Post() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "node post page")
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
			c.TplName = "users/user/searchUser/searchUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search User Page Fail")
		} else {
			c.Data["DisplayMessage"] = " "
			c.TplName = "users/user/searchUser/searchUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search User  Page Success")
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

	// utype := sess.Get("user_type").(string)
	// c.Data["utype"] = utype
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
	auth, err := utils.IsAuthorized(role, "userrmanagement-menu", "searchuser-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	input_name := c.Input().Get("input_name")
	input_email := c.Input().Get("input_email")
	input_status := c.Input().Get("input_status")

	dateRange := c.Input().Get("daterange")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Description - ", dateRange)

	c.Data["selectDate"] = dateRange

	from := ""
	to := ""

	if dateRange != "" {
		data := strings.Split(dateRange, " - ")

		if len(data) == 2 {
			from = data[0]
			to = data[1]
		}
		// log.Println(beego.AppConfig.String("loglevel"), "Debug", "fromDate, toDate ", from, to)
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "From Date - ", from)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "to Date - ", to)

	data, err := services.SearchUsersByFilter(input_name+"%", input_email+"%", from, to, input_status)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Users fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_USER_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_USER_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	var result []Row
	var ts, tdate, ttime string

	for i := range data {
		var r Row
		r.Id = data[i][0]
		r.PartialId = r.Id[0:8]
		r.FullName = data[i][1]
		r.Mobile = data[i][2]
		r.Email = data[i][3]
		r.Address1 = data[i][4]
		s1 := data[i][5]
		b1, _ := strconv.ParseBool(s1)

		if b1 == true {

			r.Status = "ACTIVE"

		} else {

			r.Status = "INACTIVE"

		}
		r.Timestamp1 = data[i][6]
		ts = data[i][6]
		tdate = ts[0:10]
		ttime = ts[11:19]
		r.Timestamp = tdate + " " + ttime
		r.UserType = data[i][7]
		r.Address2 = data[i][4] + " " + data[i][8]
		result = append(result, r)

	}
	c.Data["CustomerData"] = result

	return
}
