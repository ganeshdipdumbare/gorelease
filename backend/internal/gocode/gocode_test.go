package gocode

import (
	"net/url"
	"testing"
)

func TestNewGoCode(t *testing.T) {
	code := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"
	expectedURLEncoded := url.QueryEscape(code)

	goCode := NewGoCode(code)

	if goCode.Code != expectedURLEncoded {
		t.Errorf("Expected URLEncoded to be %s, but got %s", expectedURLEncoded, goCode.Code)
	}

	multilineCode := `// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
`

	goCode = NewGoCode(multilineCode)
	expectedURLEncoded = url.QueryEscape(multilineCode)
	if goCode.Code != expectedURLEncoded {
		t.Errorf("Expected URLEncoded to be %s, but got %s", expectedURLEncoded, goCode.Code)
	}
}
