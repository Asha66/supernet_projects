/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package updateUser

import (
	"checkermakerweb/model/db"
	"checkermakerweb/services"
	"checkermakerweb/session"
	"checkermakerweb/utils"
	"errors"
	"runtime/debug"

	"fmt"
	"strings"

	"checkermakerweb/utils/database/sql"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

type Row struct {
	Id           string
	FirstName    string
	MiddleName   string
	LastName     string
	Mobile       string
	Email        string
	Role         string
	Status       string
	Address1     string
	Address2     string
	Town         string
	City         string
	Pincode      string
	Language     string
	LocationType string
	LocationInfo string
	UserType     string
}
type Display struct {
	Fields []Field
}
type Field struct {
	Id    string
	Name  string
	Email string
}

type Display1 struct {
	Fields1 []Field1
}

type Field1 struct {
	Id   string
	Name string
}

type Display4 struct {
	Fields4 []Field4
}
type Field4 struct {
	Id   string
	Name string
}

type UpdateUser struct {
	beego.Controller
}

func (c *UpdateUser) Get() {
	AdminId := c.Ctx.Input.Param(":AdminID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "AdminId", AdminId)
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Update User Start")
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
			c.TplName = "users/user/updateUser/updateUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update User Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "users/user/updateUser/updateUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update User Page Success")
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

	data1, err := services.GetStatus()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Status fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("ENGLISH_STATUS_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FRENCH_STATUS_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	if len(data1) <= 0 {
		//err = errors.New("Status  Not Found")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("ENGLISH_STATUS_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FRENCH_STATUS_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	var Dis Display1
	for i := range data1 {
		var d Field1
		d.Id = data1[i][0]
		d.Name = data1[i][1]
		Dis.Fields1 = append(Dis.Fields1, d)
	}
	c.Data["Dis"] = Dis
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis)

	data4, err4 := services.GetUserType3()
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

	row, err := db.Db.Query(`select id,full_name,mobile,email,address1,status,address2,user_type,language from users where id=$1`, AdminId)
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
		c.Data["Status"] = data[i][5]
		c.Data["Address2"] = data[i][6]
		c.Data["UserType"] = data[i][7]
		c.Data["Language"] = data[i][8]

	}

	return
}
func (c *UpdateUser) Post() {
	var usermsg string

	AdminId := c.Ctx.Input.Param(":AdminID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "AdminId - ", AdminId)

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
			c.TplName = "users/user/updateUser/updateUser.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search User Page Fail")
		} else {
			c.Data["DisplayMessage"] = usermsg
			c.TplName = "users/user/updateUser/updateUser.html"
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

	uid := sess.Get("uid").(string)
	c.Data["uid"] = uid

	input_full_name := c.Input().Get("input_full_name")

	input_email := c.Input().Get("input_email")
	input_mobile := c.Input().Get("input_mobile")
	input_address1 := c.Input().Get("input_address1")
	input_address2 := c.Input().Get("input_address2")
	input_language := c.Input().Get("input_language")
	input_status := c.Input().Get("input_status")
	input_user_type := c.Input().Get("input_user_type")
	fmt.Println(input_status)

	var channelstatus bool

	if input_status == "ACTIVE" {

		channelstatus = true

		// err = ResetLoginCount(AdminId)
		// if err != nil {
		// 	err = errors.New("Admin User Login Count Reset Failed")
		// 	return
		// }
	} else {
		channelstatus = false
	}

	if input_full_name == "" || input_email == "" || input_mobile == "" || input_address1 == "" || input_address2 == "" ||
		input_status == "" || input_user_type == "" || input_language == "" {
		//err = errors.New("Any of the fields cannot be empty")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_ANY_FIELD_CANNOT_BE_EMPTY"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_ANY_FIELD_CANNOT_BE_EMPTY"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	fmt.Println(input_address1)

	res, err := db.Db.Exec(`UPDATE users SET full_name=$1,email=$2,mobile=$3,address1=$4,status=$5,address2=$6,updated_at=now(),user_type=$7,language=$8 WHERE id = $9`, input_full_name, input_email, input_mobile, input_address1, channelstatus, input_address2, input_user_type, input_language, AdminId)
	if err != nil {
		//err = errors.New("User updation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_USER_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_USER_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	i, err := res.RowsAffected()
	if err != nil || i == 0 {
		//err = errors.New("User updation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_USER_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_USER_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	if language == "english" {

		usermsg = beego.AppConfig.String("EN_USER_UPDATESUCCESSFULLY")

	} else if language == "french" {

		usermsg = beego.AppConfig.String("FN_USER_UPDATESUCCESSFULLY")

	}

	return
}

func ResetLoginCount(uname string) (err error) {
	count := 0
	result, err := db.Db.Exec("UPDATE users set pass_count=$1 where id=$2 ", count, uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Count Update Fail")
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Count Update Fail")
		return
	}

	if n != 1 {
		err = errors.New("User Count Update Fail")
		return
	}

	return
}
