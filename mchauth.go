package dianping

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	MCH_AUTH_URL = "https://e.dianping.com/dz-open/merchant/auth"
)

//商家授权ui组件
//http://open.dianping.com/document/v2?docId=6000164&rootDocId=5000
func (s *Service) MerchantAuth(r *ReqParams) string{
	r.Set("app_key",s.AppKey)
	h:=md5.New()
	h.Write([]byte(s.AuthEncryptState(r.Get("state"))))
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