# HomeReader
a tiny web server write by golang for home books online reading for linux


useed beego pdf.js epub.js 

usage:
modify /conf/app.conf set home=/xxx/xx and web server port

now support .pdf and .epub files online render

homereader is a binary linux x64 , run as : nohup ./homereader &

=================================================================

以下是中文说明：
本程序编写是本人自用于铁威马F220共享NAS(Linux x64平台都可使用)，方便PC或平板通过浏览器在线浏览NAS硬盘PDF和EPUB电子书。
程序用Go写成，后端使用了beego(https://github.com/astaxie/beego),
前端使用PDF.js（https://github.com/mozilla/pdf.js），EPUB.js(https://github.com/futurepress/epub.js)。

使用：

修改 conf/app.conf 设置NAS电子书主目录 home=/xxx/xx ，web服务端口 。

NAS服务器运行: $> nohup ./homereader &

PC或平板浏览器访问 。

目前编译版本为Linux X64平台 ，浏览器测试 chrome。

作者 steven yue kass307@qq.com 2017-09


