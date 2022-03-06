package greids

type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func (i *InterfaceResult) Unwrap() interface{} {
	if i.Err != nil {
		panic(i.Err)
	}
	return i.Result
}

func (i *InterfaceResult) UnwrapOrDefault(defaultValue interface{}) interface{} {
	if i.Err != nil {
		return defaultValue
	}
	return i.Result
}
