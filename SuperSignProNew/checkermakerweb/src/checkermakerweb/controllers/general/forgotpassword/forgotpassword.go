/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package forgotpassword

import (
	"checkermakerweb/model/db"
	"checkermakerweb/session"
	"checkermakerweb/utils"
	"crypto/rand"
	"errors"
	"net/smtp"
	"runtime/debug"
	"strconv"

	"checkermakerweb/utils/database/sql"

	log "github.com/sirupsen/logrus"

	"checkermakerweb/utils/encoding/base64"

	"github.com/astaxie/beego"
	"github.com/scorredoira/email"

	//"proyava.com/database/sql"

	//"proyava.com/util/log"
	"checkermakerweb/utils/util/password"
	"checkermakerweb/utils/util/pbkdf2"

	"io/ioutil"
	"net/mail"
	"strings"
)

type unencryptedAuth struct {
	smtp.Auth
}
type Forgotpassword struct {
	beego.Controller
}

func (c *Forgotpassword) Get() {

	var err error
	sessErr := false

	defer func() {
		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
			c.Data["DisplayMessage"] = "Something went wrong.Please Contact CustomerCare."
			c.TplName = "error/error.html"
		}
		if err != nil {
			if sessErr == true {
				log.Println(beego.AppConfig.String("loglevel"), "Info", "Redirecting to login")
				c.Redirect(beego.AppConfig.String("LOGIN_PATH"), 302)

			} else {
				c.Data["DisplayMessage"] = err.Error()
			}
			c.TplName = "general/forgotPassword/forgotPassword.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Forgot Password  Page Fail")
		} else {

			c.TplName = "general/forgotPassword/forgotPassword.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Forgot Password  Page Success")
		}
		return
	}()
	//	utils.SetHTTPHeader(c.Ctx)

	session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", c.Ctx.Input.IP())
	c.TplName = "general/login/login.html"

	sess, err := session.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System is unable to process your request.Please contact customer care")
		return
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	sessionId := sess.SessionID()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Session ID - ", sessionId)

	// validationpath1 := beego.AppConfig.String("VALIDATION_LANG_PATH")
	// sess.Set("VALIDATION_LANG_PATH", validationpath1)

	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "VALIDATION_LANG_PATH - ", validationpath1)

	// vpath := sess.Get("VALIDATION_LANG_PATH").(string)
	// c.Data["vpath"] = vpath
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "vpath", vpath)

	defer func() {
		utils.EventLogs(c.Ctx, sess, c.Ctx.Input.Method(), c.Input(), c.Data, err)
		sess.SessionRelease(c.Ctx.ResponseWriter)
	}()

	return
}

func (c *Forgotpassword) Post() {

	log.Println(beego.AppConfig.String("loglevel"), "Info", "Forgot Password post page")
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
			c.TplName = "general/forgotPassword/forgotPassword.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Forgot Password Page Fail")
		} else {
			c.Data["DisplayMessage"] = "Password has been reset successfully."
			c.TplName = "general/forgotPassword/forgotPassword.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Forgot Password  Page Success")
		}
		return
	}()

	utils.SetHTTPHeader(c.Ctx)
	sess, err := session.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System is unable to process your request.Please contact customer care")
		return
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	sessionId := sess.SessionID()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Session ID - ", sessionId)

	defer func() {
		utils.EventLogs(c.Ctx, sess, c.Ctx.Input.Method(), c.Input(), c.Data, err)
		sess.SessionRelease(c.Ctx.ResponseWriter)
	}()

	//utils.SetHTTPHeader(c.Ctx)

	uname := c.Input().Get("username")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "UserEmail - ", uname)

	// if uname == "" {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", "Blank User Email")
	// 	err = errors.New("User-email can't be blank.")
	// 	return
	// }

	err = SearchUser(uname)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//	err = errors.New("User Not Found")
		return
	}

	newPass, _ := password.AlphaNumericSpecial(6)

	err = UpdatePassword(uname, newPass)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System Admin User Update Password Failed")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "New Password :- ", newPass)

	data1, err := TemplateFormate()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Notification Template Not Present for Forgotpassword at AdminWeb Channel")
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

	err = SendEmail(uname, "", newPass, data1[0][1], data1[0][2], data1[0][4], data1[0][5], data1[0][6], data1[0][7], data1[0][8])
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("USER_SENDMAIL__NOT_FOUND")
		return
	}

	return
}

func SearchUser(uname string) (err error) {

	row, err := db.Db.Query("SELECT id,email,status FROM sysuser where email=$1 limit 1", uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System Admin User Not Found")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System Admin User Detail Scan Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("System Admin User Not Found")
		return
	}

	status, _ := strconv.ParseBool(data[0][2])

	if status == false {
		err = errors.New("System Admin User is currently Suspended")
		return
	}

	return
}

func UpdatePassword(uname, password string) (err error) {

	b := make([]byte, 32)
	_, err = rand.Read(b)
	var pbkdf pbkdf2.Pbkdf2
	pbkdf.Itr = 32
	pbkdf.KeyLen = 32
	pbkdf.Plain = []byte(password)
	pbkdf.Salt = b
	err = pbkdf.Encrypt()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New(" System Admin User Password Update Fail")
		return
	}
	var tmp []byte
	tmp = append(tmp, pbkdf.Salt...)
	tmp = append(tmp, pbkdf.Cipher...)

	out, err := base64.Encode(tmp)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New(" System Admin User Password Update Fail")
		return
	}

	result, err := db.Db.Exec("update sysuser set password=$1,password_updated_date=now() where email=$2 ", string(out), uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New(" System Admin User Password Update Fail")
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New(" System Admin User Password Update Fail")
		return
	}

	if n != 1 {
		err = errors.New("System Admin User Password Update Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "System Admin User new Password : ", password)
	return
}

func SendEmail(emilid string, name string, password, title, desc, url1, template1, template2, template3, describeurl string) (err error) {
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "called - ")

	uname := beego.AppConfig.String("EMAIL_NOTIFY_USERNAME")
	pass := beego.AppConfig.String("EMAIL_NOTIFY_PASSWORD")
	url := beego.AppConfig.String("EMAIL_NOTIFY_URL")
	to := beego.AppConfig.String("EMAIL_NOTIFY_TIMEOUT")
	recipients := strings.Split(emilid, "||")
	loginurl := beego.AppConfig.String("EMAIL_APPLICATION_LOGIN_URL")

	tmpFile := beego.AppConfig.String("FORGOT_EMAIL_TEMPLATE")

	buff, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", "read file -", err)
		return
	}

	msg := string(buff)
	msg = strings.Replace(string(msg), "{{.Name}}", name, -1)
	msg = strings.Replace(string(msg), "{{.Email}}", emilid, -1)
	msg = strings.Replace(string(msg), "{{.Password}}", password, -1)
	msg = strings.Replace(string(msg), "{{.LoginURL}}", loginurl, -1)

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

	row, err := db.Db.Query(`select id,title,"desc",channel,url,template1,template2,template3,describe_url,created_at from templates where channel='AdminWeb' and template_type='Forgotpassword' ORDER BY created_at Desc limit 1`)
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
