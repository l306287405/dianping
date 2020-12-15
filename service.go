package dianping

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

const (
	OPENAPI_URL = "https://openapi.dianping.com/"
)

//配置
type Config struct {
	AppKey string
	Secret string
	AuthFactor string
}

func (s *Config) check() error {
	if s.AppKey == "" {
		return errors.New("appkey没有配置")
	}
	if s.Secret == "" {
		return errors.New("secret没有配置")
	}
	if s.AuthFactor == ""{
		s.AuthFactor = RandStr(6)
	}
	return nil
}

//服务
type Service struct {
	Config
	Cache
}

func NewService(cfg Config,cache Cache) (*Service,error) {
	err:=cfg.check()
	if err!=nil{
		return nil, err
	}
	if cache==nil{
		cache=NewDefStore()
	}
	return &Service{cfg,cache},nil
}

func (s *Service) AuthEncryptState(state string) string{
	return state+s.AuthFactor
}

//参数类型
type ReqParams struct {
	url.Values
}

func (s *ReqParams) SetInt(key string,val int){
	s.Set(key,strconv.Itoa(val))
}

func (s *ReqParams) SetLong(key string,val int64) {
	s.SetInt64(key,val)
}

func (s *ReqParams) SetInt64(key string,val int64) {
	s.Set(key,strconv.FormatInt(val,10))
}

//获取参数签名
func (s *ReqParams) Sign(secret string) string {
	var (
		keys []string
		signStr = secret
		k string
		md5Ctx = md5.New()
		cipherStr []byte
	)


	for k, _ = range s.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k = range keys {
		if s.Get(k) != "" {
			signStr += k + s.Get(k)
		}
	}
	signStr+=secret
	md5Ctx.Write([]byte(signStr))
	cipherStr = md5Ctx.Sum(nil)
	return strings.ToLower(hex.EncodeToString(cipherStr))
}

func (s *ReqParams) AddKeyAndSecret(cfg *Config){
	s.Set("app_key",cfg.AppKey)
	s.Set("app_secret",cfg.Secret)
}
