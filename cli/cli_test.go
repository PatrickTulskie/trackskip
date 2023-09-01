package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "https://www.example.com/click-100145600-13142092?url=https%3A%2F%2Fgoogle.com",
			output: "https://google.com\n",
		},
		{
			input:  "https://www.example.com/?q=https%3A%2F%2Fgoogle.com",
			output: "https://google.com\n",
		},
	}

	for _, tt := range tests {
		cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcess", "--", tt.input)
		cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
		if !strings.Contains(string(out), tt.output) {
			t.Fatalf("expected output %q, but got %q", tt.output, out)
		}
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}

	// Set the arguments for the main function
	os.Args = append([]string{os.Args[0]}, args...)

	main()
}
