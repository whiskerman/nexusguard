package main

import (
	l4g "code.google.com/p/log4go"
	"github.com/astaxie/beego"
	_ "nexusguard/docs"
	_ "nexusguard/routers"
)

func main() {
	l4g.LoadConfiguration("./conf/log4go.xml")
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
