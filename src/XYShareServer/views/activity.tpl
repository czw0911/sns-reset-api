<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" type="text/css" href="static/css/css.css">
	<script type="text/javascript" src="static/js/jquery.js"></script>
	<script type="text/javascript" src="static/js/main.js"></script> 	 
	<title>乡音活动分享</title>
</head>
<body class="body2">
	<dl class="header borferbottom">
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
	<div class="title">{{.Activity.ActivityName}}</div>
	<dl class="shenying2">
		<dt>
			<img src="{{.Activity.Images}}" onclick="window.location.href='/?x={{.Activity.TalkID}}'">
			<div class="shenying_bottom">	
			{{if .Activity.Voices}}			 
				<h3><a href="javascript:;" class="myclass"><img src="static/img/shenying1.png"></a></h3>
			{{end}}
				<p>{{.Activity.TalkContent}}</p>	
				{{if .Activity.Voices}}
				<audio src="{{.Activity.Voices}}"></audio>	
				{{end}}			
			</div>			
		</dt>
		<dd>
			<span style="width:60%"></span>
		</dd>
	</dl>
	
	<dl class="author1">
		<dd>
			<a href="javascript:void(0);"><img src="static/img/author1dd.png"></a>			 
		</dd>	
		<dt>
			<img src="{{.Activity.Avatar}}"/>			 
			{{.Activity.PostUser}}		 		 
		</dt>			
	</dl>
	<div class="zans">
		<a href="" class="zan">点赞</a>
		<a href="" class="tucao">吐槽</a>
	</div>	 
	 	
</div>
  
<script type="text/javascript">
audioPlay();
function audioPlay()
{
	var a=$("body.body2 .shenying2 dt h3 a");
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