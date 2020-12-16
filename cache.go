package dianping

import (
	"strconv"
	"sync"
	"time"
)

var _ Cache = (*DefStore)(nil)

type DefStore struct {
	sync.Map
}

func (s *DefStore) Set(key string, value string, expiredIn time.Duration) {
	s.Store(key, []string{value, strconv.FormatInt(time.Now().Add(expiredIn).Unix(), 10)})
}

func (s *DefStore) Get(key string) string {
	value, loaded := s.LoadAndDelete(key)
	if !loaded {
		return ""
	}
	values, ok := value.([]string)
	if !ok {
		return ""
	}
	expiredAt, err := strconv.ParseInt(values[1], 10, 64)
	if err != nil {
		return ""
	}
	if time.Now().Unix() > expiredAt {
		return ""
	}
	return values[0]
}

func NewDefStore() *DefStore {
	return &DefStore{}
}

type Cache interface {
	//键值以及过期时间 单位为秒
	Set(key string, value string, expiredIn time.Duration)

	//获取缓存,失效或者不存在均返回空字符串
	Get(key string) string
}
