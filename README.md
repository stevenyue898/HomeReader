# HomeReader
a tiny web server for home books online reading for linux

write by beego .
useed pdf.js epub.js (pdf.js web/cmaps local files are brocked)

usage:
modify /conf/app.conf set home=/xxx/xx

support .pdf and .epub files online render

homereader is a binary linux x64 , run as : nohup ./homereader &

=================================================================
以下是中文说明：
本程序是私人用于铁威马F220共享NAS电子书，提供浏览器在线浏览硬盘PDF和EPUB电子书。
程序用Go写成，后端使用了beego(https://github.com/astaxie/beego),前端使用PDF.js（https://github.com/mozilla/pdf.js），EPUB.js(https://github.com/futurepress/epub.js)。

使用：
修改 conf/app.conf 设置NAS电子书主目录 home=/xxx/xx ，web服务端口 。
NAS服务器运行: $> nohup ./homereader &
浏览器访问 。

目前编译版本为Linux X64平台 ，浏览器测试 chrome。

作者 steven yue kass307@qq.com 2017-09


