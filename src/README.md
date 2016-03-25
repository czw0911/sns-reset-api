#    乡音服务端API文档 v1.0.0

## 目录
 
+ 规范约定
		
[字符编码](#1.1)
 
[消息格式](#1.2)
 
[签名算法](#1.3)
 
[URL格式](#1.4)
		
+ API列表
	
> 引导登录
 
[获取短信验证码](#2.1)
 
[获取地区列表](#2.2)

[注册](#2.3)

[登录](#2.4)

[设置新密码](#2.6)

[注册推送id](#2.7)

[获取二级职业列表](#2.8)

[获取标签列表](#2.9)
				
> 活动列表
 
[获取所有活动列表](#3.1)

[获取活动的所有参与内容列表](#3.2)

[发布活动的参与内容](#3.3)

[评论参与内容](#3.4)

[获取参与内容的评论列表](#3.5)

[点赞参与内容](#3.6)

[吐槽参与内容](#3.7)

[获取我的战绩排名](#3.8)

[九宫格随机认证列表](#3.9)

[修改我的乡音](#3.10)

[删除参与的活动](#3.11)

[删除参与活动的评论](#3.12)

> 同乡列表

[获取同乡列表](#4.1)

> 动态

[添加动态](#5.1)

[获取同省动态列表](#5.2)

[点赞动态](#5.3)

[转发动态统计](#5.4)

[评论动态](#5.5)

[获取动态评论列表](#5.6)

[关注用户](#5.7)

[获取用户动态列表](#5.8)

[乡音认证用户](#5.9)

[删除动态](#5.10)

[删除动态评论](#5.11)

> 个人中心

[修改个人信息](#6.1)

[我关注的用户](#6.2)

[关注我的用户](#6.3)

[我参加的活动](#6.4)

[我的认证记录](#6.5)

[消息提醒类型列表](#6.6)

[求认证](#6.7)

[获取求认证列表](#6.8)

[提醒设置](#6.9)

[反馈](#6.10)

[乡音范例](#6.11)

[设置已读提醒类型消息数](#6.12)

[获取提醒设置](#6.13)

[获取系统消息列表（乡音团队)](#6.14)

[获取昵称信息](#6.15)

+ [返回码说明](#retCode) 
 

 
 
## <span id ="1">规范约定</span>

+ <span id ="1.1">字符编码</span>
		
		接口所有字符编码为UTF-8,区分大小写
		
			
+ <span id ="1.2">消息格式</span>
		
		请求格式：
		    标准http get或post格式
			
		返回格式：
			json 格式如下：
				{
					"Code": 1, //返回码
					"Desc":"",//返回码描述
					"Info": anyobject //返回的具体信息
				}
				
+ <span id ="1.3">签名算法</span>

> 用于请求api时的身份认证

		
		签名算法为：hmac_sha256
		public key :不需要用户身份验证的加密key,固定为：f7d8bc1cdb2fdf3c844d0a15026d5ccc25daec39
		private key : 需要用户身份验证的加密key,登录后返回，有效期24小时
		签名字段名为:sign
		签名格式为：
		msg ：签名内容(urlencode的请求参数， 去除sign和二进制数据参数（比如图片，音视频文件等）参数后，按字母排序)
		key = 签名key(默认key或登录后的AccessToken)
		sign = hmac_sha256(msg,key)
			
**[[签名内容必须去除sign和表示二进制数据的参数]]** 
+ <span id ="1.4">URL格式</span>
		URL格式为：
		
		【请求地址】 / 【版本号】/ 【api类型】 / 【api方法] ? 【参数】
		
		例：
		 https://api.xiangyin.im/v1/user/get_info?sid=11111

## API列表

<span id ="2">引导登录</span>
***


+ <span id ="2.1">获取短信验证码</span>

	
		
> *接口说明*

>> 获取短信验证码

> *请求说明*
>> https请求方式： post  
 
>> 		https://api.xiangyin.im/v1/user/get_sms_verifycode 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Account</td>
      <td>是</td>
      <td>手机账号</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
			"Code": 1, 
			"Desc":"",
			"Info": null 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
   <tr>
      <td>Info</td>
      <td>返回具体信息</td>  
   </tr>
</table>

+ <span id ="2.2">获取地区列表</span>

> *接口说明*

>> 获取地区列表

> *请求说明*
>> http请求方式： get  
 
>> 		http://api.xiangyin.im/v1/common/get_region_list?region_type=0&region_id=73


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>RegionType</td>
      <td>是</td>
      <td>地区类型,取值如下<br/>0:获取全部省，市，区县三级地区列表<br/>1:获取省级地区列表<br/>2:获取市级地区列表<br/>3:获取区县地区列表</td>   
   </tr>
   <tr>
      <td>RegionID</td>
      <td>是</td>
      <td>地区编号，取值如下<br/>0:获取地区类型全部列表<br/>其它为获取地区类型和地区编号对应的列表</td>   
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		地区类型为0的返回内容：
		{
			"Code": 1,
			"Desc":"",
			"Info": [
					{
						"RegionID": "1",
						"RegionName": "北京市",
						"RegionSub": [
										{
											"RegionID": "1",
											"RegionName": "北京市",
											"RegionSub": [
															{
																"RegionID": "1",
																"RegionName": "东城区"
															},
				。。。。。。。				
		其它地区类型返回内容：	
		{
			"Code": 1,
			"Desc":"",
			"Info": [
						{
							"RegionID": "9",
							"RegionName": "上海市"
						}
					]
		}					
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
   <tr>
      <td>RegionID</td>
      <td>地区编号</td>  
   </tr>
   <tr>
      <td>RegionName</td>
      <td>地区名称</td>  
   </tr>
   <tr>
      <td>RegionSub</td>
      <td>下级地区</td>  
   </tr>
</table>


+ <span id ="2.3">注册</span>

			
> *接口说明*

>> 账号注册

> *请求说明*
>> https请求方式： post  
 
>> 		https://api.xiangyin.im/v1/user/register 

> *请求参数说明*
>> 
 
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>RegType</td>
      <td>是</td>
      <td>注册类型</td>   
   </tr>
	<tr>
      <td>Account</td>
      <td>是</td>
      <td>手机账号</td>   
   </tr>
	<tr>
      <td>PassWord</td>
      <td>是</td>
      <td>账号登录密码</td>   
   </tr>
	<tr>
      <td>Code</td>
      <td>是</td>
      <td>短信验证码</td>   
   </tr>
	<tr>
      <td>HomeProvinceID</td>
      <td>是</td>
      <td>家乡省编号</td>   
   </tr>
	<tr>
      <td>HomeCityID</td>
      <td>是</td>
      <td>家乡市编号</td>   
   </tr>
	<tr>
      <td>HomeDistrictID</td>
      <td>是</td>
      <td>家乡县编号</td>   
   </tr>
	<tr>
      <td>LivingProvinceID</td>
      <td>是</td>
      <td>现居地省编号</td>   
   </tr>
	<tr>
      <td>LivingCityID</td>
      <td>是</td>
      <td>现居地市编号</td>   
   </tr>
	<tr>
      <td>LivingDistrictID</td>
      <td>是</td>
      <td>现居地县编号</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
			"Code": 1, 
			"Desc":"",
			"Info": null 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
   <tr>
      <td>Info</td>
      <td>返回具体信息</td>  
   </tr>
</table>

+ <span id ="2.4">登录</span>

	
		
> *接口说明*

>> 账号登录

> *请求说明*
>> https请求方式： post   
 
>> 		https://api.xiangyin.im/v1/user/login 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>RegType</td>
      <td>是</td>
      <td>注册类型</td>   
   </tr>
	<tr>
      <td>Account</td>
      <td>是</td>
      <td>手机账号</td>   
   </tr>
	<tr>
      <td>PassWord</td>
      <td>是</td>
      <td>账号登录密码</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": {
				    "AccessToken": "ee59f6611c354fa2a43f8d1e1eaab83f068340ad",
				    "Base": [
				      {
				        "Account": "15900648085",
				        "AuthRecvNum": "0",
				        "AuthSendNum": "0",
				        "BindPhone": "15900648085",
				        "GrowNum": "0",
				        "IsMember": "0",
				        "MedalNum": "0",
				        "RPNum": "0",
				        "RemainDays": "0",
				        "UID": "4228047827"
				      }
				    ],
				    "Detail": [
				      {
				        "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_1.png",
				        "Birthday": null,
				        "DiySign": null,
				        "Gender": null,
				        "HomeCityID": "200001",
				        "HomeDistrictID": "300001",
				        "HomeProvinceID": "100001",
				        "HomeVoice": "",
				        "JobID": null,
				        "LivingCityID": "200001",
				        "LivingDistrictID": "300005",
				        "LivingProvinceID": "100001",
				        "NickName": "三不粘",
				        "ProfessionID": null,
				        "TagID": null,
				        "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_1.png",
				        "VoiceLen": "0"
				      }
				    ]
				  }
				}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
   <tr>
      <td>AccessToken</td>
      <td>访问令牌，用于<a href="#1.3">签名</a></td>  
   </tr>
	<tr>
      <td>Account</td>
      <td>用户账号名</td>  
   </tr>
	<tr>
      <td>AuthRecvNum</td>
      <td>获得乡音认证次数</td>  
   </tr>
	<tr>
      <td>AuthSendNum</td>
      <td>认证他人乡音次数</td>  
   </tr>
	<tr>
      <td>GrowNum</td>
      <td>成长值</td>  
   </tr>
	<tr>
      <td>IsMember</td>
      <td>是否会员</td>  
   </tr>
	<tr>
      <td>MedalNum</td>
      <td>勋章个数</td>  
   </tr><tr>
      <td>RPNum</td>
      <td>人品值</td>  
   </tr><tr>
      <td>RemainDays</td>
      <td>vip剩余天数</td>  
   </tr><tr>
      <td>UID</td>
      <td>用户id</td>  
   </tr><tr>
      <td>Thumbnail</td>
      <td>头像缩略图</td>  
   </tr><tr>
      <td>Avatar</td>
      <td>头像</td>  
   </tr><tr>
      <td>Birthday</td>
      <td>生日(时间戳)</td>  
   </tr><tr>
      <td>BindPhone</td>
      <td>绑定手机号默认（绑定手机账号)</td>  
   </tr><tr>
      <td>DiySign</td>
      <td>个性签名</td>  
   </tr><tr>
      <td>Gender</td>
      <td>性别（1：男，2：女</td>  
   </tr><tr>
      <td>JobID</td>
      <td>职业编号</td>  
   </tr>
	<tr>
      <td>ProfessionID</td>
      <td>职业所属行业编号</td>  
   </tr>
	<tr>
      <td>NickName</td>
      <td>昵称</td>  
   </tr>
	<tr>
      <td>TagID</td>
      <td>标签编号</td>  
   </tr>
	<tr>
      <td>HomeProvinceID</td>
      <td>家乡省编号</td>  
   </tr>
	<tr>
      <td>HomeCityID</td>
      <td>家乡市编号</td>  
   </tr>
	<tr>
      <td>HomeDistrictID</td>
      <td>家乡区县编号</td>  
   </tr>
	<tr>
      <td>LivingProvinceID</td>
      <td>现居地省编号</td>  
   </tr>
	<tr>
      <td>LivingCityID</td>
      <td>现居地市编号</td>  
   </tr>
   <tr>
      <td>LivingDistrictID</td>
      <td>现居地区县编号</td>  
   </tr>
	<tr>
      <td>HomeVoice</td>
      <td>乡音文件地址</td>  
   </tr>
	<tr>
      <td>VoiceLen</td>
      <td>乡音音频长度</td>  
   </tr>
</table>



+ <span id ="2.6">设置新密码</span>

	
		
> *接口说明*

>> 设置新密码

> *请求说明*
>> https请求方式： post  
 
>> 		https://api.xiangyin.im/v1/user/reset_password 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Account</td>
      <td>是</td>
      <td>手机账号</td>   
   </tr>
   <tr>
      <td>PassWord</td>
      <td>是</td>
      <td>新密码</td>   
   </tr>
   <tr>
      <td>Code</td>
      <td>是</td>
      <td>短信验证码</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
			"Code": 1, 
			"Desc":"",
			"Info": null 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
   <tr>
      <td>Info</td>
      <td>返回具体信息</td>  
   </tr>
</table>

+ <span id ="2.7">添加推送id</span>

	
		
> *接口说明*

>> 登录后添加最新的推送id

> *请求说明*
>> https请求方式： post  
 
>> 		https://api.xiangyin.im/v1/user/add_pushid 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>PushType</td>
      <td>是</td>
      <td>推送类型（1：ios，2:android</td>   
   </tr>
   <tr>
      <td>PushID</td>
      <td>是</td>
      <td>推送id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
			"Code": 1, 
			"Desc":"",
			"Info": null 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
   <tr>
      <td>Info</td>
      <td>返回具体信息</td>  
   </tr>
</table>

 
 
+ <span id ="2.8">获取二级职业列表</span>

	
> *接口说明*

>> 获取二级职业列表

> *请求说明*
>> http请求方式： get  
 
>> 		http://api.xiangyin.im/v1/common/get_job_list 


> *请求参数说明*
>>
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>无参数，不需要sign</td>
      <td></td>
      <td></td>   
   </tr>
</table>

 
> *返回说明*
>>  	
>>  	返回内容：		 
		{
			"Code": 1,
			"Desc": "成功",
			"Info": [
				{
				"JobId": "100001",
				"JobName": "IT互联网",
				"JobSub": [
							{
							"JobId": "200001",
							"JobName": "开发"
							},
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
   <tr>
      <td>JobId</td>
      <td>职业编号</td>  
   </tr>
   <tr>
      <td>JobName</td>
      <td>职业名称</td>  
   </tr>
</table>

 
+ <span id ="2.9">获取标签列表</span>

	
> *接口说明*

>> 获取标签列表

> *请求说明*
>> http请求方式： get  
 
>> 		http://api.xiangyin.im/v1/common/get_tags_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>无参数，不需要sign</td>
      <td></td>
      <td></td>   
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
			"Code": 1,
			"Desc": "成功",
			"Info": [
						{
						"TagColor": "0,188,155,1.0",
						"TagID": "900000",
						"TagName": "女神"
						},
						{
						"TagColor": "0,188,155,1.0",
						"TagID": "900001",
						"TagName": "女汉子"
						},			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
   <tr>
      <td>TagID</td>
      <td>标签编号，格式：前2位是标签类型，其余位为流水号</td>  
   </tr>
   <tr>
      <td>TagName</td>
      <td>标签名称</td>  
   </tr>
	<tr>
      <td>TagColor</td>
      <td>标签颜色</td>  
   </tr>
</table>

***
活动列表
***

+ <span id ="3.1">获取所有活动列表</span>

	
		
> *接口说明*

>> 登录后获取所有活动列表

> *请求说明*
>> http请求方式： get  
 
>> 		http://api.xiangyin.im/v1/activity/get_activity_list


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>PageType</td>
      <td>是</td>
      <td>翻页类型（1：上翻页，2:下翻页）</td>   
   </tr>
   <tr>
      <td>MaxID</td>
      <td>是</td>
      <td>返回内容里的最大id.(值为空时与PageType对应,分别为首页和尾页) </td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": {
		    "List": [
		     {
					"ActivityDesImg": "http://127.0.0.1/v1/file/view_file?FileName=2_0_0_0_201506_10000000.png",
					"ActivityID": "10000000",
					"ActivityName": "听乡音 辨家乡",
					"JoinNum": "0",
					"ShareContent": null,
					"ShareImgUrl": null,
					"Url": null
					}
		    ],
		    "MaxID": "5",
			"MinID": "1"
		  }
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
   <tr>
      <td>ActivityDesImg</td>
      <td>活动描述图片</td>  
   </tr>
    <tr>
      <td>ActivityID</td>
      <td>活动编号</td>  
   </tr>
   <tr>
      <td>ActivityName</td>
      <td>活动名称</td>  
   </tr>
 	<tr>
      <td>JoinNum</td>
      <td>活动查看人数</td>  
   </tr>	
	<tr>
      <td>ShareContent</td>
      <td>分享内容</td>  
   </tr>
	<tr>
      <td>ShareImgUrl</td>
      <td>分享地址</td>  
   </tr>
	<tr>
      <td>Url</td>
      <td>网页活动地址</td>  
   </tr>
	<tr>
      <td>MaxID</td>
      <td>最大流水号</td>  
   </tr>
   <tr>
      <td>MinID</td>
      <td>最小流水号</td>  
   </tr>
</table>


+ <span id ="3.2">获取活动的所有参与内容列表</span>

	
		
> *接口说明*

>> 登录后获取活动的所有参与内容列表

> *请求说明*
>> http请求方式： get  
 
>> 		http://api.xiangyin.im/v1/activity/get_talk_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>PageType</td>
      <td>是</td>
      <td>翻页类型（1：上翻页，2:下翻页）</td>   
   </tr>
   <tr>
      <td>MaxID</td>
      <td>是</td>
      <td>返回内容里的最大id.(值为空时与PageType对应,分别为首页和尾页) </td>   
   </tr>
	<tr>
      <td>ActivityID</td>
      <td>是</td>
      <td>活动编号</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": {
				    "List": [
				      {
				        "BadNUM": "3",
				        "CommentNUM": "34",
						"CommentUser": "贴地飞行",
						"LastComment": "gggggggg",
				        "GoodNUM": "3",
				        "Images": "",
				        "IsClickGoodOrBad": 0,
				        "PostTime": "1435287360",
				        "PostUser": {
				          "UID": 4228047827,
				          "NickName": "kill",
				          "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_1.png",
				          "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_1.png",
				          "HomeProvinceID": 1020988,
				          "HomeCityID": 200001,
				          "HomeDistrictID": 300001,
				          "LivingProvinceID": 100001,
				          "LivingCityID": 200001,
				          "LivingDistrictID": 300005,
				          "HomeVoice": "",
				          "VoiceLen": 0,
				          "ProfessionID": 0,
				          "JobID": 0,
				          "IsFollow": 0,
				          "Birthday": 0,
				          "Gender": 0,
				          "IsMember": 0,
				          "IsGuess": 1,
				          "AuthRecvNum": 12,
				          "TagsID": [
				            "666"
				          ],
				          "LastLoginTime": 1435909644
				        },
				        "TalkContent": "ssdfsdgaa",
				        "TalkID": "1434908579916858202100000014228047827",
				        "VoiceLen": "1",
				        "Voices": ""
				      }
				    ],
				    "MaxID": "201506100000014",
				    "MinID": "201506100000011"
				  }
			}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
	<tr>
      <td>BadNUM</td>
      <td>参与内容吐槽数</td>  
   </tr>
	<tr>
      <td>CommentUser</td>
      <td>最后评论的人</td>  
   </tr>
	<tr>
      <td>LastComment</td>
      <td>最后评论的内容</td>  
   </tr>
	<tr>
      <td>CommentNUM</td>
      <td>参与内容评论数</td>  
   </tr>
	<tr>
      <td>GoodNUM</td>
      <td>参与内容点赞数</td>  
   </tr>
	<tr>
      <td>Images</td>
      <td>参与内容图片</td>  
   </tr>
	<tr>
      <td>PostTime</td>
      <td>参与内容发布时间</td>  
   </tr>
	<tr>
      <td>PostUser</td>
      <td>发布人信息</td>  
   </tr>
	<tr>
      <td>UID</td>
      <td>发布人编号</td>  
   </tr>
	<tr>
      <td>NickName</td>
      <td>发布人昵称</td>  
   </tr>
	<tr>
      <td>Avatar</td>
      <td>发布人头像</td>  
   </tr>
	<tr>
      <td>Thumbnail</td>
      <td>发布人头像缩略头像</td>  
   </tr>
	<tr>
      <td>HomeProvinceID</td>
      <td>发布人家乡省</td>  
   </tr>
	<tr>
      <td>HomeCityID</td>
      <td>发布人家乡市</td>  
   </tr>
	<tr>
      <td>HomeDistrictID</td>
      <td>发布人家乡区县</td>  
   </tr>
	<tr>
      <td>LivingProvinceID</td>
      <td>发布人现居地省</td>  
   </tr>
	<tr>
      <td>LivingCityID</td>
      <td>发布人现居地市</td>  
   </tr>
	<tr>
      <td>LivingDistrictID</td>
      <td>发布人现居地县</td>  
   </tr>
	<tr>
      <td>HomeVoice</td>
      <td>发布人乡音</td>  
   </tr>
	<tr>
      <td>VoiceLen</td>
      <td>发布人乡音长度</td>  
   </tr>
	<tr>
      <td>JobID</td>
      <td>发布人职业</td>  
   </tr>
	<tr>
      <td>ProfessionID</td>
      <td>职业所属行业编号</td>  
   </tr>
	<tr>
      <td>IsFollow</td>
      <td>是否已关注发布人（0:未关注；1：已关注；2:被关注;3：互相关注）</td>  
   </tr>
	<tr>
      <td>Birthday</td>
      <td>发布人生日时间戳</td>  
   </tr>
	<tr>
      <td>Gender</td>
      <td>发布人性别</td>  
   </tr>
	<tr>
      <td>IsMember</td>
      <td>发布人是否是会员（0:不是；1：是）</td>  
   </tr>
	<tr>
      <td>IsGuess</td>
      <td>是否已猜测过发布人的乡音</td>  
   </tr>
	<tr>
      <td>AuthRecvNum</td>
      <td>发布人获得乡音认证次数</td>  
   </tr>
	<tr>
      <td>TagsID</td>
      <td>发布人标签数组</td>  
   </tr>
	<tr>
      <td>TalkContent</td>
      <td>参与内容</td>  
   </tr>
	<tr>
      <td>TalkID</td>
      <td>参与内容编号</td>  
   </tr>
	<tr>
      <td>VoiceLen</td>
      <td>上传参与内容的声音长度</td>  
   </tr>
	<tr>
      <td>Voices</td>
      <td>参与内容的声音文件</td>  
   </tr>
	<tr>
      <td>IsClickGoodOrBad</td>
      <td>是否点赞或吐槽过( 0:没有点击过;1点击过点赞；2点击过吐槽)</td>  
   </tr>
	<tr>
      <td>MaxID</td>
      <td>最大流水号</td>  
   </tr>
   <tr>
      <td>MinID</td>
      <td>最小流水号</td>  
   </tr>
</table>

+ <span id ="3.3">发布活动的参与内容</span>

	
		
> *接口说明*

>> 登录后发布活动的参与内容

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/activity/add_talk 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>ActivityID</td>
      <td>是</td>
      <td>活动编号</td>   
   </tr>
   <tr>
      <td>TalkContent</td>
      <td>是</td>
      <td>参与文字内容</td>   
   </tr>
   <tr>
      <td>Images</td>
      <td>是</td>
      <td>参与内容的图片</td>   
   </tr>
   <tr>
      <td>Voices</td>
      <td>是</td>
      <td>参与内容的声音</td>   
   </tr>
 	<tr>
      <td>VoiceLen</td>
      <td>是</td>
      <td>参与内容声音长度</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		      "Code": 1,
			  "Desc": "成功",
			  "Info": {
			          "TalkID": "1435287360240497946100000014228047827"
			          }
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>  
   </tr>
  <tr>
      <td>TalkID</td>
      <td>参与内容编号</td>  
   </tr>
</table>


+ <span id ="3.4">评论参与内容</span>

	
		
> *接口说明*

>> 登录后评论参与内容

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/activity/add_comment 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>TalkID</td>
      <td>是</td>
      <td>参与内容编号</td>   
   </tr>
 <tr>
      <td>Contents</td>
      <td>是</td>
      <td>评论内容</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": ...
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
</table>



+ <span id ="3.5">获取参与内容的评论列表</span>	

> *接口说明*

>> 登录后获取参与内容的评论列表

> *请求说明*
>> http请求方式：   get
 
>> 		http://api.xiangyin.im/v1/activity/get_comment_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>TalkID</td>
      <td>是</td>
      <td>参与内容编号</td>   
   </tr>
 <tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
 <tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：
	    {
		  "Code": 1,
		  "Desc": "成功",
		  "Info": {		 
				"List": [
					      {
								"CommentID": "1434908579916858202100000014228047827_20",
						        "Contents": "teste",
						        "PostTime": "1434964006",
						        "PostUser": {
										....
										}
							}
						],
		    		"MaxID": "10",
					"MinID": "1"
		  	}
		}					
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Contents</td>
      <td>评论内容</a> </td>   
   </tr>
	<tr>
      <td>PostTime</td>
      <td>评论时间 </td>   
   </tr>
	<tr>
      <td>PostUser</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表] </td>   
   </tr>
	<tr>
      <td>CommentID</td>
      <td>评论编号 </td>   
   </tr>
</table>



+ <span id ="3.6">点赞参与内容</span>	

> *接口说明*

>> 登录后点赞参与内容

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/activity/click_good 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>TalkID</td>
      <td>是</td>
      <td>参与内容编号</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		   "Code": 1,
		  "Desc": "成功",
		  "Info":
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>   
   </tr>
</table>


+ <span id ="3.7">吐槽参与内容</span>	

> *接口说明*

>> 登录后吐槽参与内容

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/activity/click_bad 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>TalkID</td>
      <td>是</td>
      <td>参与内容编号</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		  "Code": 1,
		  "Desc": "成功",
		  "Info":
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>   
   </tr>
</table>



+ <span id ="3.8">获取我的战绩排名</span>	

> *接口说明*

>> 登录后获取我的战绩排名

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/user/voice_record_ranking 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		      "Code": 1,
			  "Desc": "成功",
			  "Info": {
			    "AuthSendNum": 0,
			    "AuthRecvNum": 0,
			    "AccuracyRate": 0,
			    "Ranking": 0
			  }
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>AuthSendNum</td>
      <td>认证他人次数</td>   
   </tr>
	<tr>
      <td>AuthRecvNum</td>
      <td>被认证次数</td>   
   </tr>
	<tr>
      <td>AccuracyRate</td>
      <td>正确率</td>   
   </tr>
	<tr>
      <td>Ranking</td>
      <td>排名</td>   
   </tr>
</table>



+ <span id ="3.9">九宫格随机认证列表</span>	

> *接口说明*

>> 登录后九宫格随机认证列表

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/user/voice_sudoku_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": [
			    {
			      "UID": 4228047827,
			      "NickName": "豆捞",
			      "Avatar": "0_0_0_0_0_p01.png",
			      "Thumbnail": "0_0_0_0_0_p01.png",
			      "HomeProvinceID": 1,
			      "HomeCityID": 2,
			      "HomeDistrictID": 3,
			      "LivingProvinceID": 4,
			      "LivingCityID": 5,
			      "LivingDistrictID": 6,
			      "HomeVoice": "",
			      "VoiceLen": 2,
			      "JobID": 0,
			      "IsFollow": 0,
			      "Birthday": 0,
			      "Gender": 0,
			      "IsMember": 0,
			      "IsGuess": 1,
			      "AuthRecvNum": 0,
			      "TagsID": ""
			    }
			  ]
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>其它</td>
      <td> <a href="#3.2">[同获取活动的所有参与内容列表] </td>   
   </tr>
</table>


+ <span id ="3.10">修改我的乡音</span>	

> *接口说明*

>> 登录后修改我的乡音

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/user/update_home_voice 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>HomeVoice</td>
      <td>是</td>
      <td>乡音文件</td>   
   </tr>
	<tr>
      <td>VoiceLen</td>
      <td>是</td>
      <td>乡音长度</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		  "Code": 1,
		  "Desc": "成功",
		  "Info":
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码；详见<a href="#retCode">[返回码说明]</a> </td>   
   </tr>
	<tr>
      <td>Desc</td>
      <td>返回码描述</td>   
   </tr>
</table>

+ <span id ="3.11">删除参与的活动</span>	

> *接口说明*

>> 登录后删除自己参与的活动内容

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/activity/del_talk 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>TalkID</td>
      <td>是</td>
      <td>参与内容编号</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>


+ <span id ="3.12">删除参与活动的评论</span>	

> *接口说明*

>> 登录后删除自己发布的活动评论

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/activity/del_comment 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>CommentID</td>
      <td>是</td>
      <td>评论id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>






***
同乡列表
***

+ <span id ="4.1">获取同乡列表</span>	

> *接口说明*

>> 登录后获取同乡列表

> *请求说明*
>> http请求方式： get  
 
>> 		http://api.xiangyin.im/v1/user/get_townee_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>HomeProvinceID</td>
      <td>是</td>
      <td>家乡省</td>   
   </tr>
	<tr>
      <td>HomeCityID</td>
      <td>否</td>
      <td>家乡市</td>   
   </tr>
	<tr>
      <td>LivingProvinceID</td>
      <td>否</td>
      <td>居住地省</td>   
   </tr>
	<tr>
      <td>LivingCityID</td>
      <td>否</td>
      <td>居住地市</td>   
   </tr>
 <tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": {
			    "List": [
			      {
			        "UID": 879306419,
			        "NickName": "煎饼馃子",
			        "Avatar": "http://api.xiangyin.im/v1/file/view_file?FileName=0_0_0_0_0_2.png",
			        "Thumbnail": "http://api.xiangyin.im/v1/file/view_file?FileName=0_0_0_0_0_2.png",
			        "HomeProvinceID": 100003,
			        "HomeCityID": 200073,
			        "HomeDistrictID": 300719,
			        "LivingProvinceID": 100003,
			        "LivingCityID": 200073,
			        "LivingDistrictID": 300719,
			        "HomeVoice": "",
			        "VoiceLen": 0,
			        "ProfessionID": 0,
			        "JobID": 0,
			        "IsFollow": 0,
			        "Birthday": 0,
			        "Gender": 0,
			        "IsMember": 0,
			        "IsGuess": 0,
			        "AuthRecvNum": 0,
			        "TagsID": null
			      }
			    ],
			    "MaxID": "2",
			    "MinID": "1"
			  }
			}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>
      <td> </td>   
   </tr>
</table>


 
 
***
动态
***


+ <span id ="5.1">添加动态</span>	

> *接口说明*

>> 登录后添加动态

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/dynamic/add_dynamic 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>HomeProvinceID</td>
      <td>是</td>
      <td>家乡省</td>   
   </tr>
	<tr>
      <td>HomeCityID</td>
      <td>否</td>
      <td>家乡市</td>   
   </tr>
	<tr>
      <td>LivingProvinceID</td>
      <td>否</td>
      <td>居住地省</td>   
   </tr>
	<tr>
      <td>LivingCityID</td>
      <td>否</td>
      <td>居住地市</td>   
   </tr>
 <tr>
      <td>DynamicContent</td>
      <td>是</td>
      <td>动态内容</td>   
   </tr>
	<tr>
      <td>Images</td>
      <td>否</td>
      <td>动态图片</td>   
   </tr>
	<tr>
      <td>Voices</td>
      <td>否</td>
      <td>动态声音</td>   
   </tr>
	<tr>
      <td>VoiceLen</td>
      <td>否</td>
      <td>动态声音长度</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
		{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": {
		    "DynamicID": "14357185749368178751000014228047827"
		  }
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>DynamicID</td>
      <td>动态编号</td>   
   </tr>
</table>


+ <span id ="5.2">获取同省动态列表</span>	

> *接口说明*

>> 登录后获取同省动态列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/dynamic/get_dynamic_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>HomeProvinceID</td>
      <td>是</td>
      <td>家乡省</td>   
   </tr>
	<tr>
      <td>HomeCityID</td>
      <td>否</td>
      <td>家乡市</td>   
   </tr>
	<tr>
      <td>LivingProvinceID</td>
      <td>否</td>
      <td>居住地省</td>   
   </tr>
	<tr>
      <td>LivingCityID</td>
      <td>否</td>
      <td>居住地市</td>   
   </tr>
  <tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": {
			    "List": [
			      {
			        "CommentNum": "40",
			        "DynamicContent": "SDFASDG",
			        "DynamicID": "14355670062661187471000014228047827",
			        "ForwardNum": "2",
			        "GoodNUM": "7",
			        "Images": "",
			        "IsClickGood": 1,
					"LastComment": "sssssddg",
			        "PostTime": "1435567006",
			        "PostUser": {
			          "UID": 4228047827,
			          "NickName": "驼掌",
			          "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_p02.png",
			          "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_p02.png",
			          "HomeProvinceID": 1,
			          "HomeCityID": 2,
			          "HomeDistrictID": 3,
			          "LivingProvinceID": 4,
			          "LivingCityID": 5,
			          "LivingDistrictID": 6,
			          "HomeVoice": "",
			          "VoiceLen": 0,
			          "ProfessionID": 0,
			          "JobID": 0,
			          "IsFollow": 0,
			          "Birthday": 0,
			          "Gender": 0,
			          "IsMember": 0,
			          "IsGuess": 0,
			          "AuthRecvNum": 0,
			          "TagsID": null
			        },
			        "ViewNUM": "17",
			        "VoiceLen": "11",
			        "Voices": ""
			      }
			    ],
			    "MaxID": "20150628",
			    "MinID": "20150619"
			  }
			}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>CommentNum</td>
      <td>评论数</td>   
   </tr>
	<tr>
      <td>DynamicContent</td>
      <td>动态内容</td>   
   </tr>
<tr>
      <td>DynamicID</td>
      <td>动态编号</td>   
   </tr>
<tr>
      <td>ForwardNum</td>
      <td>转发数</td>   
   </tr>
<tr>
      <td>Images</td>
      <td>动态图片</td>   
   </tr>
<tr>
      <td>IsClickGood</td>
      <td>是否点赞（ 0，没有 1，有</td>   
   </tr>
<tr>
      <td>LastComment</td>
      <td>动态最后一条评论</td>   
   </tr>
<tr>
      <td>PostTime</td>
      <td>提交时间戳</td>   
   </tr>
<tr>
      <td>PostUser</td>
      <td>提交人信息</td>   
   </tr>
<tr>
      <td>ViewNUM</td>
      <td>浏览数</td>   
   </tr>
<tr>
      <td>VoiceLen</td>
      <td>动态声音长度</td>   
   </tr>
<tr>
      <td>Voices</td>
      <td>动态声音</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>返回内容最大id</td>   
   </tr>
	<tr>
      <td>MinID</td>
      <td>返回内容最小id</td>   
   </tr>
</table>



+ <span id ="5.3">点赞动态</span>	

> *接口说明*

>> 登录后点赞动态

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/dynamic/click_good 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>DynamicID</td>
      <td>是</td>
      <td>动态id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": null
			}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>

+ <span id ="5.4">转发动态统计</span>	

> *接口说明*

>> 登录后转发动态统计

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/dynamic/click_forward 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>DynamicID</td>
      <td>是</td>
      <td>动态id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": null
			}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>


+ <span id ="5.5">评论动态</span>	

> *接口说明*

>> 登录后评论动态

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/dynamic/add_comment 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>DynamicID</td>
      <td>是</td>
      <td>动态id</td>   
   </tr>
	<tr>
      <td>Contents</td>
      <td>是</td>
      <td>评论内容</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": null
			}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>


+ <span id ="5.6">获取动态评论列表</span>	

> *接口说明*

>> 登录后获取动态评论列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/dynamic/get_comment_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>DynamicID</td>
      <td>是</td>
      <td>动态id</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": {
		    "List": [
		      {
				"CommentID": "14355670062661187471000014228047827_31",
		        "Contents": "sdfsg55",
		        "PostTime": "1435641993",
		        "PostUser": {
		          "UID": 4228047827,
		          "NickName": "驼掌",
		          "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_p02.png",
		          "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_p02.png",
		          "HomeProvinceID": 1,
		          "HomeCityID": 2,
		          "HomeDistrictID": 3,
		          "LivingProvinceID": 4,
		          "LivingCityID": 5,
		          "LivingDistrictID": 6,
		          "HomeVoice": "",
		          "VoiceLen": 0,
		          "ProfessionID": 0,
		          "JobID": 0,
		          "IsFollow": 0,
		          "Birthday": 0,
		          "Gender": 0,
		          "IsMember": 0,
		          "IsGuess": 0,
		          "AuthRecvNum": 0,
		          "TagsID": null
		        }
		      }
		    ],
		    "MaxID": "40",
		    "MinID": "31"
		  }
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Contents</td>
      <td>评论内容</td>   
   </tr>
 <tr>
      <td>CommentID</td>
      <td>评论编号</td>   
   </tr>
 <tr>
      <td>PostTime</td>
      <td>评论时间</td>   
   </tr>
 <tr>
      <td>PostUser</td>
      <td>评论人信息</td>   
   </tr>
 <tr>
      <td>MaxID</td>
      <td>返回内容最大id</td>   
   </tr>
 <tr>
      <td>MinID</td>
      <td>返回内容最小id</td>   
   </tr>
</table>


+ <span id ="5.7">关注用户</span>	

> *接口说明*

>> 登录后关注用户

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/user/click_follow 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>FollowUID</td>
      <td>是</td>
      <td>关注的用户id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>


+ <span id ="5.8">获取用户动态列表</span>	

> *接口说明*

>> 登录后获取用户动态列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/dynamic/get_user_dynamic 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>ViewUID</td>
      <td>是</td>
      <td>查看该用户id的动态</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": {
			    "List": [
			      {
			        "CommentNum": "0",
			        "DynamicContent": "SDFASDG1",
			        "DynamicID": "14355669871264081271000014228047827",
			        "ForwardNum": "0",
			        "GoodNUM": "0",
			        "Images": "",
			        "IsClickGood": 0,
			        "PostTime": "1435566987",
			        "PostUser": {
			          "UID": 4228047827,
			          "NickName": "驼掌",
			          "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_p02.png",
			          "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_p02.png",
			          "HomeProvinceID": 1,
			          "HomeCityID": 2,
			          "HomeDistrictID": 3,
			          "LivingProvinceID": 4,
			          "LivingCityID": 5,
			          "LivingDistrictID": 6,
			          "HomeVoice": "",
			          "VoiceLen": 0,
			          "ProfessionID": 0,
			          "JobID": 0,
			          "IsFollow": 0,
			          "Birthday": 0,
			          "Gender": 0,
			          "IsMember": 0,
			          "IsGuess": 0,
			          "AuthRecvNum": 0,
			          "TagsID": null
			        },
			        "ViewNUM": "0",
			        "VoiceLen": "11",
			        "Voices": ""
			      }
			    ],
			    "MaxID": "2015061000017",
			    "MinID": "2015061000011"
			  }
			}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>返回内容</td>
      <td><a href="#5.6">[同获取动态列表]</td>   
   </tr>
</table>


+ <span id ="5.9">乡音认证用户</span>	

> *接口说明*

>> 登录后关注用户

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/user/auth_voice 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>AuthUID</td>
      <td>是</td>
      <td>认证该用户id的乡音</td>   
   </tr>
	<tr>
      <td>Answer</td>
      <td>是</td>
      <td>认证答案（0，错；1,对</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>

+ <span id ="5.10">删除动态</span>	

> *接口说明*

>> 登录后删除自己发布的动态

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/dynamic/del_dynamic 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>DynamicID</td>
      <td>是</td>
      <td>动态id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>


+ <span id ="5.11">删除动态评论</span>	

> *接口说明*

>> 登录后删除自己发布的动态评论

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/dynamic/del_comment 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>CommentID</td>
      <td>是</td>
      <td>评论id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>



***
个人中心
***

+ <span id ="6.1">修改个人信息</span>	

> *接口说明*

>> 登录后修改个人信息

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/user/update_user_info 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>HomeProvinceID</td>
      <td>否</td>
      <td>家乡省</td>   
   </tr>
	<tr>
      <td>HomeCityID</td>
      <td>否</td>
      <td>家乡是</td>   
   </tr>
	<tr>
      <td>HomeDistrictID</td>
      <td>否</td>
      <td>家乡区县</td>   
   </tr>
	<tr>
      <td>LivingProvinceID</td>
      <td>否</td>
      <td>现居地省</td>   
   </tr>
	<tr>
      <td>LivingCityID</td>
      <td>否</td>
      <td>现居地市</td>   
   </tr>
	<tr>
      <td>LivingDistrictID</td>
      <td>否</td>
      <td>现居地区县</td>   
   </tr>
	<tr>
      <td>NickName</td>
      <td>否</td>
      <td>昵称</td>   
   </tr>
	<tr>
      <td>ProfessionID</td>
      <td>否</td>
      <td>职业所属行业</td>   
   </tr>
	<tr>
      <td>JobID</td>
      <td>否</td>
      <td>职业编号</td>   
   </tr>
	<tr>
      <td>Gender</td>
      <td>否</td>
      <td>性别</td>   
   </tr>
	<tr>
      <td>Birthday</td>
      <td>否</td>
      <td>生日时间戳</td>   
   </tr>
	<tr>
      <td>TagID</td>
      <td>否</td>
      <td>标签编号，多个用逗号分隔</td>   
   </tr>
	<tr>
      <td>DiySign</td>
      <td>否</td>
      <td>个性签名</td>   
   </tr>
	<tr>
      <td>Avatar</td>
      <td>否</td>
      <td>上传头像图片</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
		  "Code": 1,
		  "Desc": "成功",
		  "Info": 
		}			
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>Code</td>
      <td>返回码</td>   
   </tr>
</table>



+ <span id ="6.2">我关注的用户</span>	

> *接口说明*

>> 登录后获取我关注的用户列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/follow_other_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": {
			    "List": [
			      {
			        "UID": 3501886692,
			        "NickName": "酸菜",
			        "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_3.png",
			        "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_3.png",
			        "HomeProvinceID": 100001,
			        "HomeCityID": 200001,
			        "HomeDistrictID": 300001,
			        "LivingProvinceID": 100001,
			        "LivingCityID": 200001,
			        "LivingDistrictID": 300005,
			        "HomeVoice": "",
			        "VoiceLen": 0,
			        "ProfessionID": 0,
			        "JobID": 0,
			        "IsFollow": 0,
			        "Birthday": 0,
			        "Gender": 0,
			        "IsMember": 0,
			        "IsGuess": 0,
			        "AuthRecvNum": 0,
			        "TagsID": null,
			        "LastLoginTime": 0
			      }
			    ],
			    "MaxID": "0",
			    "MinID": "9"
			  }
			}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>返回内容</td>
      <td><a href="#5.6">[同获取动态列表(PostUser)]</td>   
   </tr>
</table>

+ <span id ="6.3">关注我的用户</span>	

> *接口说明*

>> 登录后获取关注我的用户列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/follow_me_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": {
			    "List": [
			      {
			        "UID": 3501886692,
			        "NickName": "酸菜",
			        "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_3.png",
			        "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_3.png",
			        "HomeProvinceID": 100001,
			        "HomeCityID": 200001,
			        "HomeDistrictID": 300001,
			        "LivingProvinceID": 100001,
			        "LivingCityID": 200001,
			        "LivingDistrictID": 300005,
			        "HomeVoice": "",
			        "VoiceLen": 0,
			        "ProfessionID": 0,
			        "JobID": 0,
			        "IsFollow": 0,
			        "Birthday": 0,
			        "Gender": 0,
			        "IsMember": 0,
			        "IsGuess": 0,
			        "AuthRecvNum": 0,
			        "TagsID": null,
			        "LastLoginTime": 0
			      }
			    ],
			    "MaxID": "0",
			    "MinID": "9"
			  }
			}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>返回内容</td>
      <td><a href="#5.6">[同获取动态列表(PostUser)]</td>   
   </tr>
</table>


+ <span id ="6.4">我参加的活动</span>	

> *接口说明*

>> 登录后获取我参加的活动列表或他人的活动列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/activity/user_talk_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
  <tr>
      <td>ViewUID</td>
      <td>否</td>
      <td>查看他人账号的uid</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": {
				    "List": [
				      {
						"ActivityID": 10000001,
        				"ActivityName": "各地最难懂方言大PK",
				        "BadNUM": "3",
				        "CommentNUM": "34",
				        "GoodNUM": "3",
				        "Images": "",
				        "IsClickGoodOrBad": 0,
				        "PostTime": "1435287360",
				        "PostUser": {
				          "UID": 4228047827,
				          "NickName": "kill",
				          "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_1.png",
				          "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_1.png",
				          "HomeProvinceID": 1020988,
				          "HomeCityID": 200001,
				          "HomeDistrictID": 300001,
				          "LivingProvinceID": 100001,
				          "LivingCityID": 200001,
				          "LivingDistrictID": 300005,
				          "HomeVoice": "",
				          "VoiceLen": 0,
				          "ProfessionID": 0,
				          "JobID": 0,
				          "IsFollow": 0,
				          "Birthday": 0,
				          "Gender": 0,
				          "IsMember": 0,
				          "IsGuess": 1,
				          "AuthRecvNum": 12,
				          "TagsID": [
				            "666"
				          ],
				          "LastLoginTime": 1435909644
				        },
				        "TalkContent": "ssdfsdgaa",
				        "TalkID": "1434908579916858202100000014228047827",
				        "VoiceLen": "1",
				        "Voices": ""
				      }
				    ],
				    "MaxID": "201506100000014",
				    "MinID": "201506100000011"
				  }
			}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>返回内容</td>
      <td><a href="#3.2">[获取活动的所有参与内容列表]</td>   
   </tr>
</table>


+ <span id ="6.5">我的认证记录</span>	

> *接口说明*

>> 登录后获取我的认证记录列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/auth_voice_log 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": {
				    "List": [
				      {
				        "AuthTime": "1436151577",
				        "AuthType": "0",
				        "PostUser": {
				          "UID": 3435985337,
				          "NickName": "狮子头",
				          "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_5.png",
				          "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_5.png",
				          "HomeProvinceID": 100001,
				          "HomeCityID": 200001,
				          "HomeDistrictID": 300001,
				          "LivingProvinceID": 100001,
				          "LivingCityID": 200001,
				          "LivingDistrictID": 300005,
				          "HomeVoice": "",
				          "VoiceLen": 0,
				          "ProfessionID": 0,
				          "JobID": 0,
				          "IsFollow": 0,
				          "Birthday": 0,
				          "Gender": 0,
				          "IsMember": 0,
				          "IsGuess": 1,
				          "AuthRecvNum": 1,
				          "TagsID": null,
				          "LastLoginTime": 0
				        }
				      }
				    ],
				    "MaxID": "1",
				    "MinID": "1"
				  }
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>AuthTime</td>
      <td>认证时间</td>   
   </tr>
	<tr>
      <td>AuthType</td>
      <td>认证类型（0，认证别人；1，被认证</td>   
   </tr>
	<tr>
      <td>PostUser</td>
      <td>认证或被认证的人</td>   
   </tr>
</table>


+ <span id ="6.6">消息提醒类型列表</span>	

> *接口说明*

>> 登录后获取我的消息提醒类型列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/get_remind_msg_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": [
			    {
			      "LastMsg": "sdfsdyyyfd,请求你认证乡音。",
			      "LastTime": "1436243310",
			      "MsgTypeID": "1",
			      "MsgTypeName": "求乡音认证申请",
			      "UnreadNum": "1"
			    }
			  ]
			}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>LastMsg</td>
      <td>最后一条消息内容</td>   
   </tr>
	<tr>
      <td>LastTime</td>
      <td>最后一条消息时间</td>   
   </tr>
	<tr>
      <td>MsgTypeID</td>
      <td>消息类型编号,取值说明（1=求乡音认证申请；2=乡音被认证提醒；3=乡音团队）</td>   
   </tr>
	<tr>
      <td>MsgTypeName</td>
      <td>消息类型名称</td>   
   </tr>
	<tr>
      <td>UnreadNum</td>
      <td>未读消息数</td>   
   </tr>
</table>


+ <span id ="6.7">求认证</span>	

> *接口说明*

>> 登录后求乡音认证,随机发给3个用户

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/request_auth_voice 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": 
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
</table>



+ <span id ="6.8">获取求认证列表</span>	

> *接口说明*

>> 登录后获取请求我认证的列表

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/request_auth_voice_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": {
				    "List": [
				      {
				        "UID": 3962328886,
				        "NickName": "马奶酒",
				        "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_5.png",
				        "Thumbnail": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_5.png",
				        "HomeProvinceID": 100001,
				        "HomeCityID": 200001,
				        "HomeDistrictID": 300001,
				        "LivingProvinceID": 100001,
				        "LivingCityID": 200001,
				        "LivingDistrictID": 300005,
				        "HomeVoice": "",
				        "VoiceLen": 0,
				        "ProfessionID": 0,
				        "JobID": 0,
				        "IsFollow": 0,
				        "Birthday": 0,
				        "Gender": 0,
				        "IsMember": 0,
				        "IsGuess": 0,
				        "AuthRecvNum": 0,
				        "TagsID": null,
				        "LastLoginTime": 0
				      }
				    ],
				    "MaxID": "1",
				    "MinID": "1"
				  }
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
</table>


+ <span id ="6.9">提醒设置</span>	

> *接口说明*

>> 登录后获取提醒设置

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/user/set_remind 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>Comment</td>
      <td>是</td>
      <td>评论提醒（0没有，1有）</td>   
   </tr>
	<tr>
      <td>Follow</td>
      <td>是</td>
      <td>关注提醒（0没有，1有）</td>   
   </tr>
	<tr>
      <td>Activity</td>
      <td>是</td>
      <td>活动提醒（0没有，1有）</td>   
   </tr>
	<tr>
      <td>Message</td>
      <td>是</td>
      <td>留言或聊天提醒（0没有，1有）</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": 
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
</table>


+ <span id ="6.10">反馈</span>	

> *接口说明*

>> 登录后添加反馈

> *请求说明*
>> http请求方式： post  
 
>> 		http://api.xiangyin.im/v1/user/add_feedback 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>Contents</td>
      <td>是</td>
      <td>反馈内容</td>   
   </tr>
	<tr>
      <td>Contact</td>
      <td>否</td>
      <td>联系方式</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": 
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
</table>


+ <span id ="6.11">乡音范例</span>	

> *接口说明*

>> 录制乡音文字范例

> *请求说明*
>> http请求方式： get  
 
>> 		http://api.xiangyin.im/v1/common/home_voice_example 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>无参数</td>
      <td></td>
      <td></td>   
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
						"Code": 1,
						"Desc": "成功",
						"Info": [
								{
								"ProvinceID": "100001",
								"Example": [
										"真正的勇士，敢于肥还贪吃，困还熬夜，穷还追星，丑还颜控。",
										"曾经，有一份真挚的爱情摆在我面前，我没有珍惜，等到我失去的",
										"你别躲在里面不出声，我知道你在家。你有本事抢男人，怎么没本事开门啊！开门啊开门啊开门开门开门啊。"
										]
								}
							]
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
	<tr>
       <th><I>ProvinceID</I></th>
       <th><I>省编码</I></th>
   </tr>
	<tr>
       <th><I>Example</I></th>
       <th><I>文字范例数组</I></th>
   </tr>
</table>

+ <span id ="6.12">设置已读提醒类型消息数</span>	

> *接口说明*

>> 登录后设置已读提醒类型消息数

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/set_remind_read_num 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
 <tr>
      <td>MsgTypeID</td>
      <td>是</td>
      <td>消息类型编号<a href="#6.6">[同消息提醒类型列表返回字段]</a></td>   
   </tr>
 <tr>
      <td>ReadNum</td>
      <td>是</td>
      <td>已读条数</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

+ <span id ="6.13">获取提醒设置</span>	

> *接口说明*

>> 登录后获取提醒设置

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/get_remind_set 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": {
				    "Activity": "1",
				    "Comment": "1",
				    "Follow": "0",
				    "Message": "0"
				  }
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
	<tr>
      <td>Activity</td>
      <td>活动提醒（0没有，1有）</td>    
   </tr>
	<tr>
      <td>Comment</td>
      <td>评论提醒（0没有，1有）</td>   
   </tr>
	<tr>
      <td>Follow</td>
      <td>关注提醒（0没有，1有）</td> 
   </tr>
	<tr>
      <td>Message</td>
      <td>留言或聊天提醒（0没有，1有）</td> 
   </tr>
</table>

+ <span id ="6.14">获取系统消息列表（乡音团队)</span>	

> *接口说明*

>> 登录后获取系统消息列表（乡音团队)

> *请求说明*
>> http请求方式： Get  
 
>> 		http://api.xiangyin.im/v1/user/get_sys_msg_list 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>PageType</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
	<tr>
      <td>MaxID</td>
      <td>是</td>
      <td><a href="#3.2">[同获取活动的所有参与内容列表]</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
				{
				  "Code": 1,
				  "Desc": "成功",
				  "Info": {
				    "List": [
				      {
				        "ID": "2",
				        "IsRead": "1",
				        "Messages": "乡音团队系统消息测试22。。。",
				        "PostTime": "1436806247"
				      },
				      {
				        "ID": "1",
				        "IsRead": "1",
				        "Messages": "乡音团队系统消息测试。。。",
				        "PostTime": "1436506247"
				      }
				    ],
				    "MaxID": "2",
				    "MinID": "1"
				  }
				}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
	<tr>
      <td>ID</td>
      <td>流水号</td>    
   </tr>
	<tr>
      <td>IsRead</td>
      <td>是否已读（0没有，1有）</td>   
   </tr>
	<tr>
      <td>Messages</td>
      <td>消息内容</td> 
   </tr>
	<tr>
      <td>PostTime</td>
      <td>消息发送时间</td> 
   </tr>
 <tr>
      <td>MaxID</td>
      <td>返回内容最大id</td>   
   </tr>
 <tr>
      <td>MinID</td>
      <td>返回内容最小id</td>   
   </tr>
</table>

+ <span id ="6.15">获取昵称信息</span>	

> *接口说明*

>> 登录后获取昵称信息

> *请求说明*
>> http请求方式： Post  
 
>> 		http://api.xiangyin.im/v1/user/get_nickname 


> *请求参数说明*
>>  
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>是否必须</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>UID</td>
      <td>是</td>
      <td>手机账号的id</td>   
   </tr>
	<tr>
      <td>ViewUID</td>
      <td>是</td>
      <td>需要获取的uid昵称信息</td>   
   </tr>
   <tr>
      <td>Sign</td>
      <td>是</td>
      <td><a href="#1.3">[详见]</a></td>  
   </tr>
</table>

 
> *返回说明*
>>  
>>		返回内容：		 
			{
			  "Code": 1,
			  "Desc": "成功",
			  "Info": {
			    "Avatar": "http://127.0.0.1/v1/file/view_file?FileName=0_0_0_0_0_5.png",
			    "NickName": "sdfsdyyyfd"
			  }
			}		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>参数</I></th>
       <th><I>说明</I></th>
   </tr>
	<tr>
      <td>Avatar</td>
      <td>头像</td>    
   </tr>
	<tr>
      <td>NickName</td>
      <td>昵称</td>   
   </tr>
</table>


## <span id="retCode">返回码说明</span>
 
		1 - 99999  公共返回码
		100000 - 199999  账号服务接口相关返回码
		200000 - 299999  群组服务接口相关返回码
		300000 - 399999  聊天服务接口相关返回码
		
<table class="table table-striped table-condensed" style="font-size:12px;">
    <tr>
       <th><I>返回码</I></th>
       <th><I>说明</I></th>
   </tr>
   <tr>
      <td>1</td>
      <td>成功</td>   
   </tr>
	<tr>
      <td>2</td>
      <td>失败</td>  
   </tr>
	<tr>
      <td>3</td>
      <td>sign错误</td>  
   </tr>
	<tr>
      <td>4</td>
      <td>预留</td>  
   </tr>
	<tr>
      <td>5</td>
      <td>cellphone参数为空</td>  
   </tr>
	<tr>
      <td>6</td>
      <td>请求方法错误</td>  
   </tr>
	<tr>
      <td>7</td>
      <td>地区列表为空</td>  
   </tr>
	<tr>
      <td>8</td>
      <td>登录令牌过期或未登录</td>  
   </tr>
	<tr>
      <td>100000</td>
      <td>短信验证码错误</td>  
   </tr>
	<tr>
      <td>100001</td>
      <td>短信验证服务器异常</td>  
   </tr>
	<tr>
      <td>100002</td>
      <td>注册账号已存在</td>  
   </tr>
	<tr>
      <td>100003</td>
      <td>账号注册失败</td>  
   </tr>
	<tr>
      <td>100004</td>
      <td>登录失败.账号或密码错误</td>  
   </tr>
	<tr>
      <td>100005</td>
      <td>预留</td>  
   </tr>
	<tr>
      <td>100006</td>
      <td>建立家乡档案vdata参数内容为空</td>  
   </tr>
	<tr>
      <td>100007</td>
      <td>建立家乡档案vdata参数的内容大于5m(预留)</td>  
   </tr>
	<tr>
      <td>100008</td>
      <td>建立家乡档案上传文件失败</td>  
   </tr>
	<tr>
      <td>100009</td>
      <td>建立家乡档案失败</td>  
   </tr>
	<tr>
      <td>200000</td>
      <td>加入群组失败</td>  
   </tr>
<tr>
      <td>200001</td>
      <td>群组不存在</td>  
   </tr>
</table>