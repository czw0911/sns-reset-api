//@Description nosql存储key前缀定义
//@Contact czw@outlook.com

package XYLibs


const (
	
		
	//////////
	//
	//账号服务
	//
	/////////
	
	
	//短信息验证码
	NO_SQL_PREFIX_KEY_SMS = "smsSendVcode"
	//找回密码验证码
	NO_SQL_PREFIX_KEY_SMS_RECOVER_PWD = "smssendrecoverpwd"
	//登录token
	NO_SQL_USER_LOGIN_TOKEN = "loginToken"
	//用户头像信息 (mset)
	NO_SQL_USER_AVATER_INFO = "userAvater"
	//用户关注群组
	NO_SQL_USER_GROUP_INFO = "userGroup"
	//用户名片夹 (sadd）
	NO_SQL_USER_VISITING_CARD = "userVisitingCard"
	
	//所有注册注册用户uid(sadd -> userRecordVoiceList -> uid)
	NO_SQL_USER_ALL_REGISTER_UID = "userAllRegisterUID"
	
	//ios用户推送id(MSET -> pushIDiOS:uid -> pushID)
	NO_SQL_USER_IOS_PUSHI_ID = "pushIDiOS"
	
	//是否已推送(SETEX -> pushIsSend:uid -> 21600（6小时)  1 
	NO_SQL_USER_PUSH_IS_SEND = "pushIsSend"
		
	//已经录制了乡音的用户(sadd -> userRecordVoiceList -> uid）
	NO_SQL_USER_ALREADY_RECORD_VOICE_LIST = "userAlreadyRecordVoiceList"
	
	//认证过乡音的所有人列表(sadd -> userVocieAuthList:uid ->  touid)
	NO_SQL_USER_SEND_VOICE_AUTH_LIST = "userSendVocieAuthList"
	
	//被认证乡音次数(INCR -> userRecvVoiceAuthNum:uid)
	NO_SQL_USER_RECV_VOICE_AUTH_NUM = "userRecvVoiceAuthNum"
	
	//认证乡音错误次数(INCR -> userSendVoiceAuthErrorNum:uid)
	NO_SQL_USER_SEND_VOICE_AUTH_ERROR_NUM = "userSendVoiceAuthErrorNum"
	
	//乡音认证排名(zadd -> userVoiceAuthRanking -> 认证他人次数 uid)
	NO_SQL_USER_VOICE_AUTH_RANKING_LIST = "userVoiceAuthRanking"
	
	//我关注的用户(zadd -> userFollowOther:uid -> 时间戳 touid)
	NO_SQL_USER_FOLLOW_YOU ="userFollowOther"
	
	//关注我的用户(zadd -> userFollowMe:uid -> 时间戳 touid)
	NO_SQL_USER_FOLLOW_ME ="userFollowMe"
	
		
	//////////
	//
	//群组服务
	//
	/////////
	
	
	
	//群组的用户
	NO_SQL_GROUP_ALL_USERS = "groupUser"
	//群组活动标签类型（sadd 集合）
	NO_SQL_GROUP_ACTIVITY_ALL_TAGS = "groupActivityTags"
	//群组发布的活动表索引 (zadd 集合)
	NO_SQL_GROUP_PUBLISH_ACTIVITY_TABLE_INDEX = "groupActivityTableNameIndex"
	
	
	//////////
	//
	//活动服务
	//
	/////////
	
	
	//活动主题参与人数(INCR -> activityJoinNum:活动编号)
	NO_SQL_ACTIVITY_JOIN_NUM = "activityJoinNum"
	//活动谈论表索引 (zadd -> activityTalkTableNameIndex:活动编号 -> 年月 表名)
	NO_SQL_ACTIVITY_TALK_TABLE_INDEX = "activityTalkTableNameIndex"
	//用户发表活动谈论表索引(zadd -> activityTalkByUserTableNameIndex:uid -> 年月 表名)
	NO_SQL_ACTIVITY_TALK_BYUSER_TABLE_INDEX = "activityTalkByUserTableNameIndex"
	//我点赞或吐槽的谈论id(zadd -> activityTalkIDClickGoodOrBadList:uid -> [good or bad type]  talkid)
	NO_SQL_ACTIVITY_CLICK_GOOD_OR_BAD_ID = "activityTalkIDClickGoodOrBadList"
	
	
	//////////
	//
	//动态服务
	//
	/////////
	
	
	//动态表索引 (zadd -> dynamicTableNameIndex:家乡省编号 -> 年月 表名)
	NO_SQL_DYNAMIC_TABLE_INDEX = "dynamicTableNameIndex"
	//用户发表动态表索引(zadd -> dynamicByUserTableNameIndex:uid -> 年月 表名)
	NO_SQL_DYNAMIC_BYUSER_TABLE_INDEX = "dynamicByUserTableNameIndex"
	//动态浏览次数(INCR -> dynamicViewNum:动态编号)
	NO_SQL_DYNAMIC_VIEW_NUM = "dynamicViewNum"
	//我点赞过的动态id (sadd -> dynamicClickGoodList:uid -> 动态id)
	NO_SQL_DYNAMIC_CLICK_GOOD_DYNAMIC_ID = "dynamicClickGoodList"
)