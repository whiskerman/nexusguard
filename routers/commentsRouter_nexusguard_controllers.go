package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:DeviceTokenController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:EventListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:EventListController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:EventListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:EventListController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:EventListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:EventListController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:EventListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:EventListController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:EventListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:EventListController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogInController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogInController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogInController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogInController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogInController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogInController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogInController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogInController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogInController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogInController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:LogOutController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesListController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:SitesStatusController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:UserController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:UserController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:UserController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:UserController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:UserController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:UserController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nexusguard/controllers:UserController"] = append(beego.GlobalControllerRouter["nexusguard/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
