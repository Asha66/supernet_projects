/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package changePassword

import (
	//	get "checkermakerweb/model/general"
	"checkermakerweb/session"
	"checkermakerweb/utils"
	"errors"
	"runtime/debug"

	"checkermakerweb/services"

	// "proyava.com/util/log"

	"checkermakerweb/model/db"
	"strings"

	"github.com/astaxie/beego"

	"checkermakerweb/utils/database/sql"

	log "github.com/sirupsen/logrus"

	"checkermakerweb/utils/encoding/base64"
	"checkermakerweb/utils/util/pbkdf2"
	// "proyava.com/database/sql"
	// "proyava.com/util/log"
)

type ChangePassword struct {
	beego.Controller
}

func (c *ChangePassword) Get() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Change Password Page Start")
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)
	var err error
	defer func() {
		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)

			c.TplName = "error/error.html"
		}

		if err != nil && err.Error() == "Session Time Out.Please Logout and Login Again." {
			c.Abort("500")
		}

		if err != nil {
			c.Data["DisplayMessage"] = err.Error()
			c.TplName = "error/error.html"

			log.Println(beego.AppConfig.String("loglevel"), "Info", "Change Password Page Fail")
		} else {
			c.TplName = "general/changePassword/changePassword.html"

			log.Println(beego.AppConfig.String("loglevel"), "Info", "Change Password Page Success")
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

	if err = session.ValidateSession(sess); err != nil {
		sess.SessionRelease(c.Ctx.ResponseWriter)
		session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	uname := sess.Get("uname").(string)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Username - ", uname)

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

	defer func() {
		utils.EventLogs(c.Ctx, sess, c.Ctx.Input.Method(), c.Input(), c.Data, err)
		sess.SessionRelease(c.Ctx.ResponseWriter)
	}()
	return
}

func (c *ChangePassword) Post() {
	var changepasswordmsg string
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Change Password Page Start")
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)
	var err error
	defer func() {

		if l_exception := recover(); l_exception != nil {
			stack := debug.Stack()
			log.Println(beego.AppConfig.String("loglevel"), "Exception", string(stack))
			session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)

			c.TplName = "error/error.html"
		}
		if err != nil && err.Error() == "Session Time Out.Please Logout and Login Again." {
			c.Abort("500")
		}

		if err != nil {
			c.Data["DisplayMessage"] = err.Error()
			c.Data["title"] = "Error !"
			c.Data["type"] = "error"
			c.TplName = "general/changePassword/changePassword.html"

			log.Println(beego.AppConfig.String("loglevel"), "Info", "Change Password Page Fail")
		} else {
			c.Data["DisplayMessage"] = changepasswordmsg
			c.Data["title"] = "Success !"
			c.Data["type"] = "success"
			c.TplName = "general/changePassword/changePassword.html"

			log.Println(beego.AppConfig.String("loglevel"), "Info", "Change Password Page Success")
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

	if err = session.ValidateSession(sess); err != nil {
		sess.SessionRelease(c.Ctx.ResponseWriter)
		session.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}
	defer func() {
		utils.EventLogs(c.Ctx, sess, c.Ctx.Input.Method(), c.Input(), c.Data, err)
		sess.SessionRelease(c.Ctx.ResponseWriter)
	}()

	uid := sess.Get("uid").(string)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "uid - ", uid)

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

	input_oldpassword := c.Input().Get("input_oldpassword")
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "Old Login Password - ", old)

	input_newpassword := c.Input().Get("input_newpassword")
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "New Login Password 1 - ", new1)

	input_confirmnewpassword := c.Input().Get("input_confirmnewpassword")
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "Old Login Password 2 - ", new2)

	if input_newpassword != input_confirmnewpassword {
		log.Println(beego.AppConfig.String("loglevel"), "Error", "New Password Mismatch")
		//err = errors.New("New password and Confirm password can't be different")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_NEW_PASSWORD_AND_CONFIRM_PASSWORD_CAN'T_BE_DIFFERENT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_NEW_PASSWORD_AND_CONFIRM_PASSWORD_CAN'T_BE_DIFFERENT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	err = CheckOldPassword(input_oldpassword, uid)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("User  old password Incorrect")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_USER_OLD_PASSWORD_INCORRECT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_USER_OLD_PASSWORD_INCORRECT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	err = UpdatePassword(uid, input_newpassword)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Change Password Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_CHANGE_PASSWORD_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_CHANGE_PASSWORD_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	// err = get.Authenticate(uname.(string), old, "CRM")

	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	// 	err = errors.New("Admin Authentication Failed")
	// 	return
	// }

	//service := "CHANGEPASSWORD"

	// err = get.ValidateRoleAssignment(uname.(string), service)

	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	// 	err = errors.New("Admin Role Validation Failed")
	// 	return
	// }

	// err = get.UpdatePassword(uname.(string), new1)

	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	// 	err = errors.New("Admin Update Password Failed")
	// 	return
	// }
	// return

	if language == "english" {

		changepasswordmsg = beego.AppConfig.String("EN_PASSWORD_CHANGED_SUCCESSFULLY")

	} else if language == "french" {

		changepasswordmsg = beego.AppConfig.String("FN_PASSWORD_CHANGED_SUCCESSFULLY")

	}
}

func CheckOldPassword(oldpassword, userid string) (err error) {

	err = nil

	row, err := db.Db.Query(`select password from sysuser WHERE id = $1`, userid)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to fetch user")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to fetch user")
		return
	}

	// log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

	cp, err := base64.Decode([]byte(data[0][0]))
	if err != nil {
		err = errors.New("Unable to authenticate user")
		return
	}

	var pbkdf pbkdf2.Pbkdf2
	pbkdf.Itr = 32
	pbkdf.KeyLen = 32
	pbkdf.Plain = []byte(oldpassword)
	pbkdf.Salt = cp[:32]
	pbkdf.Cipher = cp[32:]
	result, err := pbkdf.Compare()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User old password incorrect")
		return
	}
	if !result {
		err = errors.New("User old password incorrect")
		return
	}

	return

}

func UpdatePassword(userid, password string) (err error) {

	pass, err := services.EncryptPassword(password)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	result, err := db.Db.Exec("update sysuser set password=$1 ,password_set=true,password_updated_date=now() where id=$2 ", pass, userid)
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
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "pass", password)
	return
}
