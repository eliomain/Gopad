package main

import (
	"beego/models"
	_ "beego/models"
	_ "beego/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("showoptions",options)
	beego.AddFuncMap("showprepage",prepage)
	beego.AddFuncMap("shownextpage",shownextpage)
	beego.SetStaticPath("/plugins", "plugins")
	beego.Run()
}

func options(data models.Options) (res string) {
	res = data.Val
	return
}

func prepage(pageindex int)(preIndex int){
	preIndex = pageindex - 1
	if preIndex <= 1 {
		preIndex = 1
	}
	return
}

func shownextpage(pageindex int)(nextIndex int){
	nextIndex = pageindex + 1
	return
}