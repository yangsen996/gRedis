package greids

type Result interface {
	Unwrap()
	UnwrapOrDefault()
}

type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{
		Result: result,
		Err:    err,
	}
}

func (s *StringResult) Unwrap() string {
	if s.Err != nil {
		panic(s.Err)
	}
	return s.Result
}

func (s *StringResult) UnwrapOrDefault(defaultValue string) string {
	if s.Err != nil {
		return defaultValue
	}
	return s.Result
}

func (s *StringResult) UnwrapOrElse(f func() string) string {
	if s.Err != nil {
		return f()
	}
	return s.Result
}
