package greids

import "time"

type DBGetter func() string

type SimpleCache struct {
	Operation *StringOperation
	Expire    time.Duration
	DBGetter  DBGetter
}

func NewSimpleCache(operation *StringOperation, expire time.Duration) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire}
}

// SetCache 设置缓存
func (s *SimpleCache) SetCache(key string, value interface{}) {
	s.Operation.Set(key, value, WithExpire(s.Expire))
}

func (s *SimpleCache) GetCache(key string) (ret interface{}) {
	ret = s.Operation.Get(key).UnwrapOrElse(s.DBGetter)
	s.SetCache(key, ret)
	return
}
