/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package emailVerfication

import (
	"checkermakerweb/model/db"
	"checkermakerweb/services"
	"checkermakerweb/utils"

	"errors"

	"io/ioutil"
	"net/mail"
	"net/smtp"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/scorredoira/email"

	"fmt"
	//"time"

	"checkermakerweb/session"

	"checkermakerweb/utils/database/sql"

	log "github.com/sirupsen/logrus"

	//"proyava.com/database/sql"
	//"checkermakerweb/utils/encoding/base64"
	//"proyava.com/util/log"
	//"checkermakerweb/utils/util/pbkdf2"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/validation"
)

type EmailVerfication struct {
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
}
type Field1 struct {
	Id   string
	Name string
}

func (c *EmailVerfication) Get() {
	TransactionId := c.Ctx.Input.Param(":TransactionID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "TransactionId - ", TransactionId)
	defer func() {
		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
			c.Data["DisplayMessage"] = "Something went wrong.Please Contact CustomerCare."
			c.TplName = "error/error.html"
		}
		return
	}()
	//	utils.SetHTTPHeader(c.Ctx)

	session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", c.Ctx.Input.IP())
	c.TplName = "email/emailVerfication/emailVerfication.html"

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

	c.Data["TxnId"] = TransactionId

	//TxnId := data[0][5]
	//log.Println(beego.AppConfig.String("loglevel"), "TxnId", TxnId)
	return
}

func (c *EmailVerfication) Post() {

	TransactionId := c.Ctx.Input.Param(":TransactionID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "TransactionId - ", TransactionId)
	c.Data["TxnId"] = TransactionId

	log.Println(beego.AppConfig.String("loglevel"), "Info", "emailVerfication Start")
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)

	msg := ""

	var err error

	defer func() {

		if err != nil {
			if l_exception := recover(); l_exception != nil {
				stack := debug.Stack()
				log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
				session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
				c.TplName = "error/error.html"
			} else {
				c.Data["DisplayMessage"] = err.Error()
			}
			c.TplName = "email/emailVerfication/emailVerfication.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "emailVerfication Fail")

		} else {

			c.Data["DisplayMessage"] = msg
			c.TplName = "email/emailVerfication/emailVerfication.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "emailVerfication Success")
		}
		return
	}()

	//	utils.SetHTTPHeader(c.Ctx)
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

	//username := c.Input().Get("username")
	password := c.Input().Get("password")
	confirmpassword := c.Input().Get("confirmpassword")

	row11, err := db.Db.Query(`select count(*) from sysuser where txn_id=$1`, TransactionId)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get SystemUser data")

		return
	}
	defer sql.Close(row11)
	_, data11, err := sql.Scan(row11)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get SystemUser data")

		return
	}

	row12, err := db.Db.Query(`select count(*) from users where txn_id=$1`, TransactionId)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get SystemUser data")

		return
	}
	defer sql.Close(row12)
	_, data12, err := sql.Scan(row12)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get SystemUser data")

		return
	}

	count1, err := strconv.Atoi(data11[0][0])
	count2, err := strconv.Atoi(data12[0][0])
	log.Println(beego.AppConfig.String("loglevel"), "sys user count", count1)
	log.Println(beego.AppConfig.String("loglevel"), "user count", count2)

	if count1 > 0 {

		fmt.Println("Email Verification For System User")

		row, err := db.Db.Query(`select fullname,mobile,email,language,password_set,txn_id from sysuser where txn_id=$1`, TransactionId)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			//err = errors.New("Unable to get SystemUser data")

			msg = "Unable to get SystemUser data"

			return
		}
		defer sql.Close(row)
		_, data, err := sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			// err = errors.New("Unable to get SystemUser data")
			msg = "Unable to get SystemUser data"

			return
		}

		log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

		email := data[0][2]
		language := data[0][3]
		passwordset1 := data[0][4]
		txnid1 := data[0][5]

		passwordset, _ := strconv.ParseBool(passwordset1)

		log.Println(beego.AppConfig.String("loglevel"), "Debug", "email", email)
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "password", password)
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "confirmpassword", confirmpassword)
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "password set value", passwordset)

		if password != confirmpassword {
			log.Println(beego.AppConfig.String("loglevel"), "Error", "New Password Mismatch")
			err = errors.New("New password and Confirm password can't be different")
			if language == "english" {
				msg = beego.AppConfig.String("EN_NEW_PASSWORD_AND_CONFIRM_PASSWORD_CAN'T_BE_DIFFERENT")
				log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
				return
			} else if language == "french" {
				msg = beego.AppConfig.String("FN_NEW_PASSWORD_AND_CONFIRM_PASSWORD_CAN'T_BE_DIFFERENT")
				log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
				return
			}
			return
		}
		if passwordset == false {

			err = UpdatePassword(password, email, txnid1)

			if err != nil {
				log.Println(beego.AppConfig.String("loglevel"), "Error", err)
				//err = errors.New("Password Updation Failed")
				msg = "Password Updation Failed"

				return
			}

			data1, err := TemplateFormate("AdminWeb")
			if err != nil {
				log.Println(beego.AppConfig.String("loglevel"), "Error", err)
				//err = errors.New("templates fetch Failed")
				if language == "english" {
					msg = beego.AppConfig.String("EN_TEMPLATE_FETCH_FAILED")
					log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
					return
				} else if language == "french" {
					msg = beego.AppConfig.String("FN_TEMPLATE_FETCH_FAILED")
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

			go SendEmail(email, data[0][0], password, data1[0][1], data1[0][2], data1[0][4], data1[0][5], data1[0][6], data1[0][7], data1[0][8], beego.AppConfig.String("USER_REGISTRATION_PATH"), beego.AppConfig.String("VERIFY_EMAIL_TEMPLATE"), "SYSUSER")

			if language == "english" {

				msg = "Your email is verified,please check your email for login url"

			} else if language == "french" {

				msg = "Votre e-mail est vérifié, veuillez vérifier votre e-mail pour l'URL de connexion"
			}

		} else {

			if language == "english" {

				msg = "Your email is already verified"

			} else if language == "french" {

				msg = "Votre email est déjà vérifié"
			}

		}

	} else if count2 > 0 {

		fmt.Println("Email Verification For Distributor")

		row, err := db.Db.Query(`select full_name,mobile,email,language,password_set,txn_id from users where txn_id=$1`, TransactionId)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			//err = errors.New("Unable to get User data")
			msg = "Unable to get User data"

			return
		}
		defer sql.Close(row)
		_, data, err := sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			//err = errors.New("Unable to get SystemUser data")
			msg = "Unable to get SystemUser data"

			return
		}

		log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

		email := data[0][2]
		language := data[0][3]
		passwordset1 := data[0][4]
		txnid1 := data[0][5]

		passwordset, _ := strconv.ParseBool(passwordset1)

		log.Println(beego.AppConfig.String("loglevel"), "Debug", "email", email)
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "password", password)
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "confirmpassword", confirmpassword)
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "password set value", passwordset)

		if password != confirmpassword {
			log.Println(beego.AppConfig.String("loglevel"), "Error", "New Password Mismatch")
			//err = errors.New("New password and Confirm password can't be different")
			if language == "english" {
				msg = beego.AppConfig.String("EN_NEW_PASSWORD_AND_CONFIRM_PASSWORD_CAN'T_BE_DIFFERENT")
				log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
				return
			} else if language == "french" {
				msg = beego.AppConfig.String("FN_NEW_PASSWORD_AND_CONFIRM_PASSWORD_CAN'T_BE_DIFFERENT")
				log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
				return
			}
			return
		}
		if passwordset == false {

			err = UpdatePasswordUser(password, email, txnid1)

			if err != nil {
				log.Println(beego.AppConfig.String("loglevel"), "Error", err)
				//err = errors.New("Password Updation Failed")
				msg = "Password Updation Failed"

				return
			}

			data1, err := TemplateFormate("DistributorWeb")
			if err != nil {
				log.Println(beego.AppConfig.String("loglevel"), "Error", err)
				//err = errors.New("templates fetch Failed")
				if language == "english" {
					msg = beego.AppConfig.String("EN_TEMPLATE_FETCH_FAILED")
					log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", msg)
					return
				} else if language == "french" {
					msg = beego.AppConfig.String("FN_TEMPLATE_FETCH_FAILED")
					log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", msg)
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

			go SendEmail(email, data[0][0], password, data1[0][1], data1[0][2], data1[0][4], data1[0][5], data1[0][6], data1[0][7], data1[0][8], beego.AppConfig.String("USER_REGISTRATION_PATH"), beego.AppConfig.String("VERIFY_EMAIL_TEMPLATE"), "USER")

			if language == "english" {

				msg = "Your email is verified,please check your email for login url"

			} else if language == "french" {

				msg = "Votre e-mail est vérifié, veuillez vérifier votre e-mail pour l'URL de connexion"
			}

		} else {

			if language == "english" {

				msg = "Your email is already verified"

			} else if language == "french" {

				msg = "Votre email est déjà vérifié"
			}

		}
	}

	return
}

func UpdatePassword(password, username, txnid string) (err error) {

	pass, err := services.EncryptPassword(password)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	result, err := db.Db.Exec("update sysuser set password=$1 ,password_set=true,password_updated_date=now() where email=$2 and txn_id=$3 ", pass, username, txnid)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System User Password Update Fail")
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System User Password Update Fail")
		return
	}

	if n != 1 {
		err = errors.New("System User Password Update Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "pass")
	return
}
func UpdatePasswordUser(password, username, txnid string) (err error) {

	pass, err := services.EncryptPassword(password)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	result, err := db.Db.Exec("update users set password=$1 ,password_set=true,password_updated_date=now() where email=$2 and txn_id=$3 ", pass, username, txnid)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System User Password Update Fail")
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System User Password Update Fail")
		return
	}

	if n != 1 {
		err = errors.New("System User Password Update Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "pass")
	return
}

func SendEmail(emilid, name, password, title, desc, url1, template1, template2, template3, describeurl, emailUsertemplate, emailtemplate, usertype string) {
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "called - ")

	uname := beego.AppConfig.String("EMAIL_NOTIFY_USERNAME")
	pass := beego.AppConfig.String("EMAIL_NOTIFY_PASSWORD")
	url := beego.AppConfig.String("EMAIL_NOTIFY_URL")
	to := beego.AppConfig.String("EMAIL_NOTIFY_TIMEOUT")
	loginurl := beego.AppConfig.String("EMAIL_APPLICATION_LOGIN_URL")

	userregistrationurl := beego.AppConfig.String("USERS_REGISTRATION_PATH")
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
func TemplateFormate(channel string) (data [][]string, err error) {

	row, err := db.Db.Query(`select id,title,"desc",channel,url,template1,template2,template3,describe_url,created_at from templates where channel=$1 ORDER BY created_at Desc limit 1`, channel)
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
