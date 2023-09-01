package trackskip

import (
	"net/url"
	"strings"
)

// isURL tries to determine if a given string looks like a URL.
func isURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

func ExtractRedirectURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	values := parsedURL.Query()

	for _, valueSlice := range values {
		for _, value := range valueSlice {
			decodedValue, err := url.QueryUnescape(value)

			if err != nil {
				return "", err
			}

			if isURL(decodedValue) {
				return decodedValue, nil
			}
		}
	}

	return "", nil
}
