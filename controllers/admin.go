package controllers

import (
	"beego/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"strconv"
	"time"
)

type AdminController struct {
	beego.Controller
}


func (c *AdminController) Prepare() {
	uri := c.Ctx.Request.RequestURI
	if uri != "/admin/login" {
		loginstate := c.GetSession("adminloginstate")
		if loginstate != 1 {
			c.Redirect("/admin/login",302)
			return
		}
	}
}

func (c *AdminController) Index() {
	c.Data["controllerName"] = "首页"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/index.html"
}

func (c *AdminController) Jump(mess string, src string) {
	c.Data["mess"] = mess
	c.Data["src"] = src
	c.TplName = "admin/jump.html"
}

func getTimestamp() int64 {
	timestamp := time.Now().Unix()
	return timestamp
}

func (c *AdminController) Adminlogin() {
	c.TplName = "admin/login.html"
}


func (c *AdminController) Handlelogin() {
	username := c.GetString("username")
	userpwd := c.GetString("userpwd")

	if username != "root" || userpwd ==""{
		c.Jump("用户名或密码不合法","")
		return
	}

	//账号密码验证
	o := orm.NewOrm()
	var uinfo models.Users
	err := o.Raw("SELECT * FROM go_users WHERE phone = ?", username).QueryRow(&uinfo)

	if err != nil {
		logs.Error(err)
	}



	//密码对比
	if uinfo.Password != md5V2(userpwd) {
		c.Jump("密码有误","")
		return
	}

	// beego中使用session 需要先在 app.conf 中添加 sessionon = true
	c.SetSession("adminloginstate", 1)
	c.Redirect("/admin/index",302)
}


func (c *AdminController) Logout() {
	c.DelSession("adminloginstate")
	c.Redirect("/admin/login",302)
}

func (c *AdminController) Setpage() {
	o := orm.NewOrm()
	var options []models.Options
	_, err := o.Raw("SELECT * FROM go_options", ).QueryRows(&options)
	if err == nil {
		c.Data["options"] = options
	}

	c.Data["navSelect"] = 1
	c.Data["controllerName"] = "基本设置"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/set.html"
}

func (c *AdminController) Setpost() {
	webtitle := c.GetString("webtitle")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	siteurl := c.GetString("siteurl")
	admin_email := c.GetString("admin_email")

	o := orm.NewOrm()
	_, err := o.Raw("UPDATE go_options SET val = ? WHERE id = 1", webtitle).Exec()
	_, err1 := o.Raw("UPDATE go_options SET val = ? WHERE id = 4", keywords).Exec()
	_, err2 := o.Raw("UPDATE go_options SET val = ? WHERE id = 5", description).Exec()
	_, err3 := o.Raw("UPDATE go_options SET val = ? WHERE id = 3", siteurl).Exec()
	_, err4 := o.Raw("UPDATE go_options SET val = ? WHERE id = 6", admin_email).Exec()

	if err == nil && err1 == nil &&  err2 == nil && err3 == nil && err4 == nil {
		c.Jump("修改成功", "/admin/set")
	}else{
		c.Jump("修改失败","")
	}
}

func (c *AdminController) Postall() {
	pid, err := c.GetInt("pid")
	if err != nil {
		logs.Error(err)
	}
	//分页处理
	//获得数据总数，总页数，当前页码
	var count uint64
	o := orm.NewOrm()

	if pid > 0 {
		err = o.Raw("SELECT count(*) FROM go_posts WHERE pid = ?", pid).QueryRow(&count)
	}else{
		err = o.Raw("SELECT count(*) FROM go_posts").QueryRow(&count)
	}

	if err != nil {
		logs.Error(err)
	}

	// 每页显示数量
	pagesize := uint64(20)

	index,err := c.GetInt("pageindex")  //当前页码
	if err != nil{
		index = 1
	}

	pageCount := math.Ceil(float64(count) / float64(pagesize))   //总页数

	start := (uint64(index) - 1) * pagesize


	//获取文章列表
	//var lists []orm.ParamsList
	var lists []models.Posts
	var num int64
	if pid > 0 {
		num, err = o.Raw("SELECT * FROM go_posts WHERE pid = ? ORDER BY ID DESC LIMIT ?,?",pid, start, pagesize).QueryRows(&lists)
	}else{
		num, err = o.Raw("SELECT * FROM go_posts ORDER BY ID DESC LIMIT ?,?", start, pagesize).QueryRows(&lists)
	}


	if err == nil && num > 0 {
		//fmt.Println(lists)
		//过滤
		for i, v := range lists {
			vtime , _ := strconv.ParseInt(v.Time, 10 ,64)
			lists[i].Time = time.Unix(vtime , 0).Format("2006-01-02 15:04")
		}
	}

	nextpage := index + 1
	if index >= int(pageCount) {
		nextpage = int(pageCount)
	}

	//获取分类目录
	catedata := Getcatedata()
	c.Data["catedata"] = catedata

	//获取当前分类
	if pid > 0 {
		cateinfo := Getcatenow(pid)
		c.Data["cateinfo"] = cateinfo
	}


	c.Data["pid"] = pid
	c.Data["count"] = count
	c.Data["nextpage"] = nextpage
	c.Data["pageCount"] = int(pageCount)
	c.Data["lists"] = lists
	c.Data["pageIndex"] = index



	c.Data["navSelect"] = 2
	c.Data["controllerName"] = "文章列表"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/postall.html"
}


func (c *AdminController) Postadd() {
	//获取分类目录
	catedata := Getcatedata()
	c.Data["catedata"] = catedata

	c.Data["navSelect"] = 3
	c.Data["controllerName"] = "写文章"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/postadd.html"
}


func (c *AdminController) Postcate() {
	o := orm.NewOrm()
	var lists []models.Postcate
	_, err := o.Raw("SELECT * FROM go_postcate ORDER BY ID DESC").QueryRows(&lists)

	if err != nil {
		logs.Error(err)
	}

	pidname := make(map[int]string)
	for _, v := range lists {
		pidname[v.Id] = v.Name
	}

	for i, v := range lists {
		lists[i].Pname = pidname[v.Pid]
	}


	c.Data["postcate"] = lists

	c.Data["navSelect"] = 4
	c.Data["controllerName"] = "分类目录"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/postcate.html"
}

func (c *AdminController) Postcateadd() {
	name := c.GetString("name")
	pid := c.GetString("pid")

	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO go_postcate (name,pid) VALUES (?,?)", name, pid).Exec()
	if err == nil {
		c.Jump("添加成功","/admin/postcate")
	}else{
		logs.Error(err)
		c.Jump("添加失败","")
	}
}

func (c *AdminController) Postcatemodify() {
	name := c.GetString("name")
	pid := c.GetString("pid")
	cid := c.GetString("cid")

	o := orm.NewOrm()
	var err error
	if pid == "no"{
		_, err = o.Raw("UPDATE go_postcate SET name=? WHERE ID = ?", name, cid).Exec()
	}else{
		_, err = o.Raw("UPDATE go_postcate SET name=?,pid=? WHERE ID = ?", name, pid, cid).Exec()
	}

	if err == nil {
		c.Jump("修改成功","/admin/postcate")
	}else{
		logs.Error(err)
		c.Jump("修改失败","")
	}
}


func (c *AdminController) Postcatesort() {
	fmt.Println("功能暂未实现")
}

func (c *AdminController) Delpostcate() {
	id := c.Ctx.Input.Param(":id")

	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM go_postcate WHERE id = ?",id).Exec()
	if err == nil {
		c.Jump("删除成功","/admin/postcate")
	}else{
		logs.Error(err)
		c.Jump("删除失败","")
	}
}

func Getcatedata() []models.Postcate {
	o := orm.NewOrm()
	var lists []models.Postcate
	_, err := o.Raw("SELECT * FROM go_postcate ORDER BY ID DESC").QueryRows(&lists)

	if err != nil {
		logs.Error(err)
	}

	return lists
}

// 获取当前分类的信息
func Getcatenow (cateid int) models.Postcate {
	if cateid == 0 {
		return models.Postcate{Id: 0, Name: "未分类"}
	}

	o := orm.NewOrm()
	var cinfo models.Postcate
	err := o.Raw("SELECT * FROM go_postcate WHERE ID = ?", cateid).QueryRow(&cinfo)

	if err != nil {
		logs.Error(err)
	}

	//fmt.Println(cinfo)

	return cinfo
}

func (c *AdminController) PushArticle() {
	title := c.GetString("title")
	pid := c.GetString("pid")
	//thumbnail := c.GetString("thumbnail")
	description := c.GetString("description")
	content := c.GetString("content")
	url2 := c.GetString("url2")
	timesql := getTimestamp()

	//处理缩略图上传
	f,h,err := c.GetFile("thumbnail")

	filename := ""

	if err == nil {
		defer f.Close()

		//1.要限定格式
		fileext := path.Ext(h.Filename)
		if fileext != ".jpg" && fileext != ".png" && fileext != ".gif"{
			c.Jump("缩略图文件格式有误","")
			return
		}
		//2.限制大小
		if h.Size > 5*1024 {
			c.Jump("缩略图文件过大","")
			return
		}

		//3.需要对文件重命名，防止文件名重复
		filename = time.Now().Format("20060102150405") + fileext  //6-1-2 3:4:5

		//4、保存, 没有文件夹要先创建
		if err != nil {
			logs.Error("getfile err", err)
		}else{
			c.SaveToFile("thumbnail", "./static/posts/" + filename)
		}
	}


	o := orm.NewOrm()
	_, err = o.Raw("INSERT INTO go_posts (title,pid,thumbnail,description,content,url2,time) VALUES (?,?,?,?,?,?,?)", title,pid,filename,description,content,url2,timesql).Exec()
	if err == nil {
		c.Jump("添加成功","/admin/postall")
	}else{
		logs.Error(err)
		c.Jump("添加失败","")
	}
}


func (c *AdminController) Showmodify() {
	pageid := c.Ctx.Input.Param(":id")

	o := orm.NewOrm()
	postinfo := models.Posts{}
	err := o.Raw("SELECT * FROM go_posts WHERE ID = ?", pageid).QueryRow(&postinfo)
	if err != nil {
		logs.Error(err)
	}

	//获取分类目录
	catedata := Getcatedata()
	c.Data["catedata"] = catedata

	//获取当前分类
	cateinfo := Getcatenow(postinfo.Pid)
	c.Data["cateinfo"] = cateinfo

	c.Data["postinfo"] = postinfo
	c.Data["navSelect"] = 2
	c.Data["controllerName"] = "修改文章"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/postmodify.html"
}

func (c *AdminController) Postmodify() {
	pageid := c.GetString("pageid")
	title := c.GetString("title")
	pid := c.GetString("pid")
	//thumbnail := c.GetString("thumbnail")
	description := c.GetString("description")
	content := c.GetString("content")
	url2 := c.GetString("url2")
	timesql := getTimestamp()

	//处理缩略图上传
	f,h,err := c.GetFile("thumbnail")

	filename := ""

	if err == nil {
		defer f.Close()

		//1.要限定格式
		fileext := path.Ext(h.Filename)
		if fileext != ".jpg" && fileext != ".png" && fileext != ".gif"{
			c.Jump("缩略图文件格式有误","")
			return
		}
		//2.限制大小
		if h.Size > 5*1024 {
			c.Jump("缩略图文件过大","")
			return
		}

		//3.需要对文件重命名，防止文件名重复
		filename = time.Now().Format("20060102150405") + fileext  //6-1-2 3:4:5

		//4、保存, 没有文件夹要先创建
		if err != nil {
			logs.Error("getfile err", err)
		}else{
			c.SaveToFile("thumbnail", "./static/posts/" + filename)
		}
	}


	o := orm.NewOrm()
	_, err = o.Raw("UPDATE go_posts SET title=?,pid=?,thumbnail=?,description=?,content=?,url2=?,ctime=? WHERE id = ?", title,pid,filename,description,content,url2,timesql,pageid).Exec()
	if err == nil {
		c.Jump("修改成功","/admin/postmodify/id/"+pageid)
	}else{
		logs.Error(err)
		c.Jump("修改失败","")
	}
}


func (c *AdminController) Delarticle() {
	pageid := c.Ctx.Input.Param(":id")

	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM go_posts WHERE id = ?",pageid).Exec()
	if err == nil {
		c.Jump("删除成功","/admin/postall")
	}else{
		logs.Error(err)
		c.Jump("删除失败","")
	}
}

func (c *AdminController) Userli() {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM go_users ORDER BY ID DESC").Values(&maps)
	if err == nil && num > 0 {
		c.Data["maps"] = maps
	}


	c.Data["navSelect"] = 5
	c.Data["controllerName"] = "用户列表"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/userli.html"
}

func (c *AdminController) Useraddshow() {
	c.Data["navSelect"] = 6
	c.Data["controllerName"] = "添加用户"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/useradd.html"
}

func md5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func (c *AdminController) Handleuseradd() {
	phone := c.GetString("phone")
	password := c.GetString("password")
	email := c.GetString("email")

	password = md5V2(password)

	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO go_users (phone,password,email) VALUES (?,?,?)", phone, password, email).Exec()
	if err == nil {
		c.Jump("用户添加成功","/admin/userli")
	}else{
		logs.Error(err)
		c.Jump("用户添加失败","")
	}
}

func (c *AdminController) Uinfo() {
	uid, err := c.GetInt("id")
	if err != nil {
		logs.Error(err)
	}

	o := orm.NewOrm()
	var uinfo models.Users
	err = o.Raw("SELECT * FROM go_users WHERE ID = ?", uid).QueryRow(&uinfo)

	if err != nil {
		logs.Error(err)
	}

	c.Data["uinfo"] = uinfo

	c.Data["navSelect"] = 5
	c.Data["controllerName"] = "用户信息"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/uinfo.html"
}
