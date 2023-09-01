package main

import (
	"fmt"
	"net/http"

	"github.com/PatrickTulskie/trackskip/trackskip"
)

func handler(w http.ResponseWriter, r *http.Request) {
	redirectURL, err := trackskip.ExtractRedirectURL(r.URL.String())
	if err != nil {
		http.Error(w, "Error extracting redirect URL", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
