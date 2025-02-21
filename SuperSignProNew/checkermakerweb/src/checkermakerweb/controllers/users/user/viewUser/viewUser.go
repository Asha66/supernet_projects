/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package viewUser

import (
	"checkermakerweb/session"

	"checkermakerweb/utils"
	"runtime/debug"

	"checkermakerweb/services"
	"errors"
	"strconv"
	"strings"

	"checkermakerweb/model/db"

	"github.com/astaxie/beego"

	"checkermakerweb/utils/database/sql"

	log "github.com/sirupsen/logrus"
)

type Row struct {
	Id       string
	FullName string
	Mobile   string
	Email    string
	Address  string
	Status   string
	UType    string
}

type Display4 struct {
	Fields4 []Field4
}
type Field4 struct {
	Id   string
	Name string
}

type ViewUser struct {
	beego.Controller
}

func (c *ViewUser) Get() {
	AdminId := c.Ctx.Input.Param(":AdminID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "AdminId", AdminId)
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Search SystemUser Start")
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
			c.TplName = "users/user/viewUser/viewUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search User Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "users/user/viewUser/viewUser.html"
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

	data4, err4 := services.GetUserType1()
	if err4 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err4 = errors.New("User Type fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("ENGLISH_USER_TYPE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FRENCH_USER_TYPE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	if len(data4) <= 0 {
		//err4 = errors.New("user Type  Not Found")
		if language == "english" {
			err4 = errors.New(beego.AppConfig.String("ENGLISH_USER_TYPE_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err4 = errors.New(beego.AppConfig.String("FRENCH_USER_TYPE_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	var Dis4 Display4
	for i := range data4 {
		var d Field4
		d.Id = data4[i][0]
		d.Name = data4[i][1]
		Dis4.Fields4 = append(Dis4.Fields4, d)
	}
	c.Data["Dis4"] = Dis4
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis4)

	//	input_user_type := c.Input().Get("input_user_type")

	row, err := db.Db.Query(`select users.id,
	users.full_name,
	users.mobile,
	users.email,
	users.address1,
	users.status,
	users.created_at,
	usertypes.name,
	users.address2,
	users.language from users
	left join usertypes on users.user_type=usertypes.id where users.id=$1`, AdminId)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get User data")
		if language == "english" {
			err4 = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err4 = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get User data")
		if language == "english" {
			err4 = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err4 = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	if len(data) <= 0 {
		//err = errors.New("User data not found")
		if language == "english" {
			err4 = errors.New(beego.AppConfig.String("EN_USER_DATA_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err4 = errors.New(beego.AppConfig.String("FN_USER_DATA_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

	for i := range data {

		c.Data["Id"] = data[i][0]
		c.Data["FullName"] = data[i][1]
		c.Data["Mobile"] = data[i][2]
		c.Data["Email"] = data[i][3]
		c.Data["Address1"] = data[i][4]
		//c.Data["Status"] = data[i][5]
		s1 := data[i][5]
		b1, _ := strconv.ParseBool(s1)

		if b1 == true {

			c.Data["Status"] = "ACTIVE"

		} else {

			c.Data["Status"] = "INACTIVE"

		}

		c.Data["UType"] = data[i][7]
		c.Data["Address2"] = data[i][4] + " " + data[i][8]
		c.Data["Language"] = data[i][9]

	}

	return

}
