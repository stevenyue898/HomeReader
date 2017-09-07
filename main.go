package main

import (
	_ "homereader/routers"
	"homereader/utils"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/book", utils.HomeUrl)
	beego.SetStaticPath("/pdfjs", "pdfjs")
	beego.SetStaticPath("/epub", "epub")
	beego.Run()
}
