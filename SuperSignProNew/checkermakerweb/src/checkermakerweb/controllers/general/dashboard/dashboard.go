/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package dashboard

import (
	"checkermakerweb/session"
	"checkermakerweb/utils"
	"errors"
	"runtime/debug"

	"checkermakerweb/utils/database/sql"

	"checkermakerweb/model/db"
	"strconv"
	"strings"

	"fmt"
	"time"

	"github.com/astaxie/beego"

	log "github.com/sirupsen/logrus"
)

type Dashboard struct {
	beego.Controller
}

type Display struct {
	SystemPool       []SystemPool
	SystemCommision  []SystemCommision
	BillerPool       []BillerPool
	BillerCommision  []BillerCommision
	PGWPool          []PGWPool
	PGWPoolCommision []PGWPoolCommision
}
type SystemPool struct {
	UserName string
	Balance  string
}
type SystemCommision struct {
	UserName string
	Balance  string
}
type BillerPool struct {
	UserName string
	Balance  string
}
type BillerCommision struct {
	UserName string
	Balance  string
}
type PGWPool struct {
	UserName string
	Balance  string
}
type PGWPoolCommision struct {
	UserName string
	Balance  string
}

func (c *Dashboard) Get() {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Dashboard Page Start")
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
			c.TplName = "general/dashboard/dashboard.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Account Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "general/dashboard/dashboard.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Create Account Page Success")
		}
		return
	}()

	//utils.SetHTTPHeader(c.Ctx)

	sess, err := session.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("System is unable to process your request.Please contact customer care")
		sessErr = true
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "UserName - ", sess.Get("uname"))
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Session ID - ", sess.SessionID())
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

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "mobile :- ", mobile)

	language := sess.Get("language").(string)
	c.Data["language"] = language

	if language == "english" {

		//fmt.Println("System Users")

		c.Data["Systemuser"] = "System Users"
		c.Data["Agent"] = "Agents"
		c.Data["Distributor"] = "Distributors"
		c.Data["Customer"] = "Customers"
		c.Data["Approved"] = "Approved"
		c.Data["Pending"] = "Pending"
		c.Data["Declined"] = "Declined"
	} else {

		//fmt.Println("Utilisateurs du système")

		c.Data["Systemuser"] = "Utilisateurs du système"
		c.Data["Agent"] = "Agentes"
		c.Data["Distributor"] = "Distributrices"
		c.Data["Customer"] = "Les clients"
		c.Data["Approved"] = "A approuvé"
		c.Data["Pending"] = "En attente"
		c.Data["Declined"] = "Diminué"

	}

	role := sess.Get("role").(string)
	c.Data["role"] = role

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "role :- ", role)

	menu := sess.Get("menu").(string)
	c.Data["menu"] = menu

	submenu := sess.Get("submenu").(string)
	c.Data["submenu"] = submenu

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "language :- ", language)

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "usertype :- ", user_type)

	auth, err := utils.IsAuthorized(role, "dashboard-menu", "dashboard-menu")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	row8, err := db.Db.Query(`SELECT password_set,password_updated_date FROM sysuser where  email=$1`, user_type)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get System User password set value ")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_SYSTEM_USER_PASSWORD_SET_VALUE"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_SYSTEM_USER_PASSWORD_SET_VALUE"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row8)
	_, data8, err := sql.Scan(row8)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("System User password set value scan fail")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SYSTEM_USER_PASSWORD_SET_VALUE_SCAN_FAIL"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SYSTEM_USER_PASSWORD_SET_VALUE_SCAN_FAIL"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	//log.Println(beego.AppConfig.String("loglevel"), "Debug", "Password Set Value - ", data8[0][0])

	s1 := data8[0][0]
	b1, _ := strconv.ParseBool(s1)

	if b1 == false {

		fmt.Println("Password is not set by system user")
		if language == "english" {

			fmt.Println("Please Reset Your System Password")

			c.Data["PasswordsetMessage"] = "Please Reset Your System Password"
		} else {

			fmt.Println("Veuillez réinitialiser votre mot de passe système")

			c.Data["PasswordsetMessage"] = "Veuillez réinitialiser votre mot de passe système"

		}

	}

	row, err := db.Db.Query(`SELECT count(*)as sysusercount FROM sysuser`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get System User Count data")
		return
	}
	defer sql.Close(row)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get  System User Count data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_SYSTEM_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_SYSTEM_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sysusercount Query Data - ", data, "Data len - ", len(data))

	number, err := strconv.ParseUint(data[0][0], 10, 32)
	finalIntNum := int(number)

	c.Data["sysusercount"] = finalIntNum

	row1, err := db.Db.Query(`SELECT count(*)as sysusercount FROM users
	left join usertypes on usertypes.id= users.user_type where usertypes.name='CUSTOMER'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Users Count data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row1)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row1)
	_, data1, err := sql.Scan(row1)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Users Count data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "customercount Query Data - ", data1, "Data len - ", len(data1))

	number1, err := strconv.ParseUint(data1[0][0], 10, 32)
	finalIntNum1 := int(number1)

	c.Data["customercount"] = finalIntNum1

	row11, err := db.Db.Query(`SELECT count(*)as sysusercount FROM users
	left join usertypes on usertypes.id= users.user_type where usertypes.name='DISTRIBUTOR'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Users Count data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row11)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row11)
	_, data11, err := sql.Scan(row11)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Users Count data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "customercount Query Data - ", data11, "Data len - ", len(data11))

	number11, err := strconv.ParseUint(data11[0][0], 10, 32)
	finalIntNum11 := int(number11)

	c.Data["distributorcount"] = finalIntNum11

	row12, err := db.Db.Query(`SELECT count(*)as sysusercount FROM users
	left join usertypes on usertypes.id= users.user_type where usertypes.name='AGENT'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Users Count data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row12)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row12)
	_, data12, err := sql.Scan(row12)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Users Count data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_COUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "customercount Query Data - ", data12, "Data len - ", len(data12))

	number12, err := strconv.ParseUint(data12[0][0], 10, 32)
	finalIntNum12 := int(number12)

	c.Data["agentcount"] = finalIntNum12

	// System Pool Account

	row2, err := db.Db.Query(`select
	users.full_name,
	accounts.balance
	from accounts
	left join accounttype on accounttype.id= accounts.account_type
	left join usertypes on usertypes.id= accounts.user_type
	left join users on users.id= accounts.user_id
	where accounttype.name='POOLACCOUNT' AND usertypes.name='SYSTEM'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get SystemPoolAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_SYSTEMPOOLACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_SYSTEMPOOLACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row2)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row2)
	_, data2, err := sql.Scan(row2)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get SystemPoolAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_SYSTEMPOOLACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_SYSTEMPOOLACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "SystemPoolAccount Query Data - ", data2, "Data len - ", len(data2))

	var Dis1 Display
	for i := range data2 {
		var d SystemPool
		d.UserName = data2[i][0]
		d.Balance = data2[i][1]
		Dis1.SystemPool = append(Dis1.SystemPool, d)
	}
	c.Data["Dis1"] = Dis1
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "SystemPoolAccount Object Data - ", Dis1)

	// finalIntNum2 := int(number2)

	// c.Data["SystemPoolAccount"] = finalIntNum2

	// System Commission Account

	row3, err := db.Db.Query(`select 
	users.full_name,
	accounts.balance
	from accounts
	left join accounttype on accounttype.id= accounts.account_type
	left join usertypes on usertypes.id= accounts.user_type
	left join users on users.id= accounts.user_id
	where accounttype.name='COMMISSIONACCOUNT' AND usertypes.name='SYSTEM'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get SystemCommisionAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_SYSTEMCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_SYSTEMCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row3)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row3)
	_, data3, err := sql.Scan(row3)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get SystemCommisionAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_SYSTEMCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_SYSTEMCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "SystemCommAccount Query Data - ", data3, "Data len - ", len(data3))

	var Dis2 Display
	for i := range data3 {
		var d SystemCommision
		d.UserName = data3[i][0]
		d.Balance = data3[i][1]
		Dis2.SystemCommision = append(Dis2.SystemCommision, d)
	}
	c.Data["Dis2"] = Dis2
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "SystemCommAccount Object Data - ", Dis2)

	// number3, err := strconv.ParseUint(data3[0][0], 10, 32)
	// finalIntNum3 := int(number3)
	// c.Data["SystemCommAccount"] = finalIntNum3

	// Biller Pool Account

	row4, err := db.Db.Query(`select 
	users.full_name,
	accounts.balance
	from accounts
	left join accounttype on accounttype.id= accounts.account_type 
	left join usertypes on usertypes.id= accounts.user_type
	left join users on users.id= accounts.user_id
	where accounttype.name='POOLACCOUNT' AND usertypes.name='BILLER'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get BillerPoolAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_BILLERPOOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_BILLERPOOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row4)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row4)
	_, data4, err := sql.Scan(row4)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get BillerPoolAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_BILLERPOOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_BILLERPOOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "BillerPoolAccount Query Data - ", data4, "Data len - ", len(data4))

	var Dis3 Display
	for i := range data4 {
		var d BillerPool
		d.UserName = data4[i][0]
		d.Balance = data4[i][1]
		Dis3.BillerPool = append(Dis3.BillerPool, d)
	}
	c.Data["Dis3"] = Dis3
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "BillerPoolAccount Object Data - ", Dis3)

	// number4, err := strconv.ParseUint(data4[0][0], 10, 32)
	// finalIntNum4 := int(number4)

	// c.Data["BillerPoolAccount"] = finalIntNum4

	// Biller Commission Account

	row5, err := db.Db.Query(`select 
	users.full_name,
	accounts.balance
	from accounts
	left join accounttype on accounttype.id= accounts.account_type 
	left join usertypes on usertypes.id= accounts.user_type
	left join users on users.id= accounts.user_id
	where accounttype.name='COMMISSIONACCOUNT' AND usertypes.name='BILLER'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get BillerCommisionAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_BILLERCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_BILLERCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row5)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row5)
	_, data5, err := sql.Scan(row5)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get BillerCommisionAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_BILLERCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_BILLERCOMMISSIONACCOUNT_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "BillerCommAccount Query Data - ", data5, "Data len - ", len(data4))

	var Dis4 Display
	for i := range data5 {
		var d BillerCommision
		d.UserName = data5[i][0]
		d.Balance = data5[i][1]
		Dis4.BillerCommision = append(Dis4.BillerCommision, d)
	}
	c.Data["Dis4"] = Dis4
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "BillerCommAccount Object Data - ", Dis4)

	// number5, err := strconv.ParseUint(data5[0][0], 10, 32)
	// finalIntNum5 := int(number5)

	// c.Data["BillerCommAccount"] = finalIntNum5

	// PGS Pool Account

	row6, err := db.Db.Query(`select 
	users.full_name,
	accounts.balance
	from accounts
	left join accounttype on accounttype.id= accounts.account_type
	left join usertypes on usertypes.id= accounts.user_type
	left join users on users.id= accounts.user_id
	where accounttype.name='POOLACCOUNT' AND usertypes.name='PAYMENTGATEWAY'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Payment Gateway PoolAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_POOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_POOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row6)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row6)
	_, data6, err := sql.Scan(row6)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Payment Gateway PoolAccounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_POOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_POOLACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "PGSPoolAccount Query Data - ", data6, "Data len - ", len(data6))

	var Dis5 Display
	for i := range data6 {
		var d PGWPool
		d.UserName = data6[i][0]
		d.Balance = data6[i][1]
		Dis5.PGWPool = append(Dis5.PGWPool, d)
	}
	c.Data["Dis5"] = Dis5
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "PGSPoolAccount Object Data - ", Dis5)

	// number6, err := strconv.ParseUint(data6[0][0], 10, 32)
	// finalIntNum6 := int(number6)

	// c.Data["PGSPoolAccount"] = finalIntNum6

	// PGS Commission Account

	row7, err := db.Db.Query(`select 
	users.full_name,
	accounts.balance
	from accounts
	left join accounttype on accounttype.id= accounts.account_type
	left join usertypes on usertypes.id= accounts.user_type
	left join users on users.id= accounts.user_id
	where accounttype.name='COMMISSIONACCOUNT' AND usertypes.name='PAYMENTGATEWAY'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Payment Gateway Commision Accounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_COMMISSIONACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_COMMISSIONACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row7)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row7)
	_, data7, err := sql.Scan(row7)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Payment Gateway Commision Accounts data")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_COMMISSIONACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_USER_PAYMENT_GATEWAY_COMMISSIONACCOUNTS_DATA"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "PGSCommAccount Query Data - ", data7, "Data len - ", len(data7))

	var Dis6 Display
	for i := range data7 {
		var d PGWPoolCommision
		d.UserName = data7[i][0]
		d.Balance = data7[i][1]
		Dis6.PGWPoolCommision = append(Dis6.PGWPoolCommision, d)
	}
	c.Data["Dis6"] = Dis6
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "PGSCommAccount Object Data - ", Dis6)

	// number7, err := strconv.ParseUint(data7[0][0], 10, 32)
	// finalIntNum7 := int(number7)

	// c.Data["PGSCommAccount"] = finalIntNum7

	// --- APPROVED ---

	row10, err := db.Db.Query(`SELECT count(*)as acount FROM transactions where status = 'APPROVED'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Approved Count")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_APPROVED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_APPROVED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row10)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row10)
	_, data10, err := sql.Scan(row10)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get  Approved Count")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_APPROVED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_APPROVED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sysusercount Query Data - ", data10, "Data len - ", len(data10))

	number10, err := strconv.ParseUint(data10[0][0], 10, 32)
	finalIntNum10 := int(number10)

	c.Data["approvedcount"] = finalIntNum10

	// --- PENDING ---

	row20, err := db.Db.Query(`SELECT count(*)as acount FROM transactions where status = 'PENDING'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get System User Count")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_SYSTEM_USER_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_SYSTEM_USER_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row20)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row20)
	_, data20, err := sql.Scan(row20)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get  Pending Count")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_PENDING_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_PENDING_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sysusercount Query Data - ", data20, "Data len - ", len(data20))

	number20, err := strconv.ParseUint(data20[0][0], 10, 32)
	finalIntNum20 := int(number20)

	c.Data["pendingcount"] = finalIntNum20

	// --- DECLINED ---

	row50, err := db.Db.Query(`SELECT count(*)as acount FROM transactions where status = 'DECLINED'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Declined Count")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_DECLINED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_DECLINED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row50)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row50)
	_, data50, err := sql.Scan(row50)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get  Declined Count ")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_DECLINED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_DECLINED_COUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sysusercount Query Data - ", data50, "Data len - ", len(data50))

	number50, err := strconv.ParseUint(data50[0][0], 10, 32)
	finalIntNum50 := int(number50)

	c.Data["declinedcount"] = finalIntNum50
	// --- APPROVED Amount ---

	row60, err := db.Db.Query(`select SUM(amount) FROM transactions as ASum where status = 'APPROVED'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Approved Amount ")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_APPROVED_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_APPROVED_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row60)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row60)
	_, data60, err := sql.Scan(row60)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Approved Amount ")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_APPROVED_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_APPROVED_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sysusercount Query Data - ", data60, "Data len - ", len(data60))

	// number60, err := strconv.ParseUint(data60[0][0], 10, 32)
	// finalIntNum60 := int(number60)

	c.Data["approvedamount"] = data60[0][0]

	// --- PENDING Amount ---

	row70, err := db.Db.Query(`select SUM(amount) FROM transactions as ASum where status = 'PENDING'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Pending Amount ")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_PENDING_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_PENDING_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row70)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row70)
	_, data70, err := sql.Scan(row70)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get  Pending Amount ")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_PENDING_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_PENDING_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sysusercount Query Data - ", data70, "Data len - ", len(data70))

	// number60, err := strconv.ParseUint(data70[0][0], 10, 32)
	// finalIntNum60 := int(number60)

	c.Data["pendingamount"] = data70[0][0]

	// --- DECLINED Amount ---

	row80, err := db.Db.Query(`select SUM(amount) FROM transactions as ASum where status = 'DECLINED'`)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Declined Amount")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_DECLINE_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_DECLINE_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	defer sql.Close(row70)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row80)
	_, data80, err := sql.Scan(row80)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Unable to get Declined Amount")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_UNABLE_TO_GET_DECLINE_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_UNABLE_TO_GET_DECLINE_AMOUNT"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sysusercount Query Data - ", data80, "Data len - ", len(data80))

	// number60, err := strconv.ParseUint(data70[0][0], 10, 32)
	// finalIntNum60 := int(number60)

	c.Data["declinedamount"] = data80[0][0]

	return
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
