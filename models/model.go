package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

//表的设计
type Options struct {
	Id int
	Name string
	Val string
}

type Posts struct {
	Id int
	Pid int
	Title string
	Content string
	Thumbnail string
	Description string
	Url2 string
	Time string
	Empty int
}

type Postcate struct {
	Id int
	Name string
	Pname string
	Pid int
	Sort int
}

type Users struct {
	Id int
	Phone string
	Password string
	Email string
	Register_time int
	Login_time int
}

func init(){
	// 设置数据库基本信息
	err := orm.RegisterDataBase("default", "mysql", "root:3vx2ye@tcp(127.0.0.1:3306)/beego?charset=utf8")

	if err != nil {
		fmt.Println("连接数据库失败！", err)
		os.Exit(2)
	}
	logs.Info(fmt.Sprintf("连接数据库成功！"))
	// 映射model数据
	//orm.RegisterModel(new(Go_options))
	// 生成表
	//orm.RunSyncdb("default", false, true)
}