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
 	<dl class="huodong"> 		 
		<dt>
			<img src="{{.Activity.Images}}" onclick="window.location.href='/?a={{.Activity.TalkID}}'">
		</dt>
		<dd>
			<h5>
				点赞<br/>{{.Activity.GoodNUM}}
			</h5>
			<h6>
				吐槽<br/>{{.Activity.BadNUM}}
			</h6>
		</dd>
	</dl>
	<div class="username">
		{{.Activity.TalkContent}}
	</div>
	<dl class="author1">
		<dd>
			<a href="javascript:void(0);"><img src="static/img/author1dd.png"></a>			 
		</dd>	
		<dt>
			<img src="{{.Activity.Avatar}}"/>			 
			{{.Activity.PostUser}}	 		 
		</dt>			
	</dl>
	<div class="pinglun">
		<a href="">{{.Activity.CommentUser}}：</a>{{.Activity.LastComment}} <span>{{.Activity.CommentNUM}}评论</span>
	</div> 
</div>
<div class="footer">
	<p>
		<a href="javascript:openApp()">我也来一发</a>
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