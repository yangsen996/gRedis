package greids

// Iterator 迭代器
type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

func (t *Iterator) HasNext() bool {
	if t.data == nil || len(t.data) == 0 {
		return false
	}
	return t.index < len(t.data)
}

func (t *Iterator) Next() (ret interface{}) {
	ret = t.data[t.index]
	t.index += 1
	return
}
