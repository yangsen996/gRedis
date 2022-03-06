package greids

import (
	"fmt"
	"time"
)

type empty struct{}

type OperationAttr struct {
	Name  string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (o OperationAttrs) FindName(name string) *InterfaceResult {
	for _, item := range o {
		if item.Name == name {
			return NewInterfaceResult(item.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("found err%s", name))
}

func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{Name: "expr", Value: t}
}

func WithNx() *OperationAttr {
	return &OperationAttr{Name: "nx", Value: empty{}}
}
