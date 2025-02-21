/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package login

import (
	"checkermakerweb/model/db"
	"checkermakerweb/utils"

	"errors"

	"runtime/debug"
	"strconv"

	"encoding/json"
	"fmt"
	"time"

	"checkermakerweb/session"

	"checkermakerweb/utils/database/sql"

	log "github.com/sirupsen/logrus"

	//"proyava.com/database/sql"
	"checkermakerweb/utils/encoding/base64"
	//"proyava.com/util/log"
	"checkermakerweb/utils/util/pbkdf2"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type Login struct {
	beego.Controller
}

type LoginData struct {
	Uname string `form:"username" valid:"Required"`
	Pass  string `form:"password" valid:"Required;MinSize(6);MaxSize(16)"`
}

type roles struct {

	// defining struct variables
	Menu    string          `json:"menu,omitempty"`
	Submenu string          `json:"submenu,omitempty"`
	Data    json.RawMessage `json:"data"`
}

func (c *Login) Get() {
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

func (c *Login) Post() {

	msg := ""

	log.Println(beego.AppConfig.String("loglevel"), "Info", "Login Start")
	pip := c.Ctx.Input.IP()
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Client IP - ", pip)

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
			c.TplName = "general/login/login.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Login Fail")

		} else {

			c.Data["DisplayMessage"] = msg
			c.TplName = "general/login/login.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Login Success")
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

	var l LoginData
	if err := c.ParseForm(&l); err != nil {
		err = errors.New("Invalid Request Received")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Form Data - ", l)
	c.Data["FormData"] = l
	valid := validation.Validation{}
	b, err := valid.Valid(&l)
	if err != nil {
		err = errors.New("Parameter validation failed")
		return
	}

	if !b {
		for _, err := range valid.Errors {
			log.Println(beego.AppConfig.String("loglevel"), "Debug", err.Key, ":", err.Message, ":", err.Field, ":", err.LimitValue, ":", err.Name, ":", err.Tmpl, ":", err.Value)
		}
		err = errors.New("Invalid Input values")
		return
	}

	err = session.CheckUserSession(l.Uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	}

	count, err := getpasscount(l.Uname)
	if err != nil {

		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get count")
		return
	}

	loginCount, _ := beego.AppConfig.Int("LOGIN_COUNT")

	row8, err := db.Db.Query(`SELECT password_set,password_updated_date,language FROM sysuser where  email=$1`, l.Uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get System User password updated date value ")
		return
	}
	defer sql.Close(row8)
	_, data8, err := sql.Scan(row8)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System User password set value scan fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Password Set Value - ", data8[0][0])

	if data8[0][1] == "" {
		fmt.Println("Password Update date value got null")
		err = errors.New("Password Update date value got null")
		return

	}

	lang := data8[0][2]

	//Password need to have expiry time of 30 calendar days, post expiry force Admin to change

	var ts, year, month, day string

	//Createdat1 := data8[0][1]
	//fmt.Println(Createdat1)
	ts = data8[0][1]
	year = ts[0:4]
	month = ts[5:7]
	day = ts[8:10]
	Createdat := year + "-" + month + "-" + day
	fmt.Println("Createdat:=", Createdat)

	y1, _ := strconv.Atoi(year)
	m1, _ := strconv.Atoi(month)
	d1, _ := strconv.Atoi(day)

	t := time.Now()
	y2 := t.Year()  // type int
	m2 := t.Month() // type time.Month
	d2 := t.Day()

	t1 := Date(y1, m1, d1)
	t2 := Date(y2, int(m2), d2)
	days := t2.Sub(t1).Hours() / 24
	fmt.Println(days)

	if days >= 30 {

		fmt.Println("Password expiry time of 30 calendar days")
		if lang == "english" {

			fmt.Println("Your Password is expired after 30 days  , Please Reset it ")

			msg = "Your Password is expired after 30 days  , Please Reset it  "
		} else {

			fmt.Println("Votre mot de passe a expiré après 30 jours, veuillez le réinitialiser")

			msg = "Votre mot de passe a expiré après 30 jours, veuillez le réinitialiser"

		}
		return

	}

	if count >= loginCount {

		result, err := db.Db.Exec(`UPDATE sysuser SET status=$1,updated_at=now() WHERE email=$2`, false, l.Uname)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("user status update fail ")
			return
		}

		i, err := result.RowsAffected()
		if err != nil || i == 0 {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("count fail , no row changed")
			return
		}

		msg = "User Auntentication Fail exceded the Limit 3 times , User is Suspended"

		return

	}

	name, id, user_type, mobile, language, role, menu, submenu, err := authinticate(l.Uname, l.Pass)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Authentication Failed")

		// increment the count
		count++
		// update count
		r_err := PasswordMismatch(count, l.Uname)
		if r_err != nil {

			return
		}

		return
	}

	if count > 0 {
		err = ResetLoginCount(l.Uname)
		if err != nil {
			err = errors.New("Admin User Login Count Reset Failed")
			return
		}
	}

	sess.Set("uname", l.Uname)
	sess.Set("username", name)
	sess.Set("uid", id)
	sess.Set("user_type", string(user_type))
	sess.Set("mobile", mobile)
	sess.Set("language", language)
	sess.Set("role", role)
	sess.Set("menu", menu)
	sess.Set("submenu", submenu)

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "User id :- ", id)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "User Name :- ", name)

	err = session.SetUserSession(sess.SessionID(), "ADMIN"+l.Uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System is unable to process your request.Please contact customer care")
		return
	}
	c.Redirect("/Dashboard", 302)
	return
}

func getpasscount(uname string) (count int, err error) {
	row, err := db.Db.Query("select pass_count from public.\"sysuser\" where email=$1 limit 1", uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Login Count Not Found")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Login Count Scan Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("User Login Count Not Found")
		return
	}

	if data[0][0] == "" {
		data[0][0] = "0"
	}

	count_str := data[0][0]

	count, _ = strconv.Atoi(count_str)
	return

}

func authinticate(uname, pass string) (name, id, user_type, mobile, language, role, menu, submenu string, err error) {
	row, err := db.Db.Query("select id,email,password,fullname,mobile, language,role_id from public.\"sysuser\" where email=$1 and status=$2", uname, true)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to authenticate user")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to authenticate user")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", " User Query Data - ", data, "Data len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("User not registered")
		return
	}

	cp, err := base64.Decode([]byte(data[0][2]))
	if err != nil {
		err = errors.New("Unable to authenticate user")
		return
	}

	var pbkdf pbkdf2.Pbkdf2
	pbkdf.Itr = 32
	pbkdf.KeyLen = 32
	pbkdf.Plain = []byte(pass)
	pbkdf.Salt = cp[:32]
	pbkdf.Cipher = cp[32:]
	result, err := pbkdf.Compare()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User password incorrect")
		return
	}
	if !result {
		err = errors.New("User password incorrect")
		return
	}

	id = data[0][0]
	name = data[0][3]
	user_type = data[0][1]
	mobile = data[0][4]
	language = data[0][5]

	row1, err := db.Db.Query("select id,role_name,privilege::json->'Menus' as menu,privilege::json->'Submenus' as submenu from roles where id=$1", data[0][6])
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to authenticate user")
		return
	}
	defer sql.Close(row1)
	_, data1, err := sql.Scan(row1)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to authenticate user")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data1))

	if len(data1) <= 0 {
		err = errors.New("User not registered")
		return
	}

	role = data1[0][1]
	menu = data1[0][2]
	submenu = data1[0][3]
	return
}

func PasswordMismatch(count int, uname string) (err error) {
	result, err := db.Db.Exec("update sysuser set pass_count=$1 where email=$2 ", count, uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New(" User Count Update Fail")
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New(" User Count Update Fail")
		return
	}

	if n != 1 {
		err = errors.New(" User Count Update Fail")
		return
	}

	return
}

func ResetLoginCount(uname string) (err error) {
	count := 0
	result, err := db.Db.Exec("UPDATE sysuser set pass_count=$1 where email=$2 ", count, uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System User Count Update Fail")
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System User Count Update Fail")
		return
	}

	if n != 1 {
		err = errors.New("System User Count Update Fail")
		return
	}

	return
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
