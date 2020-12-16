package dianping

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	MCH_AUTH_URL = "https://e.dianping.com/dz-open/merchant/auth?"
)

//商家授权ui组件
//http://open.dianping.com/document/v2?docId=6000164&rootDocId=5000
func (s *Service) MerchantAuth(r *ReqParams) string{
	r.Set("app_key",s.AppKey)
	h:=md5.New()
	h.Write([]byte(s.authEncryptState(r.Get("state"))))
	s.Set(hex.EncodeToString(h.Sum(nil)),"1",time.Minute*30)
	return MCH_AUTH_URL+r.Encode()
}

//session换取接口
//http://open.dianping.com/document/v2?docId=6000341&rootDocId=5000
func (s *Service) OauthToken(r *ReqParams)(err error){
	r.AddKeyAndSecret(&s.Config)
	result,err:=PostForm(MCH_AUTH_URL,r)
	if err!=nil{
		return err
	}
	fmt.Println(result)
	return nil
}

//session换取接口返回值
//http://open.dianping.com/document/v2?docId=6000341&rootDocId=5000
type OauthTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	Scope string `json:"scope"`
	RefreshToken string `json:"refresh_Token"`
	RemainRefreshCount int `json:"remain_refresh_count"`
	Bid string `json:"bid"`
	TokenType string `json:"token_type"`
	Code string `json:"code"`
	Msg *string `json:"msg,omitempty"`
}