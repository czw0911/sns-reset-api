function openApp(){
//	var has = isWeiXin()
//	if(has){
//		window.location.href= "static/download.html";
//		return
//	}
	window.location.href= "xiangyin://";
	setTimeout( 
	function(){ 
		window.location.href="itms-apps://itunes.apple.com/cn/app/xiang-yin-she-jiao/id1018691120?mt=8";
	}, 30);
}


function isWeiXin(){
    var ua = window.navigator.userAgent.toLowerCase();
    if(ua.match(/MicroMessenger/i) == 'micromessenger'){
        return true;
    }else{
        return false;
    }
}