<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Home Books Center -版权所有 StevenYue kass307@qq.com 2017.09 </title>

  <!-- Bootstrap -->
  <link rel="icon" href="/favicon.ico" />
  <link rel="shortcut icon" href="/favicon.ico"/>
  <link rel="bookmark" href="/favicon.ico"/>
  <link rel="stylesheet" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" href="/static/css/BootSideMenu.css">
  <link href="/static/css/font-awesome.min.css" rel="stylesheet" />
  <link rel="stylesheet" href="/epub/css/basic.css">
  <!-- HTML5 shim e Respond.js per rendere compatibile l'HTML 5 in IE8 -->
  <!-- ATTENZIONE: Respond.js non funziona se la pagina viene visualizzata in locale file:// -->
<!--[if lt IE 9]>
<script src="http://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.2/html5shiv-printshiv.min.js"></script>
<script src="js/respond.min.js"></script>
<![endif]-->

<style type="text/css">
  .user{
    padding:5px;
    margin-bottom: 5px;
  }
   .icon{
    margin-right:2px;
  }
  .center-vertical {
    position:relative;
    top:50%;
    transform:translateY(50%);
    text-align:right;
  }
</style>



</head>
<body>

  <!--Test -->
  <div id="LeftMenu">
    <div class="user" align="center" data-base="{{.base}}">
      <img src="/static/img/if_book_318577.png" width="32px" hight="32px" alt="回到主目录" class="img-thumbnail">
      <a href="/" >.主目录.</a>
    </div>
    <div class="list-group" id="Listmenu">
		
	{{.lists}}
	 

    </div>



  </div>
  <!--/Test -->
 

  <!--Normale contenuto di pagina-->

    <div id="content-iframe " style="background-color: #ffffff; " >
        <div  id="tab-content"  width="100%" height="100%">
  <blockquote class="center-vertical" >
        <p><h2>不愤不启,不悱不发。举一隅,不以三隅反,则吾不复也。</h2></p>
        <small>孔子 <cite title="论语述而">论语</cite></small>
  </blockquote>
        </div>
    </div>
    <!-- /.content -->

  <!--Normale contenuto di pagina-->



  <script src="/static/js/jquery.min.js"></script>
  <script src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/BootSideMenu.js"></script>
  <script src="/epub/build/epub.min.js"></script>
  <script src="/epub/build/libs/zip.min.js"></script>
  <script type="text/javascript">
	  function ListItemClick(){
                     var _type=$(this).data("type");
                     var _url=$(this).data("url");
                     //alert("type: "+_type+" ; URL: "+_url);
                     if (_type=="EPUB")
                     {
								 $("#tab-content").empty()
						 
						 var content =  '<iframe  src="/epub/reader/index.html?bookPath=' + encodeURI(_url) + '" width="100%" height="100%" frameborder="no" border="0" marginwidth="0" marginheight="0" scrolling="yes"  allowtransparency="yes" id="iframe_epub" ></iframe>';
 				         $("#tab-content").append(content);
						 var ifm= document.getElementById("iframe_epub");
                         ifm.height=document.documentElement.clientHeight;
						 //var ifm= $("#area");
						 //ifm.css("top", "10px"); 
                         //ifm.height($(window).height()-10);
                         //Book = ePub(_url);
                         //alert(_url);
	                     //Book.renderTo("area");
	                     $(".toggler").click();
						//$(".stoggler").css("border-color","#000");
						 return;
					 }
                    if (_type=="PDF")
					{
						$("#tab-content").empty()
						var content =  '<iframe  src="/pdfjs/web/viewer.html?file=' + encodeURI(_url) + '" width="100%" height="100%" frameborder="no" border="0" marginwidth="0" marginheight="0" scrolling="yes"  allowtransparency="yes" id="iframe_pdf" ></iframe>';
 				         $("#tab-content").append(content);
						var ifm= document.getElementById("iframe_pdf");
                         ifm.height=document.documentElement.clientHeight;
                         $(".toggler").click();
						return;
					}
					if (_type=="DIR")
					{
						 var _base=$(".user").data("base");
						 
						$.ajax({
 							url: "getDir",
							data:{Base:_base,Dir:_url},
							type: "POST",
							dataType:'json',
							success:function(data) {
						     $("#Listmenu").empty()
							 $("#Listmenu").append(data.list)
							 //$(".user").data("base",data.base);
							 //alert($(".user").data("base"))
							 $(".list-group-item").click(ListItemClick);
                              //alert(data.list)
                              },
						});
                        return;
					}
             };
      $(document).ready(function(){
      $('#LeftMenu').BootSideMenu({side:"left", autoClose:false});
      $(".list-group-item").click(ListItemClick);
      EPUBJS.cssPath = "/epub/reader/css/";
      //$('#test2').BootSideMenu({side:"right"});

  });
  
  </script>
  <footer class="main-footer"  style="top: 480px; text-align=center;display:none;">

 
    版权所有 &copy;StevenYue kass307@qq.com 2017.09 
</footer>
</body>
</html>
