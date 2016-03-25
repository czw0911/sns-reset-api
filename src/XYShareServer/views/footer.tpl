<script src="http://res.wx.qq.com/open/js/jweixin-1.0.0.js"></script>
<script>

	wx.config({
	    debug: false,
	    appId: '{{.wxMP.appId}}',
	    timestamp: {{.wxMP.timestamp}},
	    nonceStr: '{{.wxMP.nonceStr}}',
	    signature: '{{.wxMP.signature}}',
	    jsApiList: [
	      'onMenuShareTimeline',
		  'onMenuShareAppMessage',
		  'onMenuShareQQ',
		  'onMenuShareWeibo',
		  'onMenuShareQZone',
	    ]
	});
	 wx.ready(function () { 
			wx.onMenuShareTimeline({
			    title: '{{.shareData.title}}', // 分享标题
			    link: '{{.shareData.link}}', // 分享链接
			    imgUrl: '{{.shareData.imgUrl}}', // 分享图标
			    success: function () { 
			        // 用户确认分享后执行的回调函数
			    },
			    cancel: function () { 
			        // 用户取消分享后执行的回调函数
			    }
			});
			
			wx.onMenuShareAppMessage({
			    title: '{{.shareData.title}}', // 分享标题
			    desc: '{{.shareData.desc}}', // 分享描述
			    link: '{{.shareData.link}}', // 分享链接
			    imgUrl: '{{.shareData.imgUrl}}', // 分享图标
			    type: '', // 分享类型,music、video或link，不填默认为link
			    dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
			    success: function () { 
			        // 用户确认分享后执行的回调函数
			    },
			    cancel: function () { 
			        // 用户取消分享后执行的回调函数
			    }
			});
			
			wx.onMenuShareQQ({
			    title: '{{.shareData.title}}', // 分享标题
			    desc: '{{.shareData.desc}}', // 分享描述
			    link: '{{.shareData.link}}', // 分享链接
			    imgUrl: '{{.shareData.imgUrl}}', // 分享图标
			    success: function () { 
			       // 用户确认分享后执行的回调函数
			    },
			    cancel: function () { 
			       // 用户取消分享后执行的回调函数
			    }
			});
			
			wx.onMenuShareWeibo({
			    title: '{{.shareData.title}}', // 分享标题
			    desc: '{{.shareData.desc}}', // 分享描述
			    link: '{{.shareData.link}}', // 分享链接
			    imgUrl: '{{.shareData.imgUrl}}', // 分享图标
			    success: function () { 
			       // 用户确认分享后执行的回调函数
			    },
			    cancel: function () { 
			        // 用户取消分享后执行的回调函数
			    }
			});
			
			wx.onMenuShareQZone({
			    title: '{{.shareData.title}}', // 分享标题
			    desc: '{{.shareData.desc}}', // 分享描述
			    link: '{{.shareData.link}}', // 分享链接
			    imgUrl: '{{.shareData.imgUrl}}', // 分享图标
			    success: function () { 
			       // 用户确认分享后执行的回调函数
			    },
			    cancel: function () { 
			        // 用户取消分享后执行的回调函数
			    }
			});
	});

</script>
</html>