/*
*@Description 二维码登陆（后台管理员)
*@Contact czw@outlook.com
*/

package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"XYAPIServer/XYQRLogin/libs"
	"XYAPIServer/XYLibs"
	"encoding/base32"
	"crypto/aes"
	"crypto/cipher"
	"image/png"
	"bytes"
	"log"
	"net/http"
	"time"
	"fmt"
	"math/rand"
	"crypto/sha1"
	"io"
	"strings"
	"strconv"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

const (
	qrkey = "YzQ3ZGFhYTE4MWY1"
	hmacKey = "vbXyK07*7W51I5%^0fjJCN!q5Xt392ML4sr26t0T4nrJqkNLJe3zyf688essEys&"
	loginHmacKey = "OfADbzQ0SG9HiKwF3ifEgM9Q2wQgvCOlCcMwnSdxmo8JX6wb"
	mysql_xy_db_admin = "xy_db_admin"
)

var (
	Cfg *libs.AppConfig
	TokenStorage *libs.TokenList
	redisDB *XYLibs.RedisHash
	adminDB orm.Ormer
)
type ResponseMsg struct {
	Code int
	Desc string
}

func (self *ResponseMsg) Bytes() []byte {
	data,err := json.Marshal(self)
	if err != nil {
		return nil
	}
	return data
}

type QRHandler struct{
	   AccessToken []byte
}

func (self *QRHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		log.Println(r.URL.String())
		switch r.URL.Path {
			case "/ws":
			 	serveWs(w,r)
			case "/verify":
				self.VerifyQRCode(w,r)
			default:
				self.AccessToken = GenerateToken()
				self.EchoQRCode(w,r)
		}
}

func (self *QRHandler) EchoQRCode(w http.ResponseWriter, r *http.Request) {
		
		resp := new(ResponseMsg)
		sid := strings.TrimSpace(r.URL.Query().Get("SID"))
		date := strings.TrimSpace(r.URL.Query().Get("Date"))
		sign := strings.TrimSpace(r.URL.Query().Get("Sign"))
		
		if sid == "" || sign == "" || date == ""{
			resp.Code = 2
			resp.Desc = "param number error..."
			w.Write(resp.Bytes())
			return
		}
		log.Printf("sid:%s \t date:%s \t sign:%s\n",sid,date,sign)
		d ,e :=strconv.ParseInt(date,10,64)
		if e != nil {
			resp.Code = 3
			resp.Desc = "date param error..."
			w.Write(resp.Bytes())
			return
		}
		sub := time.Now().Unix() - d
		if sub > 60 {
			resp.Code = 4
			resp.Desc = "date param expire..."
			w.Write(resp.Bytes())
			return
		}
		check := fmt.Sprintf("Date=%s&SID=%s",date,sid)
		res := CheckSign(check,sign,hmacKey)
		if !res {
			resp.Code = 5
			resp.Desc = "sign verify error..."
			w.Write(resp.Bytes())
			return
		}
		
		token := fmt.Sprintf("%x",GenerateToken())
		key := fmt.Sprintf("qr:%s",sid)
		isSave := redisDB.SETEX(key,3600,token)
		if isSave != nil {
			resp.Code = 6
			resp.Desc = "qr token save fail..."
			w.Write(resp.Bytes())
			log.Fatal(isSave)
			return
		}
			
		mm := make(map[string]string)
		mm["url"] = Cfg.RunDomain
		mm["qr"] = token
		mm["sid"] = sid
		out,_ := json.Marshal(mm)
		desc,e := AesEncrypt(out,[]byte(qrkey))
		if e != nil {
			resp.Code = 7
			resp.Desc = "qr contents cipher fail..."
			w.Write(resp.Bytes())
			log.Fatal(e.Error())
			return
		}
		log.Printf("%v\n",mm)
		content := base32.StdEncoding.EncodeToString(desc)
		log.Println(content)
		code, err := qr.Encode(content, qr.L, qr.Unicode)
		if err != nil {
			log.Fatal(err)
		}
		bar, e := barcode.Scale(code, 200, 200)
		if e != nil {
			log.Fatal(e)
		}
		w.Header().Set("Content-Type","image/png")
		png.Encode(w, bar)
}

func (self *QRHandler) VerifyQRCode(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm()
	sid := r.Form.Get("SID")
	appID ,_ :=  strconv.ParseUint(r.Form.Get("APPID"),10,64)
	qrCode := r.Form.Get("QR")
	sign := r.Form.Get("Sign")
	
	log.Println(sid,appID,qrCode,sign)
	var resDB []orm.Params
	adminDB.Raw("SELECT  AccessToken FROM sys_users WHERE ID = ?  LIMIT 1",appID).Values(&resDB)
	
	if len(resDB) == 0 {
		w.Write([]byte("6"))
		return
	}
	hKey := resDB[0]["AccessToken"].(string)
	
	params := fmt.Sprintf("APPID=%d&QR=%s&SID=%s",appID,qrCode,sid)
	res := CheckSign(params,sign,hKey)
	if !res {
		w.Write([]byte("2"))
		return
	}
	key := fmt.Sprintf("qr:%s",sid)
	token ,err := redisDB.Get(key)
	if err != nil {
		log.Println(err)
		w.Write([]byte("3"))
		return
	}
	sQR := ""
	switch token.(type) {
		case []byte :
			sQR = string(token.([]byte))
		default :
			w.Write([]byte("4"))
			return
	}
	redisDB.Del(key)
	
	if qrCode != ""  &&  sQR == qrCode {
		
		token := fmt.Sprintf("%x",GenerateToken())
		keyLogin := fmt.Sprintf("login:%s",token)
		isSave := redisDB.SETEX(keyLogin,120,strconv.FormatUint(appID,10))
		if isSave != nil {
			log.Fatal(isSave)
			w.Write([]byte("7"))
			return
		}
		loginParam := fmt.Sprintf("at=%s&d=%d",token,time.Now().Unix())
		loginSign := GenerateSign(loginParam,loginHmacKey)
		loginUrl := fmt.Sprintf("%s?%s&s=%s",Cfg.LoginCallbackUrl,loginParam,loginSign)
		
		libs.HubWS.Broadcast <- &libs.Msg{ID:sid,Contents:[]byte(loginUrl)}
		
		w.Write([]byte("1"))
		
		
	}else{
		
		w.Write([]byte("5"))
	}
	
	
	
	
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	sid := strings.TrimSpace(r.URL.Query().Get("SID"))
	date := strings.TrimSpace(r.URL.Query().Get("Date"))
	sign := strings.TrimSpace(r.URL.Query().Get("Sign"))
	
	check := fmt.Sprintf("Date=%s&SID=%s",date,sid)
	res := CheckSign(check,sign,hmacKey)
	if !res {
		http.Error(w, "sign verify fail", 403)
		return
	}
	
	ws, err := libs.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	
	c := &libs.ConnWS{Send: make(chan []byte, 256), WS: ws,ID:sid}
	libs.HubWS.Register <- c
	go c.WritePump()
	c.ReadPump()
}

func GenerateToken() []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := fmt.Sprintf("%d%d",time.Now().UnixNano() ^ 0x19C6F3D8DF ,r.Int31n(99999999))
	sha := sha1.New()
	io.WriteString(sha,s)
	return sha.Sum(nil)
}

func CheckSign(check,sign,key string) bool {
		hash := hmac.New(sha256.New,[]byte(key))
		hash.Write([]byte(check))
		desc := hash.Sum(nil)
		src,err := hex.DecodeString(sign)
		if err != nil {
			return false
		}
		log.Printf("msg:%s \t client sign:%s \t server sign:%s\n",check,sign,fmt.Sprintf("%x",desc))
		return hmac.Equal(src,desc)	
}

func GenerateSign(data,key string) string {
		hash := hmac.New(sha256.New,[]byte(key))
		hash.Write([]byte(data))
		desc := hash.Sum(nil)
		return fmt.Sprintf("%x",desc)
}


func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func main() {
	port := fmt.Sprintf(":%d",Cfg.RunPort)
	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("ws"))
	})
	s := http.Server{
		Addr:           port,
		Handler:        &QRHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go libs.HubWS.Run()
	log.Printf("qr server runing on %s",port)
	log.Fatal(s.ListenAndServe())
}

func init(){
	var err error
	Cfg ,err = libs.NewAppConfig("app.json")
	if err != nil {
		panic(err)
	}
	TokenStorage = libs.NewTokenList()
	
	redisDB = XYLibs.NewRedisHash()
	redisDB.ConnectRedis(Cfg.QRRedisAddr)
	
	//
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase(mysql_xy_db_admin,"mysql",Cfg.MySqlAdmin)
	orm.RegisterDataBase("default", "mysql", Cfg.MySqlAdmin)
	orm.Debug = true
	adminDB = orm.NewOrm()
	adminDB.Using(mysql_xy_db_admin)
}
