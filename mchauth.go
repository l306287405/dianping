package dianping

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const (
	MCH_AUTH_URL    = "https://e.dianping.com/dz-open/merchant/auth?"
	OAUTH_TOKEN_URL = "https://openapi.dianping.com/router/oauth/token"
)

//商家授权ui组件
//http://open.dianping.com/document/v2?docId=6000164&rootDocId=5000
func (s *Service) MerchantAuth(r *ReqParams) string {
	r.Set("app_key", s.AppKey)
	h := md5.New()
	h.Write([]byte(s.authEncryptState(r.Get("state"))))
	s.Set(hex.EncodeToString(h.Sum(nil)), "1", time.Minute*30)
	return MCH_AUTH_URL + r.Encode()
}

type OauthTokenResp struct {
	AccessToken        string  `json:"access_token"`         //session
	ExpiresIn          int64   `json:"expires_in"`           //过期时间
	Scope              string  `json:"scope"`                //session的权限范围,对应模块名称
	RefreshToken       string  `json:"refresh_token"`        //即为refresh_session，授权可刷新次数用完后，将不再返回新的refresh_session
	RemainRefreshCount int     `json:"remain_refresh_count"` //剩余刷新次数
	Bid                string  `json:"bid"`                  //客户id
	TokenType          *string `json:"token_type,omitempty"` //bearer
	Code               int     `json:"code"`
	Msg                *string `json:"msg,omitempty"`
}

//session换取接口
//http://open.dianping.com/document/v2?docId=6000341&rootDocId=5000
func (s *Service) OauthTokenByCode(r *ReqParams) (resp *OauthTokenResp, err error) {
	resp = &OauthTokenResp{}

	err = r.CheckKeys([]string{"auth_code"})
	if err != nil {
		return
	}

	r.AddKeyAndSecret(&s.Config)
	r.Set("grant_type", "authorization_code")
	err = PostForm(OAUTH_TOKEN_URL, r, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//session刷新接口
//http://open.dianping.com/document/v2?docId=6000342&rootDocId=5000
func (s *Service) OauthTokenByRefresh(r *ReqParams) (resp *OauthTokenResp, err error) {
	resp = &OauthTokenResp{}

	err = r.CheckKeys([]string{"refresh_token"})
	if err != nil {
		return
	}

	r.AddKeyAndSecret(&s.Config)
	r.Set("grant_type", "refresh_token")
	err = PostForm(OAUTH_TOKEN_URL, r, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type OauthSessionQueryResp struct {
	Code int     `json:"code"`
	Msg  *string `json:"msg,omitempty"`
	Data struct {
		Bid   string `json:"bid"`
		Scope string `json:"scope"`
	} `json:"data"`
}

//session范围查询接口
//http://open.dianping.com/document/v2?docId=6000343&rootDocId=5000
func (s *Service) OauthSessionQuery(r *ReqParams) (resp *OauthSessionQueryResp, err error) {
	var (
		u = OPENAPI_ROUTER + "/oauth/session/query"
	)

	err = r.CheckKeys([]string{"session"})
	if err != nil {
		return
	}

	resp = &OauthSessionQueryResp{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}

type OauthSessionScopeRespBox struct {
	Code int                      `json:"code"`
	Msg  *string                  `json:"msg,omitempty"`
	Data []*OauthSessionScopeResp `json:"data"`
}

type OauthSessionScopeResp struct {
	OpenShopUuid string `json:"open_shop_uuid"` //门店id的唯一标识
	Shopname     string `json:"shopname"`       //门店名称
	Branchname   string `json:"branchname"`     //分店名称
	ShopAddress  string `json:"shop_address"`   //门店地址
	Cityname     string `json:"cityname"`       //所在城市
}

//session适用店铺查询接口
//http://open.dianping.com/document/v2?docId=6000344&rootDocId=5000
func (s *Service) OauthSessionScope(r *ReqParams) (resp *OauthSessionScopeRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/oauth/session/scope"
	)

	err = r.CheckKeys([]string{"session", "bid"})
	if err != nil {
		return
	}

	resp = &OauthSessionScopeRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}

type OauthSessionShopidmappingRespBox struct {
	Code int                              `json:"code"`
	Msg  string                           `json:"msg"`
	Data []*OauthSessionShopidmappingResp `json:"data"`
}

type OauthSessionShopidmappingResp struct {
	OpenShopUuid string `json:"open_shop_uuid"`
	ShopId       string `json:"shop_id"`
	ErrMsg       string `json:"err_msg"`
}

//session对应客户门店ID映射关系
//http://open.dianping.com/document/v2?docId=6000635&rootDocId=5000
func (s *Service) OauthSessionShopidmapping(r *ReqParams) (resp *OauthSessionShopidmappingRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/oauth/session/shopidmapping"
	)

	err = r.CheckKeys([]string{"session", "bid", "shopids"})
	if err != nil {
		return
	}

	resp = &OauthSessionShopidmappingRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}
