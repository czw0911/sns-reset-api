<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" type="text/css" href="static/css/css.css">
	<script type="text/javascript" src="static/js/jquery.js"></script> 
	<script type="text/javascript" src="static/js/main.js"></script> 
	<title>乡音动态分享</title>
</head>
<body class="body2">
	<dl class="header">
		<dt>
			<h5>乡音</h5>
			<p>同城老乡社交</p>
			<img src="static/img/header2tx.png">
		</dt>
		<dd>
			<a href="javascript:openApp()">立即下载</a>
		</dd>
	</dl>
	<div class="white">
	<dl class="shenying">
		<dt>
			<img src="{{.Dynamic.Images}}">
			<div class="shenying_bottom">
				<h4><img src="static/img/shenying2.png">{{.Dynamic.ViewNUM}}</h4>
				{{if .Dynamic.Voices}}
				<h3><a href="javascript:;" class="myclass"><img src="static/img/shenying1.png"></a></h3>					
				<audio src="{{.Dynamic.Voices}}"></audio>		
				{{end}}		
			</div>			
		</dt>
		<dd>
			<span style="width:60%"></span>
		</dd>
	</dl>

	<p class="localMusicP">{{.Dynamic.DynamicContent}}</p>
	
	<dl class="author">
		<dt>
			<img src="{{.Dynamic.Avatar}}" />			 
			<h5>{{.Dynamic.NickName}}</h5>
			<p>{{.Dynamic.PostTime}}</p>			 
		</dt>
		<dd>
			<a href="javascript:void(0);" class="cur">{{.Dynamic.GoodNUM}}</a>
			<a href="javascript:void(0);">{{.Dynamic.ForwardNum}}</a>
		</dd>		
	</dl>
	<div class="clearboth"></div>
	<div class="discuss">
		<h3>{{.Dynamic.CommentNum}}条评论</h3>
		{{range .Comment}}
			<dl>
				<dt>
					<a href="javascript:void(0)"><img src="{{.Avatar}}"></a>
					<span>{{.PostTime}}</span>
					<h5>{{.NickName}}</h5>				
				</dt>
				<dd>{{.Contents}}</dd>
			</dl>
		{{end}}
	</div>	
</div>
<div class="footer">
	<p>
		<a href="javascript:openApp()">我也来说两句</a>
	</p>
</div>	 
<script type="text/javascript">
audioPlay();
function audioPlay()
{
	var a=$("body.body2 .shenying dt h3 a");
	a.click(function(){
		if($(this).hasClass('myclass'))
		{
			$('audio').trigger('play');
			$(this).removeClass("myclass");
		}
		else
		{
			$('audio').trigger('pause');
			$(this).addClass("myclass");
		}

		
	}) 
}
</script>	 
</body>
{{template "footer.tpl" .}}