package wxlogin

import (
	"encoding/base64"
	"crypto/cipher"
	"crypto/aes"
	"encoding/json"
)

type wxDataCrypt struct {
	appId string
	sessionKey string
}
type watermark struct {
	Appid string `json:"appid"`
	timestamp int64 `json:"timestamp"`
}
type wxEncryptedData  struct {
	OpenId string `json:"openId"`
	NickName string `json:"nickName"`
	Gender byte `json:"gender"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId string `json:"unionId"`
	watermark `json:"watermark"`
}

func NewWXDataCrypt(appId string,sessionKey string) *wxDataCrypt{
	return &wxDataCrypt{appId,sessionKey}
}

func (wx *wxDataCrypt)decryptData(encryptedData string,iv string)(wxed *wxEncryptedData,err error){
	sessionKeyByte, err := base64.StdEncoding.DecodeString(wx.sessionKey)
	if err != nil {
		return
	}
	encryptedDataByte, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return
	}
	ivByte, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(sessionKeyByte)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()

	blockMode := cipher.NewCBCDecrypter(block, ivByte[:blockSize])
	origData := make([]byte, len(encryptedDataByte))
	blockMode.CryptBlocks(origData, encryptedDataByte)
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	origData = origData[:(length - unpadding)]
	err=json.Unmarshal(origData,&wxed)
	return
}