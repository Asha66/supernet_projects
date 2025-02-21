/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package services

import (
	"strings"

	"errors"
	"io/ioutil"
	"net/mail"
	"net/smtp"

	"checkermakerweb/model/db"
	"fmt"

	"checkermakerweb/utils/database/sql"

	"github.com/scorredoira/email"

	"checkermakerweb/utils/encoding/base64"

	log "github.com/sirupsen/logrus"

	"checkermakerweb/utils/util/pbkdf2"

	"crypto/rand"

	"github.com/astaxie/beego"
	//"proyava.com/encoding/base64"
	//"royava.com/util/pbkdf2"
)

type unencryptedAuth struct {
	smtp.Auth
}

func SendEmail(emilid, name, password, emailtemplate string) {
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "called - ")

	uname := beego.AppConfig.String("EMAIL_NOTIFY_USERNAME")
	pass := beego.AppConfig.String("EMAIL_NOTIFY_PASSWORD")
	url := beego.AppConfig.String("EMAIL_NOTIFY_URL")
	to := beego.AppConfig.String("EMAIL_NOTIFY_TIMEOUT")
	loginurl := beego.AppConfig.String("EMAIL_APPLICATION_LOGIN_URL")
	recipients := strings.Split(emilid, "||")

	tmpFile := emailtemplate

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

	m := email.NewHTMLMessage("Email", msg)
	m.From = mail.Address{Name: "OMINAYA", Address: uname}
	m.To = recipients

	// send it
	//auth := smtp.PlainAuth("", uname, pass, url)

	config := beego.AppConfig.String("EMAIL_AUTH_CONFIG_MODE")

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

func EncryptPassword(pass string) (out []byte, err error) {

	//commenting display of password in logs
	//log.Println(beego.AppConfig.String("loglevel"), "Debug", "inside encryption password rec: ", pass)
	b := make([]byte, 32)
	_, err = rand.Read(b)
	var pbkdf pbkdf2.Pbkdf2
	pbkdf.Itr = 32
	pbkdf.KeyLen = 32
	pbkdf.Plain = []byte(pass)
	pbkdf.Salt = b
	err = pbkdf.Encrypt()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to encrypt password")
		return
	}
	var tmp []byte
	tmp = append(tmp, pbkdf.Salt...)
	tmp = append(tmp, pbkdf.Cipher...)

	out, err = base64.Encode(tmp)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to encrypt password")
		return
	}
	//	log.Printf("%s %s ", "inside encrypt pass after encryption:", out)
	return
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func GetStatus() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from status`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Status Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Status  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT users.id,users.full_name,usertypes.name from users
	left join usertypes on users.user_type=usertypes.id where usertypes.name!='AGENT' AND usertypes.name!='CUSTOMER' `)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Service Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Service  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUsers3() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT users.id,users.full_name from users
	left join usertypes on users.user_type=usertypes.id where usertypes.name='BILLER'`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Service Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Service  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType1() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from usertypes where name!='AGENT'`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType5() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,full_name from users`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType3() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from usertypes where name!='CUSTOMER' AND name!='DISTRIBUTOR' AND name!='AGENT'`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType4() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from usertypes where name='DISTRIBUTOR'`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType2(id string) (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from usertypes where id=$1`, id)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType7() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from usertypes where name='AGENT'`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUserType8() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from usertypes where name='CUSTOMER' OR name='AGENT'`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("User Type  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetUsers() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,full_name from users where status=true`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Users Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Users  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}
func SearchSystemUsers() (data [][]string, err error) {

	row, err := db.Db.Query(`select sysuser.id,sysuser.fullname,sysuser.mobile,sysuser.email,sysuser.address,sysuser.status,sysuser.created_at,roles.role_name from sysuser
	left join roles on roles.id=sysuser.role_id `)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get SystemUser info")
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
		err = errors.New("Unable to get the SystemUser info")
		return

	}

	return
}
func SearchSystemUsersByFilter(name, email, from, to, status string) (data [][]string, err error) {

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(from)
	fmt.Println(to)
	fmt.Println("channel got at SearchChannelByFilter is ", status)

	if status == "" {

		row, err := db.Db.Query(`select sysuser.id,sysuser.fullname,sysuser.mobile,sysuser.email,sysuser.address,sysuser.status,sysuser.created_at,roles.role_name from sysuser
	left join roles on roles.id=sysuser.role_id  where (sysuser.fullname='' OR sysuser.fullname like $1) AND (sysuser.email='' OR sysuser.email like $2) 
	AND ($3='' OR TO_DATE(sysuser.created_at::text,'YYYY/MM/DD') >= TO_DATE( $3,'MM/DD/YYYY')) AND ($4='' OR TO_DATE(sysuser.created_at::text,'YYYY/MM/DD') <= TO_DATE( $4,'MM/DD/YYYY'))`, name, email, from, to)

		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get Channel info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Channel Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the Channel info")
			return data, err
		}

	} else {

		var channelstatus bool

		if status == "ACTIVE" {

			channelstatus = true
		} else if status == "INACTIVE" {
			channelstatus = false
		} else {

		}

		log.Println(beego.AppConfig.String("loglevel"), "Debug", "channelstatus", channelstatus)

		row, err := db.Db.Query(`select sysuser.id,sysuser.fullname,sysuser.mobile,sysuser.email,sysuser.address,sysuser.status,sysuser.created_at,roles.role_name from sysuser
	left join roles on roles.id=sysuser.role_id  where (sysuser.fullname='' OR sysuser.fullname like $1) AND (sysuser.email='' OR sysuser.email like $2) And sysuser.status=$3
	AND ($4='' OR TO_DATE(sysuser.created_at::text,'YYYY/MM/DD') >= TO_DATE( $4,'MM/DD/YYYY')) 
	AND ($5='' OR TO_DATE(sysuser.created_at::text,'YYYY/MM/DD') <= TO_DATE( $5,'MM/DD/YYYY'))`, name, email, channelstatus, from, to)

		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get User info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Message Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the User info")
			return data, err
		}
	}
	return

}

func SearchUsers() (data [][]string, err error) {

	row, err := db.Db.Query(`select users.id,
	users.full_name,
	users.mobile,
	users.email,
	users.address1,
	users.status,
	users.created_at,
	usertypes.name,
	users.address2 from users
	left join usertypes on users.user_type=usertypes.id where usertypes.name!='CUSTOMER' AND name!='DISTRIBUTOR' AND name!='AGENT' `)

	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get SystemUser info")
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
		err = errors.New("Unable to get the SystemUser info")
		return

	}

	return
}

func SearchUsersByFilter(name, email, from, to, status string) (data [][]string, err error) {

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(from)
	fmt.Println(to)
	fmt.Println("channel got at SearchChannelByFilter is ", status)

	if status == "" {

		row, err := db.Db.Query(`select users.id,
	users.full_name,
	users.mobile,
	users.email,
	users.address1,
	users.status,
	users.created_at,
	usertypes.name, users.address2 from users
	left join usertypes on users.user_type=usertypes.id  where usertypes.name!='CUSTOMER' AND name!='DISTRIBUTOR' AND name!='AGENT' AND (users.full_name='' OR users.full_name like $1) AND (users.email='' OR users.email like $2) AND ($3='' OR TO_DATE(users.created_at::text,'YYYY/MM/DD') >= TO_DATE( $3,'MM/DD/YYYY')) AND ($4='' OR TO_DATE(users.created_at::text,'YYYY/MM/DD') <= TO_DATE( $4,'MM/DD/YYYY'))`, name, email, from, to)

		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get User info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Message Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the User info")
			return data, err
		}

	} else {

		var channelstatus bool

		if status == "ACTIVE" {

			channelstatus = true
		} else if status == "INACTIVE" {
			channelstatus = false
		} else {

		}

		log.Println(beego.AppConfig.String("loglevel"), "Debug", "channelstatus", channelstatus)

		row, err := db.Db.Query(`select users.id,
	users.full_name,
	users.mobile,
	users.email,
	users.address1,
	users.status,
	users.created_at,
	usertypes.name, users.address2 from users
	left join usertypes on users.user_type=usertypes.id  where usertypes.name!='CUSTOMER' AND name!='DISTRIBUTOR' AND name!='AGENT' AND (users.full_name='' OR users.full_name like $1) AND (users.email='' OR users.email like $2) AND users.status= $3 AND ($4='' OR TO_DATE(users.created_at::text,'YYYY/MM/DD') >= TO_DATE( $4,'MM/DD/YYYY')) AND ($5='' OR TO_DATE(users.created_at::text,'YYYY/MM/DD') <= TO_DATE( $5,'MM/DD/YYYY'))`, name, email, channelstatus, from, to)

		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get User info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Message Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the User info")
			return data, err
		}
	}
	return

}

func SearchStatus() (data [][]string, err error) {
	row, err := db.Db.Query(`select id,name,"desc",created_at from status  `)
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

func SearchStatusByFilter(name, from, to string) (data [][]string, err error) {

	fmt.Println(name)
	fmt.Println(from)
	fmt.Println(to)
	row, err := db.Db.Query(`select 
	id,
	name,
	"desc",
	created_at 
	from status where (name='' OR name like $1) AND ($2='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $2,'MM/DD/YYYY')) AND ($3='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $3,'MM/DD/YYYY'))`, name, from, to)

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

func SearchUsertype() (data [][]string, err error) {
	row, err := db.Db.Query(`select id,name,"desc",created_at ,status from usertypes  `)
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

func SearchUsertypeByFilter(name, from, to, status string) (data [][]string, err error) {

	fmt.Println(name)
	fmt.Println(from)
	fmt.Println(to)

	if status == "" {

		row, err := db.Db.Query(`select 
	id,
	name,
	"desc",
	created_at ,
	status
	from usertypes where (name='' OR name like $1) AND ($2='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $2,'MM/DD/YYYY')) AND ($3='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $3,'MM/DD/YYYY')) `, name, from, to)

		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get the error message info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Message Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the error message info")
			return data, err
		}

	} else {

		var channelstatus bool

		if status == "ACTIVE" {

			channelstatus = true
		} else if status == "INACTIVE" {
			channelstatus = false
		} else {

		}

		log.Println(beego.AppConfig.String("loglevel"), "Debug", "channelstatus", channelstatus)
		row, err := db.Db.Query(`select 
	id,
	name,
	"desc",
	created_at ,
	status
	from usertypes where (name='' OR name like $1) AND status=$2 AND ($3='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $3,'MM/DD/YYYY')) AND ($4='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $4,'MM/DD/YYYY')) `, name, channelstatus, from, to)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get the error message info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Message Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the error message info")
			return data, err
		}
	}
	return

}
func SearchTransactionReport() (data [][]string, err error) {
	row, err := db.Db.Query(`select 
	id,
	created_at,
	request_id,
	amount,
	status,
	service_name,
	operator,
	mobile,
	account_number,
	user_type,
	txn_number,
	channel,
	pay_gateway_ref_number,
	biller_ref_number,
	customer_metadata
	from transactions  `)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get AccountType info")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("AccountType Scan Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("Unable to get AccountType info")
		return
	}
	return

}

func SearchTransactionReportByFilter(mob, name, from, to, status, reqid, operator, txn, accnum, amount, channel, usertype string) (data [][]string, err error) {

	fmt.Println(mob)
	fmt.Println(name)
	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(status)
	fmt.Println(reqid)
	fmt.Println(operator)
	fmt.Println(txn)
	fmt.Println(accnum)
	fmt.Println(channel)
	fmt.Println(usertype)
	//	fmt.Println(amount)
	row, err := db.Db.Query(`select 
	id,
	created_at,
	request_id,
	amount,
	status,
	service_name,
	operator,
	mobile,
	account_number,
	user_type,
	txn_number,
	channel,
	pay_gateway_ref_number,
	biller_ref_number,
	customer_metadata from transactions  where (mobile='' OR mobile like $1) AND (service_name='' OR service_name like $2) AND (status='' OR status like $3) 
	AND (request_id='' OR request_id like $4) AND (operator='' OR operator like $5) AND (txn_number='' OR txn_number like $6) AND (account_number='' OR account_number like $7) 
	AND (amount ISNULL OR cast(amount as text) like $8) AND (channel='' OR channel like $9) AND ($10='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $10,'MM/DD/YYYY')) AND ($11='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $11,'MM/DD/YYYY')) AND  (user_type='' OR user_type like $12) `, mob, name, status, reqid, operator, txn, accnum, amount, channel, from, to, usertype)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get Transaction Report info1")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("AccountType Scan Fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("Unable to get Transaction Report info2")
		return
	}
	return

}

func SearchRole() (data [][]string, err error) {
	row, err := db.Db.Query(`select id,role_name,"privilege",created_at ,status from roles  `)
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

func SearchRoleByFilter(name, from, to, status string) (data [][]string, err error) {

	fmt.Println(name)
	fmt.Println(from)
	fmt.Println(to)

	if status == "" {

		row, err := db.Db.Query(`select 
	id,
	role_name,
	"privilege",
	created_at ,
	status
	from roles where (role_name='' OR role_name like $1) AND status=$2 AND ($3='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $3,'MM/DD/YYYY')) AND ($4='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $4,'MM/DD/YYYY')) `, name, status, from, to)

		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get the error message info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Message Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the error message info")
			return data, err
		}

	} else {

		var channelstatus bool

		if status == "ACTIVE" {

			channelstatus = true
		} else if status == "INACTIVE" {
			channelstatus = false
		} else {

		}

		log.Println(beego.AppConfig.String("loglevel"), "Debug", "channelstatus", channelstatus)
		row, err := db.Db.Query(`select 
	id,
	role_name,
	"privilege",
	created_at ,
	status
	from roles where (role_name='' OR role_name like $1) AND status=$2 AND ($3='' OR TO_DATE(created_at::text,'YYYY/MM/DD') >= TO_DATE( $3,'MM/DD/YYYY')) AND ($4='' OR TO_DATE(created_at::text,'YYYY/MM/DD') <= TO_DATE( $4,'MM/DD/YYYY')) `, name, channelstatus, from, to)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Unable to get the error message info")
			return data, err
		}
		defer sql.Close(row)
		_, data, err = sql.Scan(row)
		if err != nil {
			log.Println(beego.AppConfig.String("loglevel"), "Error", err)
			err = errors.New("Message Scan Fail")
			return data, err
		}
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, " Data len - ", len(data))
		if len(data) <= 0 {
			err = errors.New("Unable to get the error message info")
			return data, err
		}
	}
	return

}
func GetRole() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,role_name from roles`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Status Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Status  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}

func GetService() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from services where status=true`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Service Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Service  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}
func GetChannel() (data [][]string, err error) {
	row, err := db.Db.Query(`SELECT id,name from channels where status=true`)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Channel Name Get Fail")
		return
	}
	defer sql.Close(row)
	_, data, err = sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Channel  Name Scan Fail")
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "Data len - ", len(data))
	return

}
