/*
Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
*/
package routers

import (
	"checkermakerweb/controllers/email/emailVerfication"
	"checkermakerweb/controllers/error"
	"checkermakerweb/controllers/general/changePassword"
	"checkermakerweb/controllers/general/dashboard"
	"checkermakerweb/controllers/general/forgotpassword"
	"checkermakerweb/controllers/general/login"
	"checkermakerweb/controllers/general/logout"
	"checkermakerweb/controllers/reports/transactionReport"
	"checkermakerweb/controllers/systemconfiguration/role/createRole"
	"checkermakerweb/controllers/systemconfiguration/role/searchRole"
	"checkermakerweb/controllers/systemconfiguration/role/updateRole"
	"checkermakerweb/controllers/systemconfiguration/role/viewRole"
	"checkermakerweb/controllers/systemconfiguration/status/createStatus"
	"checkermakerweb/controllers/systemconfiguration/status/searchStatus"
	"checkermakerweb/controllers/systemconfiguration/status/updateStatus"
	"checkermakerweb/controllers/systemconfiguration/status/viewStatus"
	"checkermakerweb/controllers/systemconfiguration/usertype/createuserType"
	"checkermakerweb/controllers/systemconfiguration/usertype/searchuserType"
	"checkermakerweb/controllers/systemconfiguration/usertype/updateuserType"
	"checkermakerweb/controllers/systemconfiguration/usertype/viewuserType"
	"checkermakerweb/controllers/users/sysusers/createsysUser"
	"checkermakerweb/controllers/users/sysusers/searchsysUser"
	"checkermakerweb/controllers/users/sysusers/updatesysUser"
	"checkermakerweb/controllers/users/sysusers/viewsysUser"
	"checkermakerweb/controllers/users/user/createUser"
	"checkermakerweb/controllers/users/user/searchUser"
	"checkermakerweb/controllers/users/user/updateUser"
	"checkermakerweb/controllers/users/user/viewUser"
	"checkermakerweb/model/db"
	"checkermakerweb/session"

	"github.com/astaxie/beego"
)

func init() {
	if session.Init() != nil {
		return
	}
	if db.Init() != nil {
		return
	}
	beego.ErrorController(&error.Error{})
	beego.SetStaticPath(beego.AppConfig.String("NDASENDA_DOC_DIR_STATIC_URL"), beego.AppConfig.String("NDASENDA_DOC_DIR_STATIC_PATH"))
	beego.SetStaticPath(beego.AppConfig.String("NDASENDA_ADMIN_STATIC_PATH"), beego.AppConfig.String("NDASENDA_ADMIN_STATIC_DIR"))

	beego.SetStaticPath(beego.AppConfig.String("VALIDATION_LANG_PATH"), beego.AppConfig.String("VALIDATION_LANG_PATH"))
	beego.Router(beego.AppConfig.String("MAIN_PATH"), &login.Login{})
	beego.Router(beego.AppConfig.String("LOGIN_PATH"), &login.Login{})
	beego.Router(beego.AppConfig.String("LOGOUT_PATH"), &logout.Logout{})

	beego.Router(beego.AppConfig.String("DASHBOARD_PATH"), &dashboard.Dashboard{})

	beego.Router(beego.AppConfig.String("CHANGE_PASSWORD_PATH"), &changePassword.ChangePassword{})
	beego.Router(beego.AppConfig.String("FORGOT_PASSWORD_PATH"), &forgotpassword.Forgotpassword{})

	beego.Router(beego.AppConfig.String("CREATE_SYS_USER_PATH"), &createsysUser.CreatesysUser{})
	beego.Router(beego.AppConfig.String("SEACRH_SYS_USER_PATH"), &searchsysUser.SearchsysUser{})
	beego.Router(beego.AppConfig.String("UPDATE_SYS_USER_PATH"), &updatesysUser.UpdatesysUser{})
	beego.Router(beego.AppConfig.String("VIEW_SYS_USER_PATH"), &viewsysUser.ViewsysUser{})

	beego.Router(beego.AppConfig.String("CREATE_USER_TYPE_PATH"), &createuserType.CreateuserType{})
	beego.Router(beego.AppConfig.String("SEACRH_USER_TYPE_PATH"), &searchuserType.SearchuserType{})
	beego.Router(beego.AppConfig.String("UPDATE_USER_TYPE_PATH"), &updateuserType.UpdateuserType{})
	beego.Router(beego.AppConfig.String("VIEW_USER_TYPE_PATH"), &viewuserType.ViewuserType{})

	beego.Router(beego.AppConfig.String("CREATE_STATUS_PATH"), &createStatus.CreateStatus{})
	beego.Router(beego.AppConfig.String("SEACRH_STATUS_PATH"), &searchStatus.SearchStatus{})
	beego.Router(beego.AppConfig.String("UPDATE_STATUS_PATH"), &updateStatus.UpdateStatus{})
	beego.Router(beego.AppConfig.String("VIEW_STATUS_PATH"), &viewStatus.ViewStatus{})

	beego.Router(beego.AppConfig.String("CREATE_USER_PATH"), &createUser.CreateUser{})
	beego.Router(beego.AppConfig.String("SEACRH_USER_PATH"), &searchUser.SearchUser{})
	beego.Router(beego.AppConfig.String("UPDATE_USER_PATH"), &updateUser.UpdateUser{})
	beego.Router(beego.AppConfig.String("VIEW_USER_PATH"), &viewUser.ViewUser{})

	beego.Router(beego.AppConfig.String("EMAIL_VERFICATION_PATH"), &emailVerfication.EmailVerfication{})

	beego.Router(beego.AppConfig.String("TRANSACTION_REPORT_PATH"), &transactionReport.TransactionReport{})

	beego.Router(beego.AppConfig.String("CREATE_ROLE_PATH"), &createRole.CreateRole{})
	beego.Router(beego.AppConfig.String("SEACRH_ROLE_PATH"), &searchRole.SearchRole{})
	beego.Router(beego.AppConfig.String("UPDATE_ROLE_PATH"), &updateRole.UpdateRole{})
	beego.Router(beego.AppConfig.String("VIEW_ROLE_PATH"), &viewRole.ViewRole{})

}
