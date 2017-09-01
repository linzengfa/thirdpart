package wxlogin

import (
	"net/http"
	"errors"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

const (
	WXLOGIN_RETURN_ERRCODE_KEY ="errcode"
	WXLOGIN_RETURN_ERRMSG_KEY ="errmsg"
	WXLOGIN_RETURN_OPENID_KEY ="openid"
	WXLOGIN_RETURN_SESSIONKEY_KEY ="session_key"
)
var(
	wxlogin_param_error = errors.New("wxlogin_param_error")
	wxlogin_watermark_check_error = errors.New("encryptedData error")
)
type WXLogin struct {
	appId string
	appSecret string
}

type WXUserInfo struct {
	OpenId string `json:"openId"`
	NickName string `json:"nickName"`
	Gender byte `json:"gender"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId string `json:"unionId"`
}

type wxAuthorization struct {
	Openid string `json:"openid"` //用户唯一标识
	Session_key string `json:"session_key"` //会话密钥
	Unionid string `json:"unionid"` //用户在开放平台的唯一标识符。
}

func New(appId, appSecret string) (wx *WXLogin,err error) {
	if appId == "" || appSecret==""{
		err = wxlogin_param_error
		return
	}
	return &WXLogin{appId, appSecret},nil
}

//微信登录
//加密数据( encryptedData )进行解密
func (wx *WXLogin)Login(encryptedData,iv,code string)(wxu *WXUserInfo,err error){
	wxa,err:=codeExchangeSessionkey(wx.appId,wx.appSecret,code)
	if err!=nil{
		return
	}

	wxDataCrypt := NewWXDataCrypt(wx.appId,wxa.Session_key)

	wxData,err:=wxDataCrypt.decryptData(encryptedData,iv)
	if err !=nil{
		return
	}
	if wxData.Appid =="" ||wxData.Appid !=wx.appId {
		err = wxlogin_watermark_check_error
		return
	}
	wxu =&WXUserInfo{
		wxData.OpenId,
		wxData.NickName,
		wxData.Gender,
		wxData.City,
		wxData.Province,
		wxData.Country,
		wxData.AvatarUrl,
		wxData.UnionId,
	}
	return
}

//code换取session_key
func codeExchangeSessionkey(appId string,appSecret string,code string)(wxa *wxAuthorization,err error){
	url:=buildExchangeUrl(appId,appSecret,code)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err !=nil{
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		return
	}
	var resultMap map[string]interface{}

	err=json.Unmarshal(body,&resultMap)
	if err !=nil{
		fmt.Println("codeExchangeSessionkey error %v",err)
		return
	}

	//code换取session_key失败
	if _,ok:=resultMap[WXLOGIN_RETURN_ERRCODE_KEY];ok{
		errmsg :=fmt.Sprintf("codeExchangeSessionkey error，errcode=%v,errmsg=%v",resultMap[WXLOGIN_RETURN_ERRCODE_KEY],resultMap[WXLOGIN_RETURN_ERRMSG_KEY])
		err= errors.New(errmsg)
		return
	}
	if _,ok:=resultMap[WXLOGIN_RETURN_SESSIONKEY_KEY];!ok{
		err= errors.New("codeExchangeSessionkey error,session_key does not exist")
		return
	}
	if _,ok:=resultMap[WXLOGIN_RETURN_OPENID_KEY];!ok{
		err= errors.New("codeExchangeSessionkey error,openid does not exist")
		return
	}
	err = json.Unmarshal(body,&wxa)
	return
}

//组织code换取session_key请求url
func buildExchangeUrl(appId string,appSecret string,code string)(string){
	return fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",appId,appSecret,code)
}