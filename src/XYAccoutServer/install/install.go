package install

import (
	
	"strings"
	"fmt"
	//"database/sql"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"os"
)

func InstallDB(dns string){
	

	f := fmt.Sprintf("use %s ;",dns)
	
	for i := 0 ;i < 8 ; i++ {	
	str :=  `
		
		DROP TABLE IF EXISTS  user_base_[numeber] ;
        CREATE TABLE  user_base_[numeber]  (
           ID  int(11) unsigned NOT NULL AUTO_INCREMENT, 
           UID  int(11) unsigned NOT NULL COMMENT '用户id',
           Account  varchar(45) NOT NULL COMMENT '用户账号',
           PassWord  varchar(45) NOT NULL COMMENT '登录密码',
		   BindPhone  bigint DEFAULT NULL COMMENT '绑定手机号',
           RegType  tinyint(1) DEFAULT '1' COMMENT '注册类型,1:手机登陆 2：微博登陆 3:微信登陆',
           IsMember  tinyint(1) DEFAULT '0' COMMENT '是否会员',
           AuthSendNum  int(11) unsigned DEFAULT '0' COMMENT '认证他人乡音次数',
           AuthRecvNum  int(11) unsigned  DEFAULT '0' COMMENT '获得乡音认证次数',
           RemainDays  int(11) unsigned DEFAULT '0' COMMENT 'vip剩余天数',
           GrowNum  int(11) unsigned DEFAULT '0' COMMENT '成长值',
           RPNum  int(11) unsigned DEFAULT '0' COMMENT '人品值',
           MedalNum  int(11) unsigned  DEFAULT '0' COMMENT '勋章个数',
		   RegisterTime  int(11)  DEFAULT '0' COMMENT '注册时间',
		   HomeProvinceID  int(11) DEFAULT NULL COMMENT '家乡省id',
          PRIMARY KEY ( ID ),
          UNIQUE KEY  UID_UNIQUE  ( UID  ),
		  UNIQUE KEY  Account_UNIQUE  ( Account ),
          KEY  User_base_index  ( UID , Account  , IsMember , RPNum , MedalNum )
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户基本信息表';
        
        
        DROP TABLE IF EXISTS  user_detail_info_[numeber] ;
        CREATE TABLE  user_detail_info_[numeber]  (
           ID  int(11) unsigned NOT NULL AUTO_INCREMENT, 
           UID  int(11) unsigned NOT NULL COMMENT '用户id',
           NickName  varchar(45) DEFAULT NULL COMMENT '昵称',
           Avatar  varchar(100) DEFAULT NULL COMMENT '头像',
           Thumbnail  varchar(100) DEFAULT NULL COMMENT '头像缩略图',
           TagID  varchar(100) DEFAULT NULL COMMENT '标签id,多个用逗号分隔',
           DiySign  varchar(45) DEFAULT NULL COMMENT '个性签名',
           Gender  int(11) DEFAULT NULL COMMENT '性别',
           Birthday  int(11) DEFAULT NULL COMMENT '生日',
		   JobID  int(11) DEFAULT NULL COMMENT '职业id',
		   ProfessionID  int(11) DEFAULT NULL COMMENT '职业所属行业id',
		   HomeProvinceID  int(11) DEFAULT NULL COMMENT '家乡省id',
           HomeCityID  int(11) DEFAULT NULL COMMENT '家乡城市id',
           HomeDistrictID  int(11) DEFAULT NULL COMMENT '家乡区县id',
           LivingProvinceID  int(11) DEFAULT NULL COMMENT '居住地省id',
           LivingCityID  int(11) DEFAULT NULL COMMENT '居住地城市id',
           LivingDistrictID  int(11) DEFAULT NULL COMMENT '居住地区县id',
		   PushType  tinyint(1) DEFAULT '1' COMMENT '推送类型,1:ios 2：android ',
           PushID  varchar(100) DEFAULT NULL,
           HomeVoice  varchar(100) DEFAULT NULL,
           VoiceLen  int(11) DEFAULT NULL,
           NorthLatitude  int(11) DEFAULT NULL COMMENT '北纬',
           EastLongtude  int(11) DEFAULT NULL COMMENT '东经',
          PRIMARY KEY ( ID ),
          UNIQUE KEY  UID_UNIQUE  ( UID),
          KEY  DETAIL_INDEX  ( NickName , UID , Birthday , Gender , TagID,HomeProvinceID , HomeCityID , LivingProvinceID , LivingCityID , LivingDistrictID ,JobID,ProfessionID,PushType,PushID)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户详细信息表(按家乡省 hash)';
                

		
		DROP TABLE IF EXISTS  voice_auth_log_[numeber] ;
		CREATE TABLE   voice_auth_log_[numeber]  (
		   ID  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		   UID  INT(11) UNSIGNED NULL ,
		   AuthType  TINYINT(1) DEFAULT 0 ,
		   AuthUID  INT(11) unsigned  NULL ,
		   AuthTime  INT(11) unsigned  NULL ,
		  PRIMARY KEY ( ID ),
		  UNIQUE INDEX  auth_msg_UNIQUE  ( UID  ASC,  AuthType  ASC,  AuthUID  ASC),
		  INDEX  auth_msg_INDEX  ( UID  ASC,  AuthType  ASC,  AuthTime  ASC))
		ENGINE = InnoDB
		DEFAULT CHARACTER SET = utf8
		COMMENT = '乡音认证成功日志';
		
		
		DROP TABLE IF EXISTS  user_remind_set_[numeber] ;
		CREATE TABLE  user_remind_set_[numeber]  (
		   ID  BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
		   UID  INT(11) unsigned  NULL,
		   Comment  TINYINT(1) NULL DEFAULT 0  COMMENT '评论提醒开关',
		   Follow  TINYINT(1) NULL DEFAULT 0  COMMENT '关注提醒开关',
		   Activity  TINYINT(1) NULL DEFAULT 0 COMMENT '活动提醒开关',
		   Message  TINYINT(1) NULL DEFAULT 0  COMMENT '留言或聊天提醒开关',
		  PRIMARY KEY ( ID ),
		  UNIQUE INDEX   REMIND_SET_UNIQUE  ( UID  ASC),
		  INDEX  REMIND_SET_INDEX  ( UID  ASC))
		ENGINE = InnoDB
		DEFAULT CHARACTER SET = utf8
		COMMENT = '用户提醒设置表';
		
		
		DROP TABLE IF EXISTS  voice_auth_request_[numeber] ;
		CREATE TABLE   voice_auth_request_[numeber]  (
		   ID  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		   UID  INT(11) UNSIGNED NULL COMMENT '用户uid',
		   ReuestUID  INT(11) unsigned  NULL COMMENT '求认证uid',
		  PRIMARY KEY ( ID ),
		  UNIQUE INDEX  auth_request_UNIQUE  ( UID  ASC,  ReuestUID  ASC),
		  INDEX  auth_requesnt_INDEX  ( UID  ASC,  ReuestUID  ASC))
		ENGINE = InnoDB
		DEFAULT CHARACTER SET = utf8
		COMMENT = '求乡音认证';
		
		DROP TABLE IF EXISTS  user_remind_msg_[numeber] ;
		CREATE TABLE  user_remind_msg_[numeber]  (
		   ID  BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
		   UID  INT(11) unsigned  NULL,
		   MsgTypeID  TINYINT(1) NULL DEFAULT 0 COMMENT '消息类型',
		   UnreadNum  INT(11) unsigned  NULL DEFAULT 0  COMMENT '未读消息数',
		   LastMsg  varchar(150) DEFAULT NULL COMMENT '最后一条消息',
		   LastTime  INT(11) unsigned  NULL DEFAULT 0 COMMENT '最后一天消息时间',
		  PRIMARY KEY ( ID ),
		  UNIQUE INDEX   REMIND_MSG_UNIQUE  ( UID  ASC,MsgTypeID  ASC),
		  INDEX  REMIND_MSG_INDEX  ( UID  ASC,MsgTypeID  ASC))
		ENGINE = InnoDB
		DEFAULT CHARACTER SET = utf8
		COMMENT = '用户提醒消息';
		

    
 `     
	n := fmt.Sprintf("%d",i)
	f += strings.Replace(str,"[numeber]",n,-1)
	println("---")
	}
	
	
	
	e := ioutil.WriteFile("./user_db.sql",[]byte(f),777)
	if e != nil {
		println(e.Error())
		return
	}
	path,_ := os.Getwd() 
	println("111",path)
	file,_ := filepath.Abs(path)
	println("2222:",file)
	println(filepath.Base(path))
	cmd := fmt.Sprintf(" mysql -u root -p123456  < %s","/Users/david/project_code/go_project/src/XYAPIServer/XYAccoutServer/user_db.sql")
	println(cmd)
	c := exec.Command("/bin/sh",cmd)
	err := c.Start()
	
	if err != nil {
			println("exec err:",err.Error())
	}
	//println(sqlStr)
}

func installGroupDB(){
	
}

