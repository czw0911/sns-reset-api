package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"XYAPIServer/XYLibs"
	"fmt"
	"encoding/json"
	"time"
	"strconv"
	"XYAPIServer/XYShareServer/libs"
	"XYAPIServer/XYShareServer/models"
)

var (
	dynamic_share_cache_key = "dynamicShareCache"
	activity_share_cache_key = "activityShareCache"
) 

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	q := c.Ctx.Request.URL.RawQuery
	c.Data["downloadUrl"] = beego.AppConfig.String("appDownloadUrl")
	wx := libs.NewWXMPjsSdk(beego.AppConfig.String("mp_appid"),beego.AppConfig.String("mp_appsecret"),libs.RedisDBShare)
	c.Data["wxMP"] = wx.GetSignPackage(c.Ctx.Request)
	switch{
		case strings.HasPrefix(q,"a="):
			c.shareActivity(c.GetString("a"))
			c.TplNames = "activity.tpl"
		case strings.HasPrefix(q,"x="):
			c.shareActivity(c.GetString("x"))
			c.TplNames = "activity_detail.tpl"
		case strings.HasPrefix(q,"d="):
			c.shareDynamic(c.GetString("d"))
			c.TplNames = "dynamic.tpl"
		default:
			shareData := make(map[string]string)
			shareData["title"] = "乡音，真实而暖心的同乡社交应用！"
			shareData["desc"] = "有一种情怀叫乡音，有一种思念叫乡味！我们在乡音等你，一起品味乡情！"
			shareData["link"] = "http://share.xiangyin.im"
			shareData["imgUrl"] = "http://share.xiangyin.im/img/share.png"
			c.Data["shareData"] = shareData
			c.TplNames = "index.tpl"
			
	}
	
}

//动态分享
func  (c *MainController) shareDynamic(dynamicID string) {
		if dynamicID == ""{
			return
		}
		data := new(models.DynamicCache)
		key := fmt.Sprintf("%s:%s",dynamic_share_cache_key,dynamicID)
		cache,err := libs.RedisDBShare.Get(key)
		//fmt.Printf("%#v\n",cache)
		if cd,ok :=  cache.([]byte);ok {
			var tmp models.DynamicCache
			e := json.Unmarshal(cd,&tmp)
			if e == nil {
				c.Data["Dynamic"] = tmp.Dynamic
				c.Data["Comment"] = tmp.Comment
			}else{
				beego.Error(e)
			}
		}else{
				
			    beego.Error(err)
				url := fmt.Sprintf("http://%s?DynamicID=%s",beego.AppConfig.String("dynamic_addr"),dynamicID)
				resp ,err := XYLibs.HttpGet(url)
				if len(resp) > 0 {
					var res XYLibs.XYAPIResponse
					err = json.Unmarshal([]byte(resp),&res)
					if err != nil {
						println(string(resp))
						beego.Error(err)
						return
					}
					if res.Code == 1 {
					  if dContent,ok := res.Info.(map[string]interface{});ok {
						data.Dynamic = make(map[string]string)
						 if dd,ok := dContent["dynamic"].(map[string]interface{});ok {
							sc,_ := dd["DynamicContent"].(string)
							if sc != "" {
								data.Dynamic["DynamicContent"] = strings.Replace(sc,"\\r\\n","\r\n",-1)
							}			
							data.Dynamic["ViewNUM"],_ = dd["ViewNUM"].(string)
							data.Dynamic["ForwardNum"],_ = dd["ForwardNum"].(string)
							data.Dynamic["CommentNum"],_ = dd["CommentNum"].(string)
							data.Dynamic["GoodNUM"],_ = dd["GoodNUM"].(string)
							data.Dynamic["Images"],_ = dd["Images"].(string)
							data.Dynamic["Voices"],_ = dd["Voices"].(string)
							if avatar, o := dd["PostUser"].(map[string]interface{});o {
								data.Dynamic["Avatar"],_ = avatar["Avatar"].(string)
								data.Dynamic["NickName"],_ = avatar["NickName"].(string)
							}
							
							postTime,_ := dd["PostTime"].(string)
							if postTime != "" {
								t,_ := strconv.ParseInt(postTime,10,64)
								if t > 0 {
									n := time.Unix(t,0)		
									data.Dynamic["PostTime"] = n.Format("2006年01月02日")
								}
								
							}
							c.Data["Dynamic"] = data.Dynamic
						}
						//fmt.Printf("%#v\n",dContent)
						//评论
						if dd,ok := dContent["comment"].([]interface {});ok {
							data.Comment = make([]map[string]string,0,len(dd))
							for _,v := range dd {
								if cc,ok := v.(map[string]interface{});ok {
									//fmt.Printf("%#v\n\n\n",cc["PostUser"])
									 t := make(map[string]string)
									 if avatar, o := cc["PostUser"].(map[string]interface{});o {
										t["Avatar"],_ = avatar["Avatar"].(string)
										t["NickName"],_ = avatar["NickName"].(string)
									 }
									 t["Contents"],_ = cc["Contents"].(string)
									 postTime,_ := cc["PostTime"].(string)
									if postTime != "" {
										y,_ := strconv.ParseInt(postTime,10,64)
										if y > 0 {
											n := time.Unix(y,0)		
											t["PostTime"] = n.Format("2006年01月02日")
										}
										
									}
									 data.Comment = append(data.Comment,t)
								}
							}
							
							c.Data["Comment"] = data.Comment
							//fmt.Printf("%#v\n",comment)
						}
						if len(c.Data) > 0 {
							cacheData,err := json.Marshal(data)
							if err != nil {
								beego.Error(err)
							}else{
								libs.RedisDBShare.SETEX(key,86400,cacheData)
							}
						}			 
					  }
					}
				}
	   }
		link := fmt.Sprintf("%s%s",c.Ctx.Input.Site(),c.Ctx.Request.URL.RequestURI())
		shareData := make(map[string]string)
		shareData["title"] = "来自“乡音”的家乡动态"
		shareData["desc"] = c.Data["Dynamic"].(map[string]string)["DynamicContent"]
		shareData["link"] = link
		shareData["imgUrl"] = c.Data["Dynamic"].(map[string]string)["Images"]
		c.Data["shareData"] = shareData
} 

//活动分享
func  (c *MainController) shareActivity(talkID string) {
	if talkID == "" {
		return
	}
	
	key := fmt.Sprintf("%s:%s",activity_share_cache_key,talkID)
	cache,err := libs.RedisDBShare.Get(key)
	if err != nil {
		beego.Error(err)
	}
	if cd,ok :=  cache.([]byte);ok {
		
		var tmp map[string]interface{}
		e := json.Unmarshal(cd,&tmp)
		if e == nil {
				c.Data["Activity"] = tmp
		}else{
				beego.Error(e)
		}
		
	}else{
		url := fmt.Sprintf("http://%s?TalkID=%s",beego.AppConfig.String("activity_addr"),talkID)
		resp ,err := XYLibs.HttpGet(url)
		if err != nil {
			beego.Error(err)
		}
		if resp != nil {
			var res XYLibs.XYAPIResponse
			err = json.Unmarshal([]byte(resp),&res)
			if err != nil {
				//println(string(resp))
				beego.Error(err)
				return
			}
			if res.Code == 1 {
				if data,ok := res.Info.([]interface{});ok {
					if val,ok := data[0].(map[string]interface {});ok{
						if _,ok := val["TalkContent"];ok {
							val["TalkContent"] = strings.Replace(val["TalkContent"].(string),"\\r\\n","\r\n",-1)
						}
						c.Data["Activity"] = val
						cacheData,err := json.Marshal(val)
						if err != nil {
								beego.Error(err)
						}else{
								libs.RedisDBShare.SETEX(key,86400,cacheData)
						}
					}
				}
			}
		}
		
	}
	link := fmt.Sprintf("%s%s",c.Ctx.Input.Site(),c.Ctx.Request.URL.RequestURI())
		shareData := make(map[string]string)
		shareData["title"] = "来自“乡音”的醇浓乡味"
		shareData["desc"] = c.Data["Activity"].(map[string]interface{})["TalkContent"].(string)
		shareData["link"] = link
		shareData["imgUrl"] = c.Data["Activity"].(map[string]interface{})["Images"].(string)
		c.Data["shareData"] = shareData
	
}