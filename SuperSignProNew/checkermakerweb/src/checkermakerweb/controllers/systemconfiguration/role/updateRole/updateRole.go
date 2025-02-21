/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package updateRole

import (
	"checkermakerweb/model/db"
	"checkermakerweb/services"
	"checkermakerweb/session"
	"checkermakerweb/utils"
	"encoding/json"
	"errors"
	"runtime/debug"

	"fmt"
	"strings"

	"checkermakerweb/utils/database/sql"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

type UpdateRole struct {
	beego.Controller
}

type Row struct {
	Id   string
	Name string
	Desc string
}

type Display struct {
	Fields  []Field
	Fields1 []Field1
}
type Field struct {
	Id    string
	Name  string
	Email string
}

type Field1 struct {
	Id   string
	Name string
}
type RBACStruct struct {
	Menus    []string `json:"Menus"`
	Submenus []string `json:"Submenus"`
}

func (c *UpdateRole) Get() {
	AdminId := c.Ctx.Input.Param(":AdminID")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "AdminId", AdminId)
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Usertype Start")
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
			c.TplName = "systemconfiguration/role/updateRole/updateRole.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Role Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "systemconfiguration/role/updateRole/updateRole.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Role Page Success")
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

	auth, err := utils.IsAuthorized(role, "systemmanagement-menu", "searchrole-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	row, err := db.Db.Query(`select id,role_name,status,privilege::json->'Menus' as menu,privilege::json->'Submenus' as submenu from roles where id=$1`, AdminId)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Usertype data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USERTYPE_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USERTYPE_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Usertype data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USERTYPE_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USERTYPE_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	if len(data) <= 0 {
		//err = errors.New("Usertype data not found")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_USERTYPE_DATA_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_USERTYPE_DATA_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

	for i := range data {

		c.Data["Id"] = data[i][0]
		c.Data["Name"] = data[i][1]
		c.Data["Status"] = data[i][2]
		c.Data["Menu"] = data[i][3]
		c.Data["Submenu"] = data[i][4]

	}

	//var menusjson, submenusjson string

	var menuarr []string
	_ = json.Unmarshal([]byte(data[0][3]), &menuarr)
	log.Printf("Unmarshaled: %v", menuarr)

	var submenuarr []string
	_ = json.Unmarshal([]byte(data[0][4]), &submenuarr)
	log.Printf("Unmarshaled: %v", submenuarr)

	c.Data["Menuarray"] = menuarr
	c.Data["Submenuarray"] = submenuarr

	fmt.Println("menu in updated page", menuarr)
	fmt.Println("submenu in updated page", submenuarr)

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
	for i := range data1 {
		var d Field1
		d.Id = data1[i][0]
		d.Name = data1[i][1]
		Dis.Fields1 = append(Dis.Fields1, d)
	}
	c.Data["Dis"] = Dis
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis)

	return
}
func (c *UpdateRole) Post() {
	var usertypemsg string

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
			c.TplName = "systemconfiguration/role/updateRole/updateRole.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Role Page Fail")
		} else {
			c.Data["DisplayMessage"] = usertypemsg
			c.TplName = "systemconfiguration/role/updateRole/updateRole.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Update Role  Page Success")
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
	uid := sess.Get("uid").(string)
	c.Data["uid"] = uid
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

	auth, err := utils.IsAuthorized(role, "systemmanagement-menu", "searchrole-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	// row, err := db.Db.Query(`select id,name,"desc" from usertype where id=$1`, AdminId)
	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	// 	err = errors.New("Unable to get user data")
	// 	return
	// }
	// defer sql.Close(row)
	// _, data, err := sql.Scan(row)
	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	// 	err = errors.New("Unable to get user data")
	// 	return
	// }
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	// if len(data) <= 0 {
	// 	err = errors.New("User data not found")
	// 	return
	// }

	// log.Println(beego.AppConfig.String("loglevel"), "Debug", data)

	// for i := range data {

	// 	c.Data["Id"] = data[i][0]
	// 	c.Data["Name"] = data[i][1]
	// 	c.Data["Desc"] = data[i][2]

	// }

	input_role_name := c.Input().Get("input_role_name")
	input_status := c.Input().Get("input_status")
	fmt.Println("role", input_role_name)
	fmt.Println("status", input_status)
	menuarray := c.Input().Get("menuArr")
	fmt.Println("menu String array from html ", menuarray)
	menuarray1 := strings.Split(menuarray, ",")

	var arraylen = len(menuarray1)
	fmt.Println("menuarray1 arraylen :- ", arraylen)

	submenuArr := c.Input().Get("subArr")
	fmt.Println("submenu String array from html ", submenuArr)
	submenuarray1 := strings.Split(submenuArr, ",")

	var arraylen2 = len(submenuarray1)
	fmt.Println("submenuarray1 arraylen :- ", arraylen2)

	if menuarray == "" || submenuArr == "" {

		if language == "english" {
			err = errors.New("Menu or Submenu cannot be empty")
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			return
		} else {

			err = errors.New("Le menu ou le sous-menu ne peut pas être vide")
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			return

		}

	}

	rbacstruct := RBACStruct{
		Menus:    menuarray1,
		Submenus: submenuarray1,
	}

	var jsonData2 []byte
	jsonData2, err = json.Marshal(rbacstruct)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "RBAC struct  json Data - ", string(jsonData2))

	var channelstatus bool

	if input_status == "ACTIVE" {

		channelstatus = true
	} else {
		channelstatus = false
	}

	res, err := db.Db.Exec(`UPDATE roles SET role_name=$1,status=$2,privilege=$3,updated_by=$4,updated_at=now() WHERE id = $5`, input_role_name, channelstatus, string(jsonData2), uid, AdminId)
	if err != nil {
		//err = errors.New("Usertype updation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_ROLE_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_ROLE_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	i, err := res.RowsAffected()
	if err != nil || i == 0 {
		//err = errors.New("Usertype updation failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_ROLE_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_ROLE_UPDATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	if language == "english" {

		usertypemsg = beego.AppConfig.String("EN_ROLE_UPDATESUCCESSFULLY")

	} else if language == "french" {

		usertypemsg = beego.AppConfig.String("FN_ROLE_UPDATESUCCESSFULLY")

	}

	return
}
