//@Description 通用函数
//@Contact czw@outlook.com

package XYLibs

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"	
	"errors"	
	"net"
	"net/http"
	"io/ioutil"
	"fmt"
	"hash/fnv"
	"io"
	"crypto/sha1"
	"math/rand"
	"time"
	"strings"
	"sort"
)

const (
	//通用签名key
	BASE_SIGN_KEY = "f7d8bc1cdb2fdf3c844d0a15026d5ccc25daec39"
	
	//上传文件上限4m
	UPLOAD_FILE_MAX_SIZE = 1 << 22
	
	//上翻页
	PAGE_TYPE_UP = 1
	
	//下翻页
	PAGE_TYPE_DOWN = 2
	
	//分页返回行数
	TABLE_LIMIT_NUM = 10
	
	//日期表开始日期
	TABLE_NAME_START_DATE = "201501"
)

const (
	
	//用户信息上传文件根目录
	UPLOAD_FILE_ROOT_USER = "0"
	
	//动态上传文件根目录
	UPLOAD_FILE_ROOT_DYNAMIC = "1"
	
	//活动上传文件根目录
	UPLOAD_FILE_ROOT_ACTIVITY = "2"
	
	//聊天上传文件根目录
	UPLOAD_FILE_ROOT_CHAT = "3"
)

const (
	
	//现居地群组
	GROUP_TYPE_BY_LIVING = "10"
	
	//职业群组
	GROUP_TYPE_BY_JOB = "11"
	
	//省编号前缀
	GEGION_TYPE_P_PREFIX  = "1"
	
	//市编号前缀
	GEGION_TYPE_C_PREFIX  = "2"
	
	//区县编号前缀
	GEGION_TYPE_D_PREFIX  = "3"
	
	//职业大类前缀
	JOB_TYPE_B_PREFIX  = "1"
	
	//职业小类前缀
	JOB_TYPE_S_PREFIX  = "2"
	
)

const (
	//求乡音认证申请
	REMIND_MESSAGE_TYPE_A = "1"
	//被认证提醒
	REMIND_MESSAGE_TYPE_B = "2"
	//乡音团队
	REMIND_MESSAGE_TYPE_C = "3"
	//活动中心
	REMIND_MESSAGE_TYPE_D = "4"
	//新增同乡
	REMIND_MESSAGE_TYPE_E = "5"
	//关注乡友新增动态
	REMIND_MESSAGE_TYPE_F = "6"
)
//提醒消息类型定义
var RemindMessageTypeDefine = map[string]string{
	REMIND_MESSAGE_TYPE_A:"求乡音认证申请",
	REMIND_MESSAGE_TYPE_B:"被认证提醒",
	REMIND_MESSAGE_TYPE_C:"乡音团队",
	REMIND_MESSAGE_TYPE_D:"新增活动",
	REMIND_MESSAGE_TYPE_E:"新增同乡",
	REMIND_MESSAGE_TYPE_F:"乡友动态",	
}


func PageErrorSet(){
	beego.Errorhandler("401",PageErrorInfo)
	beego.Errorhandler("402",PageErrorInfo)
	beego.Errorhandler("403",PageErrorInfo)
	beego.Errorhandler("404",PageErrorInfo)
	beego.Errorhandler("405",PageErrorInfo)
	beego.Errorhandler("501",PageErrorInfo)
	beego.Errorhandler("502",PageErrorInfo)
	beego.Errorhandler("503",PageErrorInfo)
	beego.Errorhandler("504",PageErrorInfo)
}

func PageErrorInfo(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("page error"))
}



func Substr(str string, length int) string {
	res := ""
	if str == "" {
		return res
	}
	if length <= 0 {
		return str
	}
	for k, v := range str {
		if k >= length {
			break
		}
		res += string(v)
	}
	return res
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func HttpGet(url string) ([]byte, error) {
	trans := &http.Transport{
		Dial:TimeoutDialer(60 * time.Second,60 * time.Second),
	}
	client := http.Client{
		Transport:trans,
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "xyAPIServer")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(body) <= 0 {
		return nil, errors.New("get data is null")
	}
	return body, nil

}

//检查sign
func CheckSign(req *context.Context, sign string,delParame []string) bool {
	
	return CheckLoginSign(req,sign,BASE_SIGN_KEY,delParame)
}
//因为urlencode问题改为如下
func CheckLoginSign(c *context.Context, sign,key string,delParame []string) bool {
	req := c.Request
	if req == nil{
		fmt.Printf("%#v\n",req)
		return false
	}
	msg := make([]string,0,10)
	switch req.Method {
		case "GET":
			
		    if t := strings.Split(req.URL.RawQuery,"&"); t != nil {
				has := false
				for _,p := range t {
					labContinue:
					if has {
						has = false
						continue
					} 
					for _,v := range delParame {
						if strings.HasPrefix(p,fmt.Sprintf("%s=",v)) {
							has = true
							goto labContinue
						}
					}
					msg = append(msg,p)
				}
			}	
			fmt.Printf("client get RawQuery:%s\n",req.URL.RawQuery) 			 
		case "POST":
		
		   if strings.Contains(req.Header.Get("Content-Type"),"multipart/form-data") {

				for _,v := range delParame {
					req.Form.Del(v)		
				}
				for k, v := range req.Form {
					x := ""
					if len(v) > 0 {
						x = v[0]
					}
					msg = append(msg,fmt.Sprintf("%s=%s",k,x ))
				}
				
			}else{
				
				if len(c.Input.RequestBody) == 0 {
					return false
				}
				if t := strings.Split(string(c.Input.RequestBody),"&"); t != nil {
					has := false
					for _,p := range t {
						labPostContinue:
						if has {
							has = false
							continue
						} 
						for _,v := range delParame {
							if strings.HasPrefix(p,fmt.Sprintf("%s=",v)) {
								has = true
								goto labPostContinue
							}
						}
						msg = append(msg,p)
					}
				}
				
			}
		default:
			fmt.Printf("%s\n","not method")
			return false
		
	}
	if len(msg) == 0 {
		fmt.Printf("%s\n","not recv data")
		return false
	}
	sort.Strings(msg)
	mac := hmac.New(sha256.New, []byte(key))
	check := strings.Join(msg,"&")
	mac.Write([]byte(check))
	h := mac.Sum(nil)
	code, err := hex.DecodeString(sign)
	if err != nil {
		println(err.Error())
		return false
	}
	fmt.Printf("msg:%s \t server sign=%x \t client sign=%s\n ",check,h,sign)
	return hmac.Equal(code, h)
}

//账号转uid
func ConvertAccountToUID(account string) uint32 {
	h := fnv.New32a()
	io.WriteString(h,account)
	return h.Sum32()
}

//登录密码hash
func HashLoginPassword(account,pwd string) string {
	f := fnv.New32a()
	io.WriteString(f,account)
	d := f.Sum32()
	p := fmt.Sprintf("%s%d",pwd,d)
	sh := sha1.New()
	io.WriteString(sh,p)
	h := sh.Sum(nil)
	return fmt.Sprintf("%x",h)
}

//生成19位id前缀
func GenerateNineteenPrefixID() string {
	s := fmt.Sprintf("%d",time.Now().UnixNano())
	l := len(s)
	sub := 19 - l
	suffix := ""
	if sub > 0 {
		for i := 0 ; i < sub ; i++ {
			suffix += "0"
		}
		s = fmt.Sprintf("%s%s",s,suffix)
	}
	return s[:19]
}

//生成随机token
func GenerateToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := fmt.Sprintf("%d%d",time.Now().UnixNano() ^ 0x19C6F3D4DF ,r.Int31n(99999999))
	sha := sha1.New()
	io.WriteString(sha,s)
	return fmt.Sprintf("%x",sha.Sum(nil))
}

//上传文件目录hash
func GetUpLoadFileNameAndPath(uid uint32,ym string,fileType string) (string,string) {
	h := uid % 256
	fName := fmt.Sprintf("user_voice_%d_%d_%s_%d.%s",uid,h,ym,time.Now().UnixNano(),fileType)
	fPath := fmt.Sprintf("user/voice/%d/%d/%s",h,uid,ym)
	return fName,fPath
}

//上传用户信息乡音文件目录hash
func GetUserUpLoadFileNameAndPath(uid uint32,ym string,fileType string) (string,string) {
	h := uid % 256
	fName := fmt.Sprintf("%s_%s_%d_%d_%s_%d.%s",UPLOAD_FILE_ROOT_USER,UPLOAD_FILE_ROOT_USER,h,uid,ym,time.Now().UnixNano(),fileType)
	fPath := fmt.Sprintf("%s/%s/%d/%d/%s",UPLOAD_FILE_ROOT_USER,UPLOAD_FILE_ROOT_USER,h,uid,ym)
	return fName,fPath
}


//上传活动文件目录hash
func GetActivityUpLoadFileNameAndPath(aid int64,uid uint32,ym string,fileType string) (string,string) {
	h := uid % 256
	fName := fmt.Sprintf("%s_%d_%d_%d_%s_%d.%s",UPLOAD_FILE_ROOT_ACTIVITY,aid,h,uid,ym,time.Now().UnixNano(),fileType)
	fPath := fmt.Sprintf("%s/%d/%d/%d/%s",UPLOAD_FILE_ROOT_ACTIVITY,aid,h,uid,ym)
	return fName,fPath
}

//上传动态文件目录hash
func GetDynamicUpLoadFileNameAndPath(HomeProvinceID int,uid uint32,ym string,fileType string) (string,string) {
	h := uid % 256
	fName := fmt.Sprintf("%s_%d_%d_%d_%s_%d.%s",UPLOAD_FILE_ROOT_DYNAMIC,HomeProvinceID,h,uid,ym,time.Now().UnixNano(),fileType)
	fPath := fmt.Sprintf("%s/%d/%d/%d/%s",UPLOAD_FILE_ROOT_DYNAMIC,HomeProvinceID,h,uid,ym)
	return fName,fPath
}

//随机昵称
func GenerateRandomNickName() string {
	
	nickName := []string{"炸酱面", "驴打滚儿", "中南海豆汁儿", "烤鸭", "糖葫芦", "二锅头", "景泰蓝", "果脯", "茯苓饼","狗不理", "十八街", "耳朵眼", "煎饼馃子", "茶汤儿", "锅巴菜", "独流老醋", "栗子", "猫不闻饺子", "烫面炸糕",
		"四喜烤麸", "大中华", "生煎", "排骨年糕", "蟹粉小笼", "狮子头", "阳春面", "浦东鸡", "鲜肉粽", "腌笃鲜","水煮鱼", "麻辣烫", "毛血旺", "怪味胡豆", "涪陵榨菜", "辣子鸡", "米花糖", "陈麻花", "桃片", "泡椒凤爪",
		"驴肉火烧", "金丝大枣", "油面窝窝", "焖饼", "杏梅", "柿饼", "素面", "春不老", "松花蛋", "豆腐筋","刀削面", "栲栳栳", "老陈醋", "豆面饸饹", "太谷饼", "酱肉", "沁州黄", "平遥牛肉", "过油肉", "竹叶青",
		"烤全羊", "酪蛋子", "驼掌", "马奶酒", "发菜", "奶茶", "奶皮", "勒巴达", "牛肉干", "豆包","虾爬子", "红参", "酸枣仁", "海蛎子", "粘豆包", "麻辣拌", "大酱", "大骨鸡", "蘸酱菜", "大米",
		"土豆", "溜肉段", "打糕", "辣白菜", "人参", "黑木耳", "鹿茸", "冷面", "猴头菇", "大白菜","挂浆冰溜子", "杀猪菜", "乱炖", "锅包肉", "大列巴", "红肠", "酸菜", "酱骨头", "炖粉条儿", "汆白肉",
		"杠头", "芙蓉鸡", "九转大肠", "甏肉饭", "金枣", "苹果", "锅塌豆腐", "阿胶", "大葱", "馓子","酱排骨", "猪肉脯", "烧饼", "炒饭", "锅盖面", "芝麻酱", "活珠子", "茴香豆", "八珍糕", "咸水鸭",
		"牛肉汤", "芙蓉糕", "臭鳜鱼", "格拉条", "大救驾", "包公鱼", "虾子面", "墨子酥", "板面", "水烙馍","汤团", "火腿", "叫花鸡", "松糕", "牛肉羹", "叔嫂传珍", "熏鱼", "年糕", "麻糍", "酥饼",
		"粉蒸肉", "脐橙", "三杯鸡", "酸枣糕", "瓦罐汤", "糯米饭", "萝卜饼", "酥糖", "炒粉", "鸡腿","佛跳墙", "蚵仔煎", "沙茶面", "卷饼", "锅边糊", "扁肉", "金骏眉", "粿条", "茯苓糕", "手抓面",
		"芋圆", "卤肉饭", "太阳饼", "棺材板", "凤梨酥", "麻薯", "牛轧糖", "鸭肉粥","白吉馍", "烧鸡", "鲤鱼跳龙门", "三不粘", "烩面", "胡辣汤", "花生糕", "桶子鸡", "葱花饼",
		"热干面", "豆皮", "鸭脖子", "米酒", "黄鹤楼", "糯米鸡", "面窝", "武昌鱼", "麻糖", "白马尿","口味虾", "臭豆腐", "剁椒鱼头", "猪血丸子", "米线", "红烧肉", "腊肉", "团子", "酱板鸭", "粑粑",
		"月饼", "藕饼", "虾仁饼", "地瓜饼", "双皮奶", "橄榄菜", "桂圆肉","烤乳猪", "酱豆腐", "米粉", "螺蛳粉","肉夹馍", "biangbiang", "羊肉泡馍", "凉皮", "臊子面", "油酥饼", "千层饼", "酸汤饺", "葫芦鸡", "金线油塔", "钱钱饭", "窝窝",
		"醪糟", "酿皮", "老酸奶", "手抓饭", "狗浇尿", "尕面片", "青稞酒", "干板鱼 ", "牦牛干", "奶皮",
		"水煮牛肉", "担担面", "口水鸡", "夫妻肺片", "酸菜鱼", "泡菜", "棒棒鸡", "豆瓣酱", "龙抄手", "串串", "兔头", "麻婆豆腐","丝袜奶茶", "叉烧包", "老婆饼", "铜锣烧", "烧腊","蛋挞", "豆捞",
		"拉面", "打卤面", "浆水面", "炮仗面", "茄鲞", "蜜瓜","老干妈", "茅台", "狗肉", "丝娃娃", "牛肉粉", "酸汤鱼", "宫保鸡丁", "折耳根", "刷把头", "猪儿粑", "波波糖",
		"葡萄干", "大盘鸡", "羊肉串", "拉条子", "馕包肉", "曲曲汤", "馕饼","公婆饼", "风干肉", "炸灌肺", "吧啦饼", "炖羊肉", "糌粑", "红花",
		"过桥米线", "菠萝饭", "凉糕","饸饹", "油香", "枸杞"}
		
	dataLen := len(nickName) - 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(dataLen)
	return nickName[index]
	
}

// 随机图像
func GenerateRandomAvatar() string {
	avatar := []string{
		"0_0_0_0_0_1.png",
		"0_0_0_0_0_2.png",
		"0_0_0_0_0_3.png",
		"0_0_0_0_0_5.png",
		"0_0_0_0_0_7.png",
	}
	
	dataLen := len(avatar) - 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(dataLen)
	return avatar[index]
}

//随机机器人
func RobotRandom()uint32{
	arr := []uint32{
		 1073925456,1124258313,1107480694,1157813551,1141035932,1191368789,1174591170,
         1224924027,1208146408,3674206138,3690983757,3640650900,3657428519,3607095662,
 		 3623873281,3573540424,3590318043,3539985186,3556762805,1292328693,1275551074,
 		 1258773455,2214527778,2231305397,2180972540,2197750159,2147417302,2164194921,
 		 2113862064,2130639683,2348748730,2365526349,2315340587,2298562968,2348895825,
		 2332118206,2382451063,2365673444,2416006301,2399228682,2181119635,2164342016,
 		 268721332,285498951,302276570,163281110,180058729,129725872,146503491,230391586,247169205,
	}
	dataLen := len(arr) - 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(dataLen)
	return arr[index]
}