package gocodeplayer

import (
	"testing"

	"backend/internal/gocode"
)

func TestNewPlaygroundClient(t *testing.T) {
	url := "https://example.com"
	client, err := NewPlaygroundClient(url)
	if err != nil {
		t.Errorf("Failed to create PlaygroundClient: %v", err)
	}

	if client.url != url {
		t.Errorf("Expected URL to be %s, but got %s", url, client.url)
	}

	if client.logger == nil {
		t.Error("Expected logger to be initialized")
	}

	if client.reqClient == nil {
		t.Error("Expected reqClient to be initialized")
	}
}

func TestPlaygroundClient_Play(t *testing.T) {
	tests := []struct {
		name           string
		code           *gocode.GoCode
		expectedOutput string
		expectedError  string
	}{
		{
			name: "Valid code",
			code: &gocode.GoCode{
				Code: `// Our first program will print the classic "hello world"
	// message. Here's the full source code.
	package main

	import "fmt"

	func main() {
		fmt.Println("Hello, World!")
	}`,
			},
			expectedOutput: "Hello, World!\n",
			expectedError:  "",
		},
		{
			name: "Empty code",
			code: &gocode.GoCode{
				Code: "",
			},
			expectedOutput: "",
			expectedError:  "code cannot be empty",
		},
		{
			name:           "Nil code",
			code:           nil,
			expectedOutput: "",
			expectedError:  "code cannot be nil",
		},
		{
			name: "Code with compile errors",
			code: &gocode.GoCode{
				Code: `package main

	import "fmt"

	func main() {
		fmt.Println("Hello, World!")
		fmt.Println(a)
	}`,
			},
			expectedOutput: "",
			expectedError:  "./prog.go:7:15  undefined: a\n",
		},
	}

	client, err := NewPlaygroundClient("https://go.dev")
	if err != nil {
		t.Errorf("Failed to create PlaygroundClient: %v", err)
		return
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := client.Play(test.code)
			if err != nil {
				if test.expectedError == "" {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != test.expectedError {
					t.Errorf("Expected error to be %q, but got %q", test.expectedError, err.Error())
				}
				return
			}

			if output.Output != test.expectedOutput {
				t.Errorf("Expected output to be %q, but got %q", test.expectedOutput, output.Output)
			}

			if output.Error != test.expectedError {
				t.Errorf("Expected error to be %q, but got %q", test.expectedError, output.Error)
			}
		})
	}
}
