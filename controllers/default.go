package controllers

import (
	"homereader/utils"
	"html/template"

	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type GetDirController struct {
	beego.Controller
}

func (c *MainController) Get() {

	list := utils.GetLists(utils.HomeUrl)

	//fmt.Printf("%s", list)
	c.Data["lists"] = template.HTML(list)
	c.Data["base"] = "/book/"
	c.TplName = "init.tpl"

}
func (c *GetDirController) Post() {
	base := c.GetString("Dir")
	if !strings.HasPrefix(base, utils.BookUrl) {
		c.Redirect("/", 404)
	}

	redir, _ := utils.GetRel(utils.BookUrl, base)
	dir := utils.HomeUrl
	if redir != "" {
		dir = utils.HomeUrl + "/" + redir //文件系统绝对路径
	}

	//utils.BookUrl = c.GetString("Dir") //虚拟路径
	//fmt.Printf("getdir post: %s\n", dir)
	list := utils.GetLists(dir)
	//fmt.Printf("%s", list)
	pdir := getParentDirectory(base)
	//fmt.Printf("虚拟路径:%s\n", pdir)
	if base != utils.BookUrl {
		up := "<a href=\"#\" class=\"list-group-item \"  data-url=\"" + pdir + "\" data-type=\"DIR\"><i class=\"fa fa-align-justify icon\"></i>..回到上级目录</a>\n"
		list = up + list
	}

	c.Data["json"] = map[string]interface{}{"list": list, "base": c.GetString("Dir")}

	c.ServeJSON()

}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	//runes := []rune(s)
	//return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
	return dirctory[:strings.LastIndex(dirctory, "/")]
}
