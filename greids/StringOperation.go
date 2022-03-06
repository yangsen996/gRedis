package greids

import (
	"context"
	"time"
)

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (s *StringOperation) Set(key string, value interface{}, attr ...*OperationAttr) *InterfaceResult {
	expr := OperationAttrs(attr).FindName("expr")
	nx := OperationAttrs(attr).FindName("nx")
	if nx != nil {
		return NewInterfaceResult(Redis().SetNX(s.ctx, key, value, nx.UnwrapOrDefault(0*time.Second).(time.Duration)).Result())
	}
	return NewInterfaceResult(Redis().Set(s.ctx, key, value, expr.UnwrapOrDefault(0*time.Second).(time.Duration)).Result())
}

func (s *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(s.ctx, key).Result())
}

func (s *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(s.ctx, keys...).Result())
}
