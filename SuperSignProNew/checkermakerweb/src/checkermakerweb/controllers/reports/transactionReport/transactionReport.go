/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package transactionReport

import (
	"checkermakerweb/session"

	"checkermakerweb/utils"
	"encoding/json"
	"errors"
	"fmt"
	"runtime/debug"

	"strconv"

	"checkermakerweb/model/db"
	"checkermakerweb/services"
	"checkermakerweb/utils/database/sql"

	"strings"

	"github.com/astaxie/beego"

	log "github.com/sirupsen/logrus"
)

type Row struct {
	Id             string
	Timestamp      string
	RequestId      string
	Amount         string
	Status         string
	Service_name   string
	Operator       string
	Mobile         string
	AccountNumber  string
	UserType       string
	PartialId      string
	Timestamp1     string
	TxnNumber      string
	Channel        string
	CustomerMobile string
	CustomerEmail  string
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
	Id      string
	Name    string
	Channel string
}

type TransactionReport struct {
	beego.Controller
}
type Customerdata struct {
	CustomerMobile string `json:"customer_mobile"`
	CustomerEmail  string `json:"customer_email"`
}

func (c *TransactionReport) Get() {
	Utype := c.Ctx.Input.Param(":Utype")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Utype", Utype)
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Search Transactionreport Start")
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
			c.TplName = "reports/transactionReport/transactionReport.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search Transactionreport Page Fail")
		} else {
			c.Data["DisplayMessage"] = ""
			c.TplName = "reports/transactionReport/transactionReport.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search Transactionreport  Page Success")
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
	auth, err := utils.IsAuthorized(role, "reports-menu", "transactionreport-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	data3, err := services.GetService()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Service fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SERVICE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SERVICE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	if len(data3) <= 0 {
		//err = errors.New("Service  Not Found")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SERVICE_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SERVICE_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	var Dis3 Display
	for i := range data3 {
		var d Field1
		d.Id = data3[i][0]
		d.Name = data3[i][1]
		Dis3.Fields1 = append(Dis3.Fields1, d)
	}
	c.Data["Dis3"] = Dis3
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis3)

	// data4, err := services.GetOperator()
	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	// 	err = errors.New("Operator fetch Failed")
	// 	return
	// }
	// if len(data4) <= 0 {
	// 	err = errors.New("Operator  Not Found")
	// 	return
	// }
	// var Dis4 Display
	// for i := range data4 {
	// 	var d Field1
	// 	//d.Id = data4[i][0]
	// 	d.Name = data4[i][0]
	// 	Dis4.Fields1 = append(Dis4.Fields1, d)
	// }
	// c.Data["Dis4"] = Dis4
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis4)

	data5, err := services.GetChannel()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Channel fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_CHANNEL_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_CHANNEL_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	if len(data5) <= 0 {
		//err = errors.New("Channel  Not Found")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_CHANNEL_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_CHANNEL_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	var Dis5 Display
	for i := range data5 {
		var d Field1
		d.Id = data5[i][0]
		d.Name = data5[i][1]
		Dis5.Fields1 = append(Dis5.Fields1, d)
	}
	c.Data["Dis5"] = Dis5
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis5)

	data, err := services.SearchTransactionReport()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Transaction Report Data fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_TRANSACTION_REPORTS_DATA_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_TRANSACTION_REPORTS_DATA_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	var result []Row
	var ts, tdate, ttime string

	var result5 Customerdata

	for i := range data {
		var r Row
		r.Id = data[i][0]
		r.PartialId = r.Id[0:8]
		r.Timestamp1 = data[i][1]
		ts = data[i][1]
		tdate = ts[0:10]
		ttime = ts[11:19]
		r.Timestamp = tdate + " " + ttime
		r.RequestId = data[i][2]
		r.Amount = data[i][3]
		r.Status = data[i][4]
		r.Service_name = data[i][5]
		r.Operator = data[i][6]
		r.Mobile = data[i][7]
		r.AccountNumber = data[i][8]
		r.UserType = data[i][9]
		r.TxnNumber = data[i][10]
		r.Channel = data[i][11]

		if data[i][9] == "AGENT" {

			if err := json.Unmarshal([]byte(data[i][14]), &result5); err != nil {
				// Parse []byte to go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}

			r.CustomerMobile = result5.CustomerMobile
			r.CustomerEmail = result5.CustomerEmail

		} else {

			r.CustomerMobile = "NA"
			r.CustomerEmail = "NA"

		}

		result = append(result, r)

	}
	c.Data["CustomerData"] = result

	row2, err2 := db.Db.Query(`select 
	count(*)
	from account_transactions 
	left join transactions on transactions.txn_number=account_transactions.txn_number
	WHERE transactions.status=$1 and transactions.channel=$2`, "APPROVED", "B2CWebApp")

	if err2 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err2)
		err2 = errors.New("Unable to fetch Count of b2c")
		return
	}
	defer sql.Close(row2)
	_, data2, err2 := sql.Scan(row2)
	if err2 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err2)
		err2 = errors.New("Unable to fetch Count of b2c")
		return
	}

	row3, err3 := db.Db.Query(`select 
	count(*)
	from account_transactions 
	left join transactions on transactions.txn_number=account_transactions.txn_number
	WHERE transactions.status=$1 and transactions.channel=$2`, "APPROVED", "B2BWebApp")

	if err3 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err3)
		err3 = errors.New("Unable to fetch Count of b2b")
		return
	}
	defer sql.Close(row3)
	_, data33, err3 := sql.Scan(row3)
	if err3 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err3)
		err3 = errors.New("Unable to fetch Count of b2b")
		return
	}

	row44, err44 := db.Db.Query(`select 
	count(*)
	from transactions 
	WHERE transactions.status=$1 and transactions.channel=$2`, "APPROVED", "B2BWebApp")

	if err44 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err44)
		err44 = errors.New("Unable to fetch estimated Count of b2b")
		return
	}
	defer sql.Close(row44)
	_, data44, err44 := sql.Scan(row44)
	if err3 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err44)
		err44 = errors.New("Unable to fetch estimated Count of b2b")
		return
	}

	row45, err45 := db.Db.Query(`select 
	count(*)
	from transactions 
	WHERE transactions.status=$1 and transactions.channel=$2`, "APPROVED", "B2CWebApp")

	if err45 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err45)
		err45 = errors.New("Unable to fetch estimated Count of b2c")
		return
	}
	defer sql.Close(row45)
	_, data45, err44 := sql.Scan(row45)
	if err45 != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err45)
		err45 = errors.New("Unable to fetch estimated Count of b2c")
		return
	}

	b2b, err44 := strconv.Atoi(data44[0][0])
	fmt.Println(err44)

	b2c, err45 := strconv.Atoi(data45[0][0])
	fmt.Println(err45)

	fmt.Println("----------------------------------------")

	fmt.Println("Total Account Transactions Entry for B2C : ", data2[0][0])
	fmt.Println("Estimated Account Transactions Entry for B2C : ", b2c*9)
	fmt.Println("Total Account Transactions Entry for B2B : ", data33[0][0])
	fmt.Println("Estimated Account Transactions Entry for B2B : ", b2b*13)
	fmt.Println("----------------------------------------")

	return
}

func (c *TransactionReport) Post() {
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
			c.TplName = "reports/transactionReport/transactionReport.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search Transactionreport Page Fail")
		} else {
			c.Data["DisplayMessage"] = " "
			c.TplName = "reports/transactionReport/transactionReport.html"
			log.Println(beego.AppConfig.String("loglevel"), "Info", "Search Transactionreport  Page Success")
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
	auth, err := utils.IsAuthorized(role, "reports-menu", "transactionreport-active")
	if !auth {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "UnAuthorized")
		Autherr = errors.New("UnAuthorized")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "IsAuthorized - ", "Authorized")

	data5, err := services.GetChannel()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Channel fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_CHANNEL_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_CHANNEL_CREATION_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	if len(data5) <= 0 {
		//err = errors.New("Channel  Not Found")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_CHANNEL_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_CHANNEL_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	var Dis5 Display
	for i := range data5 {
		var d Field1
		d.Id = data5[i][0]
		d.Name = data5[i][1]
		Dis5.Fields1 = append(Dis5.Fields1, d)
	}
	c.Data["Dis5"] = Dis5
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis5)

	data3, err := services.GetService()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Service fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SERVICE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SERVICE_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	if len(data3) <= 0 {
		//err = errors.New("Service  Not Found")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_SERVICE_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_SERVICE_NOT_FOUND"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}
	var Dis3 Display
	for i := range data3 {
		var d Field1
		d.Id = data3[i][0]
		d.Name = data3[i][1]
		Dis3.Fields1 = append(Dis3.Fields1, d)
	}
	c.Data["Dis3"] = Dis3
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis3)

	// data4, err := services.GetOperator()
	// if err != nil {
	// 	log.Println(beego.AppConfig.String("loglevel"), "Error", err)
	// 	err = errors.New("Operator fetch Failed")
	// 	return
	// }
	// if len(data4) <= 0 {
	// 	err = errors.New("Operator  Not Found")
	// 	return
	// }
	// var Dis4 Display
	// for i := range data4 {
	// 	var d Field1
	// 	//d.Id = data4[i][0]
	// 	d.Name = data4[i][0]
	// 	Dis4.Fields1 = append(Dis4.Fields1, d)
	// }
	// c.Data["Dis4"] = Dis4
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "Object Data - ", Dis4)

	input_name := c.Input().Get("input_name")
	//input_description := c.Input().Get("input_description")
	input_status := c.Input().Get("input_status")
	input_mobile_number := c.Input().Get("input_mobile_number")
	input_request_id := c.Input().Get("input_request_id")
	input_operator := c.Input().Get("input_operator")
	input_acc_number := c.Input().Get("input_acc_number")
	input_amount := c.Input().Get("input_amount")
	input_txn_number := c.Input().Get("input_txn_number")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "input_txn_number - ", input_txn_number)
	input_channel := c.Input().Get("input_channel")

	input_usertype := c.Input().Get("input_usertype")

	// var channelstatus bool

	// if input_status == "ACTIVE" {

	// 	channelstatus = true
	// } else if input_status == "INACTIVE" {
	// 	channelstatus = false
	// } else {

	// }

	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "channelstatus", channelstatus)

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

	data, err := services.SearchTransactionReportByFilter(input_mobile_number+"%", input_name+"%", from, to, input_status+"%", input_request_id+"%", input_operator+"%", input_txn_number+"%", input_acc_number+"%", input_amount+"%", input_channel+"%", input_usertype+"%")
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		//err = errors.New("Channel fetch Failed")
		if language == "english" {
			err = errors.New(beego.AppConfig.String("EN_TRANSACTION_REPORTS_DATA_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "ENGLISH Error", err)
			return
		} else if language == "french" {
			err = errors.New(beego.AppConfig.String("FN_TRANSACTION_REPORTS_DATA_FETCH_FAILED"))
			log.Println(beego.AppConfig.String("loglevel"), "FRENCH Error", err)
			return
		}
		return
	}

	var result []Row
	var ts, tdate, ttime string

	var result5 Customerdata

	for i := range data {
		var r Row
		r.Id = data[i][0]
		r.PartialId = r.Id[0:8]
		r.Timestamp1 = data[i][1]
		ts = data[i][1]
		tdate = ts[0:10]
		ttime = ts[11:19]
		r.Timestamp = tdate + " " + ttime
		r.RequestId = data[i][2]
		r.Amount = data[i][3]
		r.Status = data[i][4]
		r.Service_name = data[i][5]
		r.Operator = data[i][6]
		r.Mobile = data[i][7]
		r.AccountNumber = data[i][8]
		r.UserType = data[i][9]
		r.TxnNumber = data[i][10]
		r.Channel = data[i][11]

		if data[i][9] == "AGENT" {

			if err := json.Unmarshal([]byte(data[i][14]), &result5); err != nil {
				// Parse []byte to go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}

			r.CustomerMobile = result5.CustomerMobile
			r.CustomerEmail = result5.CustomerEmail

		} else {

			r.CustomerMobile = "NA"
			r.CustomerEmail = "NA"

		}

		result = append(result, r)

	}
	c.Data["CustomerData"] = result

	row1, err := db.Db.Query(`select SUM(amount) as ASum
	from transactions
	where (mobile='' OR mobile like $1) AND (service_name='' OR service_name like $2) AND (status='' OR status like $3) 
	AND (request_id='' OR request_id like $4) AND (operator='' OR operator like $5) AND (txn_number='' OR txn_number like $6) AND (account_number='' OR account_number like $7) 
	AND (amount ISNULL OR cast(amount as text) like $8) 
	AND (channel='' OR channel like $9) 
	AND ($10='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $10,'MM/DD/YYYY')) 
	AND ($11='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $11,'MM/DD/YYYY')) `,
		input_mobile_number+"%", input_name+"%", input_status+"%", input_request_id+"%", input_operator+"%", input_txn_number+"%", input_acc_number+"%", input_amount+"%", input_channel+"%", from, to)

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
	defer sql.Close(row1)
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row1)
	_, data1, err := sql.Scan(row1)
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

	c.Data["TotalDepositAmount"] = data1[0][0]

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "TotalDepositAmount - ", data1[0][0])

	row5, err := db.Db.Query(`select count(amount) as ASum
	from transactions
	where (mobile='' OR mobile like $1) AND (service_name='' OR service_name like $2) AND (status='' OR status like $3) 
	AND (request_id='' OR request_id like $4) AND (operator='' OR operator like $5) AND (txn_number='' OR txn_number like $6) AND (account_number='' OR account_number like $7) 
	AND (amount ISNULL OR cast(amount as text) like $8) 
	AND (channel='' OR channel like $9) 
	AND ($10='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $10,'MM/DD/YYYY')) 
	AND ($11='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $11,'MM/DD/YYYY')) `,
		input_mobile_number+"%", input_name+"%", input_status+"%", input_request_id+"%", input_operator+"%", input_txn_number+"%", input_acc_number+"%", input_amount+"%", input_channel+"%", from, to)
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
	defer sql.Close(row5)
	// log.Println(beego.AppConfig.String("loglevel"), "Debug", "Row Data - ", row5)
	_, data7, err := sql.Scan(row5)
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
	c.Data["TotalCountDeposit"] = data7[0][0]

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "TotalCountDeposit - ", data7[0][0])

	return
}
