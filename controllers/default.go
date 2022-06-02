package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//
//func (c *MainController) ShowIndex() {
//	c.TplName = "index.tpl"
//}
//
//func (c *MainController) Get() {
//	c.TplName = "register.html"
//}
//
//func (c *MainController) Post() {
//	//1、拿到数据
//	userName := c.GetString("userName")
//	pwd := c.GetString("pwd")
//
//	//2、数据校验
//	if userName == "" || pwd == "" {
//		fmt.Println("数据不能为空")
//		c.Redirect("/register",302)
//		return
//	}
//
//	//3、插入数据库
//	o := orm.NewOrm()
//
//	user := models.User{}
//	user.Name = userName
//	user.Pwd = pwd
//	_, err := o.Insert(&user)
//	if err != nil {
//		fmt.Println("插入数据失败")
//		c.Redirect("/register",302)
//		return
//	}
//
//	//4、返回登陆界面
//	c.Redirect("/login", 302)
//}
//
//func (c *MainController) ShowLogin() {
//	c.TplName = "login.html"
//}
//
//func (c *MainController) HandleLogin() {
//	// 接受数据
//	userName := c.GetString("userName")
//	pwd := c.GetString("pwd")
//
//	if userName == "" || pwd == "" {
//		fmt.Println("输入数据不合法")
//		c.TplName = "login.html"
//		return
//	}
//
//	// 校验
//	o := orm.NewOrm()
//	user := models.User{}
//
//	user.Name = userName
//	err := o.Read(&user, "Name")
//	if err != nil {
//		fmt.Println("用户不存在")
//		c.TplName = "login.html"
//		return
//	}else if user.Pwd != pwd {
//		fmt.Println("密码错误")
//		c.TplName = "login.html"
//		return
//	}
//
//	c.Redirect("/index",302)
//}