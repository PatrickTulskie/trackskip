package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PatrickTulskie/trackskip/trackskip"
)

func main() {
	var inputURL string

	// Check if there's a command line argument.
	if len(os.Args) > 1 {
		inputURL = os.Args[1]
	} else {
		// Read from standard input.
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			inputURL = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			os.Exit(1)
		}
	}

	if inputURL == "" {
		fmt.Println("No URL provided.")
		os.Exit(1)
	}

	redirectURL, err := trackskip.ExtractRedirectURL(inputURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if strings.TrimSpace(redirectURL) == "" {
		fmt.Println("No redirect URL found.")
		os.Exit(1)
	}

	fmt.Println(redirectURL)
}
