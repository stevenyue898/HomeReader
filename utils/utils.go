// utils.gp
package utils

import (
	"errors"
	"fmt"
	"io/ioutil"

	"unicode"

	"strings"

	"github.com/astaxie/beego"
)

var HomeUrl string
var BookUrl string

func init() {
	HomeUrl = beego.AppConfig.String("home")
	BookUrl = "/book"
}

func GetLists(dir string) (ret string) {
	ret = ""
	if !strings.HasPrefix(dir, HomeUrl) {
		return "error"
	}
	base := BookUrl
	if HomeUrl != dir {
		base, _ = GetRel(HomeUrl, dir)
		fmt.Printf("base after getrel: %s \n", base)
		base = BookUrl + "/" + base
	}

	fmt.Printf("GetList: %s\n", base)
	files, dirs, _ := ListDir(dir, ".pdf|.epub")

	for _, i := range dirs {
		ch := show_substr(i, 22)
		//fmt.Printf("%s\n%s\n", i, ch)
		ret = ret + "<a href=\"#\" class=\"list-group-item \" title=\"" + i + "\" data-url=\"" + base + "/" + i + "\" data-type=\"DIR\"><i class=\"fa fa-align-justify icon\"></i>" + ch + "</a>\n"

		//fmt.Printf("%s\n", i)
	} //end of dirs
	for _, i := range files {

		if strings.HasSuffix(strings.ToUpper(i), ".PDF") {
			ret = ret + "<a href=\"#\" class=\"list-group-item \" title=\"" + i + "\" data-url=\"" + base + "/" + i + "\" data-type=\"PDF\"><i class=\"fa fa-file-pdf-o icon\"></i>" + show_substr(i[:len(i)-4], 22) + "</a>\n"
		} else {
			ret = ret + "<a href=\"#\" class=\"list-group-item \" title=\"" + i + "\" data-url=\"" + base + "/" + i + "\" data-type=\"EPUB\"><i class=\"fa fa-book icon\"></i>" + show_substr(i[:len(i)-4], 22) + "</a>\n"
		}

	}
	return ret
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, dirs []string, err error) {
	files = make([]string, 0, 30)
	dirs = make([]string, 0, 30)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	//	PthSep := string(os.PathSeparator)
	suffs := strings.Split(suffix, "|")
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			//continue
			//fmt.Println(fi.Name() + " >>> ")
			//dirs = append(dirs, dirPth+PthSep+fi.Name())
			dirs = append(dirs, fi.Name())
			continue
		}
		for _, suff := range suffs {
			suff = strings.ToUpper(suff)
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suff) { //匹配文件
				//files = append(files, dirPth+PthSep+fi.Name())
				files = append(files, fi.Name())
				break
			}
		}

	}
	return files, dirs, nil
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func show_substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

func show_strlen(s string) int {
	sl := 0
	rs := []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			sl++
		} else {
			sl += 2
		}
	}
	return sl
}

func GetRel(base, targ string) (string, error) {
	//fmt.Printf("base %s , target %s \n", base, targ)
	if !strings.HasPrefix(targ, base) {
		return "error", errors.New("error")
	} else {
		if base == targ {
			return "", nil
		} else {
			return targ[len(base)+1:], nil
		}

	}

}
