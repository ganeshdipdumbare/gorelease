package gocode

type GoCode struct {
	Code string
}

// NewGoCode creates a new GoCode instance with the given code
// and returns a pointer to it.
func NewGoCode(code string) *GoCode {
	return &GoCode{
		Code: code,
	}
}

type GoCodePlayOutput struct {
	Output string
	Error  string
}
