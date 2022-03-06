package lib

import (
	"github.com/yangsen996/gRedis/greids"
	"sync"
	"time"
)

var NewPool *sync.Pool

func init() {
	NewPool = &sync.Pool{
		New: func() interface{} {
			return greids.NewSimpleCache(greids.NewStringOperation(), time.Second*15)
		},
	}
}

func NewCache() *greids.SimpleCache {
	return NewPool.Get().(*greids.SimpleCache)
}

func RelaseNew(cache *greids.SimpleCache) {
	NewPool.Put(cache)
}
