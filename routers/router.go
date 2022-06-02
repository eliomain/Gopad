package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

//分配URL URL变量命名采用全小写
//$this->assign('dashboardurl',U('Admin/Index/index')); //dashboard
//$this->assign('seturl',U('Admin/Index/set')); //设置->基本设置
//$this->assign('postallurl',U('Admin/Index/postall')); //文章->所有文章
//$this->assign('postaddurl',U('Admin/Index/postadd')); //文章->写文章
//$this->assign('postcateurl',U('Admin/Index/postcate')); //文章->分类目录
//
//$this->assign('menucateurl',U('Admin/Index/menucate')); //菜单->菜单分类
//$this->assign('menuaddurl',U('Admin/Index/menuadd')); //菜单->菜单添加
////url
//$this->assign('menuliurl',U('Admin/Index/menuli')); //菜单->菜单列表
//$this->assign('userliurl',U('Admin/Index/userli')); //用户->用户列表
//$this->assign('useraddurl',U('Admin/Index/useradd')); //用户->用户添加

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin/login", &controllers.AdminController{}, "get:Adminlogin;post:Handlelogin")
	beego.Router("/admin/logout", &controllers.AdminController{}, "get:Logout")

	//后台基础
	beego.Router("/admin/index", &controllers.AdminController{}, "get:Index")
	beego.Router("/admin/notice", &controllers.AdminController{}, "get:Jump")
	beego.Router("/admin/set", &controllers.AdminController{}, "get:Setpage;post:Setpost")

	beego.Router("/uecontroller", &controllers.UeditorController{}, "*:ControllerUE")

	// 文章
	beego.Router("/admin/postall", &controllers.AdminController{}, "get:Postall")
	beego.Router("/admin/postadd", &controllers.AdminController{}, "get:Postadd;post:PushArticle")
	beego.Router("/admin/postcate", &controllers.AdminController{}, "get:Postcate")
	beego.Router("/admin/postcateadd", &controllers.AdminController{}, "post:Postcateadd")
	beego.Router("/admin/postcatemodify", &controllers.AdminController{}, "post:Postcatemodify")
	beego.Router("/admin/postcatesort", &controllers.AdminController{}, "post:Postcatesort")

	beego.Router("/admin/delpostcate/id/:id([0-9]+)", &controllers.AdminController{}, "get:Delpostcate")
	beego.Router("/admin/postmodify/id/:id([0-9]+)", &controllers.AdminController{}, "get:Showmodify;post:Postmodify")
	beego.Router("/admin/delarticle/id/:id([0-9]+)", &controllers.AdminController{}, "get:Delarticle")

	// 用户
	beego.Router("/admin/userli", &controllers.AdminController{}, "get:Userli")
	beego.Router("/admin/useradd", &controllers.AdminController{}, "get:Useraddshow;post:Handleuseradd")
	beego.Router("/admin/uinfo", &controllers.AdminController{}, "get:Uinfo")

}