# HomeReader
a tiny web server write by golang for home books (.pdf .epub)online reading for linux
微型电子书浏览服务器

useed beego pdf.js epub.js 

usage:

download homereader.x64.linux.0.1.tar.gz (pdf.js 2.0.4 )here.
  tar zxvf homereader.x64.linux.0.1.tar.gz 
  
modify /conf/app.conf set home=/xxx/xx and web server port
  home is root director of your ebook
  
homereader is a binary execute file for linux x64 , run as : nohup ./homereader &

use your browser to view :
  http://xxx.xxx.xxx.xxx:port

2018.7 now abandon beego ,rewrite used iris-go ,add mp4 movie play by video.js 

  tar zxvf homereader.x64.linux.0.2.iris.tar.gz 

=================================================================

以下是中文说明：
本程序编写原先初衷是本人自用于铁威马F220 NAS平台(Linux x64平台都可使用)，方便PC或平板通过浏览器在线浏览NAS硬盘里PDF和EPUB电子书，后来发现nextCloud早就实现。

程序用Go写成，后端使用了beego(https://github.com/astaxie/beego),后因为增加了mp4播放功能，用iris-go重写。
前端浏览器预览使用PDF.js（https://github.com/mozilla/pdf.js），EPUB.js(https://github.com/futurepress/epub.js)。video.js

使用：

下载 homereader.x64.linux.0.1.tar.gz (pdf.js 2.0.4 )，解压.
  tar zxvf homereader.x64.linux.0.1.tar.gz 

修改 程序主目录下 conf/app.conf 设置NAS电子书主目录 home=/xxx/xx ，及web服务端口 。

NAS服务器运行: $> nohup ./homereader &

使用PC或平板浏览器访问 。
  http://xxx.xxx.xxx.xxx:port
 

目前编译版本为Linux X64平台 ，浏览器测试 chrome。

作者 steven yue kass307@qq.com 2017-09

2018.6 epub 改为epub.js reader ，更改/views/init.tpl即可。
2018.7 homereader.x64.linux.0.2.iris.tar.gz 打包下载（用Iris-go重写，加入video.js 支持mp4 mkv播放）
