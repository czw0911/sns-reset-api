//@Description api统一输出结构
//@Contact czw@outlook.com

package XYLibs


var RespStateCode = map[string]XYAPIResponse{

		//公共类1 - 99999
		//成功
		"nil" : {0,"未知错误",nil},
		"ok" : {1,"成功",nil},
		//失败
		"fail" :{2,"失败",nil},
		//签名不匹配
		"sign_error" : { 3,"签名错误",nil},
		//短信发送次数超过限制
		"sms_send_full" : { 4,"短信验证发送次数太频繁",nil},
		//短信发送号码为空
		"sms_phone_null" : { 5,"短信发送号码为空",nil},
		//请求方法错误
		"method_not_find" : { 6,"请求方法错误",nil},
		//地区列表为空
		"region_null" : { 7,"地区类别为空",nil},
		//登录令牌过期或未登录
		"login_token_expire" : { 8,"登录令牌过期或未登录",nil},
		"view_file_format_fail" : { 9,"文件名格式错误",nil},
		"view_file_name_fail" : { 10,"文件名错误",nil},
		"view_file_not_find" : { 11,"文件找不到",nil},
		"panging_list_is_null" : { 12,"没有数据",nil},
		"comment_content_null" : { 13,"评论内容不能为空",nil},
		"feedback_content_null" : { 14,"反馈内容不能为空",nil},
		"feedback_content_fail" : { 15,"反馈失败",nil},
		"user_detail_error" : { 16,"用户信息异常",nil},
		
		//用户类100000 - 199999
		//短信验证码错误
		"sms_verify_code_error" : {100000,"短信验证码错误",nil},
		//短信验证服务器异常
		"sms_server_error" : { 100001,"短信验证码不存在",nil},
		//注册账号已存在
		"reg_user_isexist" : { 100002,"账号已存在",nil},
		//账号注册失败
		"reg_user_fail" : { 100003,"账号注册失败",nil},
		//登录失败.账号或密码错误
		"login_fail" : { 100004,"账号或密码错误",nil},
		//登录失败次数太多，账号被锁定1小时
		"login_user_lock" : { 100005,"账号被锁定",nil},
		//上传乡音文件为空
		"upload_voice_null" : { 100006,"上传文件为空",nil},
		//上传乡音文件太大
		"upload_voice_max" : { 100007,"上传文件太大",nil},
		//上传文件失败
		"upload_fail" : { 100008,"上传文件失败",nil},
		//建立家乡档案失败
		"hometown_fail" : { 100009,"建立家乡档案失败",nil},
		//找回密码失败
		"recover_fail" : { 100010,"修改密码错误",nil},
		//账号注册失败
		"reg_user_pwd_null" : { 100011,"账号或密码为空",nil},
		//推送id添加失败
		"add_pushid_fail" : { 100012,"添加推送id失败",nil},
		"user_update_homevoice_fail" : { 100013,"修改乡音失败",nil},
		"user_mulauth_homevoice_fail" : { 100014,"不能重复认证乡音",nil},
		"user_save_homevoice_fail" : { 100015,"保持认证乡音结果失败",nil},
		"user_self_homevoice_fail" : { 100016,"不能认证自己",nil},
		"user_uid_homevoice_null" : { 100017,"被认证用户不能为空",nil},
		"user_update_user_info_fail" : { 100018,"修改用户信息失败",nil},
		"user_remind_set_fail" : { 100019,"提醒设置失败",nil},
		"user_remind_read_set_fail" : { 100020,"提醒已读写入失败",nil},
		"user_remind_set_get_fail" : { 100021,"提醒设置获取失败",nil},
		"user_account_not_exist" : { 100022,"账号不存在",nil},
		"user_nickname_is_exist" : { 100023,"昵称已存在",nil},
		
		//加入群组失败
		"join_group_fail" : { 200000,"加入群组失败",nil},
		//群组不存在
		"join_group_not_exist" : { 200001,"群组不存在",nil},
		//退出群组失败
		"exit_group_fail" : { 200002,"退出群组失败",nil},
		"add_group_fail" : { 200003,"添加群组失败",nil},
		"add_group_region_fail" : { 200004,"未知家乡名",nil},
		"add_group_living_fail" : { 200005,"未知现居地或职业名",nil},
		"exchange_visiting_card_fail" : { 200006,"递交名片失败",nil},
		"exchange_visiting_card_recv_fail" : { 200007,"递交名片时对方接收失败",nil},
		"activity_upload_file_max_size_err" : { 200008,"上传文件不能大于4m",nil},
		"activity_upload_file_fail" : { 200009,"上传文件失败",nil},
		"activity_tail_fail" : { 200010,"提交失败",nil},
		"activity_not_find" : { 200011,"活动不存在",nil},
		"activity_click_good_fail" : { 200012,"点赞失败",nil},
		"activity_click_bad_fail" : { 200013,"吐槽失败",nil},
		"activity_list_is_null" : { 200014,"没有数据",nil},
		"activity_delete_fail" : { 200015,"活动删除失败",nil},
		"activity_comment_id_error" : { 200016,"活动评论id错误",nil},
		"activity_comment_delete_fail" : { 200017,"活动评论删除失败",nil},
		
		"dynamic_provinceid_null" : { 300000,"家乡所在省不能为空",nil},
		"dynamic_add_fail" : { 300001,"动态发布失败",nil},
		"dynamic_cityid_null" : { 300002,"居住地省或市不能为空",nil},
		"dynamic_list_is_null" : { 300003,"没有数据",nil},
		"dynamic_click_good_fail" : { 300004,"点赞失败",nil},
		"dynamic_mulclick_good_fail" : { 300005,"不能重复点赞",nil},
		"dynamic_click_forward_fail" : { 300006,"转发失败",nil},
		"dynamic_not_find" : { 300007,"动态不存在",nil},
		"dynamic_comment_fail" : { 300008,"评论失败",nil},
		"dynamic_mulclick_follow" : { 300009,"不能重复关注",nil},
		"dynamic_follow_fail" : { 300010,"关注失败",nil},
		"dynamic_content_null" : { 300011,"动态内容不能为空",nil},
		"dynamic_delete_fail" : { 300012,"动态删除失败",nil},
		"dynamic_comment_id_error" : { 300013,"动态评论id错误",nil},
		"dynamic_comment_delete_fail" : { 300014,"动态评论删除失败",nil},
}

type XYAPIResponse struct {
	//状态码
	Code int
	//状态码描述
	Desc string
	//输出内容
	Info interface{}
	
}



