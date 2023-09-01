package trackskip

import (
	"testing"
)

func TestExtractRedirectURL(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "https://www.example.com/?click-100145600-13142092?url=https%3A%2F%2Fgoogle.com",
			output: "https://google.com",
		},
		{
			input:  "https://www.example.com/?q=https%3A%2F%2Fgoogle.com",
			output: "https://google.com",
		},
		{
			input:  "https://www.example.com/?q=845999338855553234",
			output: "",
		},
	}

	for _, tt := range tests {
		result, _ := ExtractRedirectURL(tt.input)
		if result != tt.output {
			t.Errorf("For input %q, expected %q, got %q", tt.input, tt.output, result)
		}
	}
}
