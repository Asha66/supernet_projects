/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package createsysUser

import (
	"checkermakerweb/session"

	"io/ioutil"
	"net/mail"

	"net/smtp"
	"runtime/debug"
	"strings"

	//	"time"

	"checkermakerweb/model/db"
	"checkermakerweb/utils"
	"errors"

	"github.com/scorredoira/email"

	"github.com/astaxie/beego"

	"checkermakerweb/services"

	// "proyava.com/database/sql"
	// "proyava.com/util/log"

	"checkermakerweb/utils/database/sql"

	"fmt"

	log "github.com/sirupsen/logrus"

	"checkermakerweb/utils/util/password"
	"checkermakerweb/utils/util/txnno"

	"github.com/google/uuid"
)

type CreatesysUser struct {
	beego.Controller
}

type unencryptedAuth struct {
	smtp.Auth
}

type Field struct {
	Id    string
	Name  string
	Email string
}

type Display struct {
	Fields1 []Field1
	Fields2 []Field2
}
type Field1 struct {
	Id   string
	Name string
}
type Field2 struct {
	Id   string
	Name string
}

func (c *CreatesysUser) Get() {
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
			c.TplName = "users/sysusers/createsysUsers/createsysUsers.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search System User Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "users/sysusers/createsysUsers/createsysUsers.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search System User  Page Success")
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

	username := sess.Get("username").(string)
	username1 := strings.ToUpper(username)
	c.Data["username"] = username1

	user_type := sess.Get("user_type").(string)
	user_type1 := strings.ToUpper(user_type)
	c.Data["user_type"] = user_type1

	mobile := sess.Get("mobile").(string)
	mobile1 := strings.ToUpper(mobile)
	c.Data["mobile"] = mobile1

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

	auth, err := utils.IsAuthorized(role, "sysusermanagement-menu", "searchsysuser-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	data, err := services.GetStatus()
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
	if len(data) <= 0 {
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
	var Dis Display
	for i := range data {
		var d Field1
		d.Id = data[i][0]
		d.Name = data[i][1]
		Dis.Fields1 = append(Dis.Fields1, d)
	}
	c.Data["Dis"] = Dis
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis)

	data1, err := services.GetRole()
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
	var Dis1 Display
	for i := range data1 {
		var d Field2
		d.Id = data1[i][0]
		d.Name = data1[i][1]
		Dis1.Fields2 = append(Dis1.Fields2, d)
	}
	c.Data["Dis1"] = Dis1
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis1)

	return
}
func (c *CreatesysUser) Post() {
	var systemusermsg string

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
			c.TplName = "users/sysusers/createsysUsers/createsysUsers.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search System User Page Fail")
		} else {
			c.Data["DisplayMessage"] = systemusermsg
			c.TplName = "users/sysusers/createsysUsers/createsysUsers.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search System User  Page Success")
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

	auth, err := utils.IsAuthorized(role, "sysusermanagement-menu", "searchsysuser-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	data, err := services.GetStatus()
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
	if len(data) <= 0 {
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
	var Dis Display
	for i := range data {
		var d Field1
		d.Id = data[i][0]
		d.Name = data[i][1]
		Dis.Fields1 = append(Dis.Fields1, d)
	}
	c.Data["Dis"] = Dis
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis)

	data3, err := services.GetRole()
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
	if len(data3) <= 0 {
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
	var Dis1 Display
	for i := range data3 {
		var d Field2
		d.Id = data3[i][0]
		d.Name = data3[i][1]
		Dis1.Fields2 = append(Dis1.Fields2, d)
	}
	c.Data["Dis1"] = Dis1
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis1)

	uid := sess.Get("uid").(string)
	c.Data["uid"] = uid

	input_full_name := c.Input().Get("input_full_name")
	input_mobile := c.Input().Get("input_mobile")
	input_email := c.Input().Get("input_email")
	input_address := c.Input().Get("input_address")
	input_language := c.Input().Get("input_language")
	input_status := c.Input().Get("input_status")
	input_role := c.Input().Get("input_role")

	var channelstatus bool

	if input_status == "ACTIVE" {

		channelstatus = true
	} else {
		channelstatus = false
	}

	data1, err := TemplateFormate()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("templates fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_TEMPLATE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_TEMPLATE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	for i := range data1 {

		c.Data["Id"] = data1[i][0]
		c.Data["Title"] = data1[i][1]
		c.Data["Desc"] = data1[i][2]
		c.Data["Channel"] = data1[i][3]
		c.Data["Url"] = data1[i][4]
		c.Data["Template1"] = data1[i][5]
		c.Data["Template2"] = data1[i][6]
		c.Data["Template3"] = data1[i][7]
		c.Data["DescribeUrl"] = data1[i][8]

	}

	err = CheckUserAlreadyExists(input_email, input_mobile)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("SystemUser already Exists")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SYSTEMUSER_ALREADY_EXISTS"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SYSTEMUSER_ALREADY_EXISTS"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	id := uuid.New()
	fmt.Println(id.String())

	loginPass, _ := password.AlphaNumericSpecial(6)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Login Pass - ", loginPass)

	pass, err := services.EncryptPassword(loginPass)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	txn_id := txnno.Generate13Digit()

	result, err := db.Db.Exec(`INSERT INTO sysuser (id,
	    fullname,
		password,
		mobile,
		email,
		address,
		status,
		language,
		password_set,
		created_by,
		txn_id,
		role_id,
		created_at,
		password_updated_date)
		VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10,$11,$12,now(),now())`,
		id,
		input_full_name,
		pass,
		input_mobile,
		input_email,
		input_address,
		channelstatus,
		input_language,
		"false",
		uid,
		txn_id,
		input_role)
	if err != nil {
		//err = errors.New("SystemUser creation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SYSTEMUSER_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SYSTEMUSER_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	i, err := result.RowsAffected()
	if err != nil || i == 0 {
		//err = errors.New("SystemUser creation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SYSTEMUSER_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SYSTEMUSER_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	data2, err := SearchSysuser(input_email)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("templates fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SYSTEMUSER_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SYSTEMUSER_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	for i := range data2 {

		c.Data["TxnId"] = data2[i][0]

	}

	log.Println(beego.AppConfig.String("loglevel"), "Transid", data2[0][0])

	go SendEmail(input_email, input_full_name, loginPass, data1[0][1], data1[0][2], data1[0][4], data1[0][5], data1[0][6], data1[0][7], data1[0][8], beego.AppConfig.String("USER_REGISTRATION_PATH"), beego.AppConfig.String("EMAIL_TEMPLATE"), data2[0][0])

	if language == "english" {

		systemusermsg = beego.AppConfig.String("EN_SYSTEMUSER_CREATESUCCESFULLY")

	} else if language == "french" {

		systemusermsg = beego.AppConfig.String("FN_SYSTEMUSER_CREATESUCCESFULLY")

	}

	return
}

func CheckUserAlreadyExists(email, mobile string) (err error) {

	err = nil

	row, err := db.Db.Query(`select count(*) from sysuser WHERE email = $1 and mobile = $2`, email, mobile)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to fetch SystemUser")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to fetch SystemUser")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

	countlen := data[0][0]

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "countlen", countlen)

	if countlen != "0" {
		err = errors.New("SystemUser already exists")
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	return

}

func SendEmail(emilid, name, password, title, desc, url1, template1, template2, template3, describeurl, emailUsertemplate, emailtemplate, txnid string) {
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "called - ")

	uname := beego.AppConfig.String("EMAIL_NOTIFY_USERNAME")
	pass := beego.AppConfig.String("EMAIL_NOTIFY_PASSWORD")
	url := beego.AppConfig.String("EMAIL_NOTIFY_URL")
	to := beego.AppConfig.String("EMAIL_NOTIFY_TIMEOUT")
	loginurl := beego.AppConfig.String("EMAIL_APPLICATION_LOGIN_URL")

	userregistrationurl := beego.AppConfig.String("USERS_REGISTRATION_URL") + txnid
	recipients := strings.Split(emilid, "||")

	tmpFile := emailtemplate

	buff, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", "read file -", err)
		return
	}

	// tmpFile1 := emailUsertemplate

	// buff1, err := ioutil.ReadFile(tmpFile1)
	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", "read file -", err)
	// 	return
	// }

	msg := string(buff)
	msg = strings.Replace(string(msg), "{{.Name}}", name, -1)
	msg = strings.Replace(string(msg), "{{.Email}}", emilid, -1)
	msg = strings.Replace(string(msg), "{{.Password}}", password, -1)
	msg = strings.Replace(string(msg), "{{.LoginURL}}", loginurl, -1)

	// msg1 := string(buff1)
	msg = strings.Replace(string(msg), "{{.UserRegistrationUrl}}", userregistrationurl, -1)

	msg = strings.Replace(string(msg), "{{.Title}}", title, -1)
	msg = strings.Replace(string(msg), "{{.Desc}}", desc, -1)
	msg = strings.Replace(string(msg), "{{.Url}}", url1, -1)
	msg = strings.Replace(string(msg), "{{.Template1}}", template1, -1)
	msg = strings.Replace(string(msg), "{{.Template2}}", template2, -1)
	msg = strings.Replace(string(msg), "{{.Template3}}", template3, -1)
	msg = strings.Replace(string(msg), "{{.DescribeUrl}}", describeurl, -1)

	m := email.NewHTMLMessage("Email", msg)
	m.From = mail.Address{Name: "SUPERPAY", Address: uname}
	m.To = recipients

	// send it
	//auth := smtp.PlainAuth("", uname, pass, url)

	config := beego.AppConfig.String("EMAIL_AUTH_CONFIG_MODE")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "config", config)

	if config == "1" {
		auth := smtp.PlainAuth("", uname, pass, url)
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "auth")
		if err = email.Send(url+":"+to, auth, m); err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			return
		}

	} else if config == "2" {

		auth := unencryptedAuth{
			smtp.PlainAuth(
				"",
				uname,
				pass,
				url,
			),
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "no tls auth")
		if err = email.Send(url+":"+to, auth, m); err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			return
		}
	} else {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "no auth")
		if err = email.Send(url+":"+to, nil, m); err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			return
		}
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Email sent successfully")

	return
}

func TemplateFormate() (data [][]string, err error) {

	row, err := db.Db.Query(`select id,title,"desc",channel,url,template1,template2,template3,describe_url,created_at from templates where channel='AdminWeb' and template_type='Registration' ORDER BY created_at Desc limit 1`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get the error message info")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Message Scan Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("Unable to get the error message info")
		return
	}
	return

}

func SearchSysuser(email string) (data [][]string, err error) {

	row, err := db.Db.Query(`select txn_id from sysuser where email=$1 limit 1`, email)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get the error message info")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Message Scan Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("Unable to get the error message info")
		return
	}
	return

}
