package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqPath := r.URL.Path
		switch reqPath {
		case "/":
			fmt.Fprintf(w, "Hello GoLang!")
		case "/health":
			var health bool
			health = checkHealth()
			if health {
				fmt.Fprintf(w, "Healthy")
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Error - Site is in unhealthy state!"))
			}
		case "/make-healthy":
			res := createHealthyFile()
			if res {
				fmt.Fprintf(w, "Done!")
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Error - Set healthy state failed"))
			}
		case "/make-unhealthy":
			res := removeHealthyFile()
			if res {
				fmt.Fprintf(w, "Done!")
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Set healthy state failed"))
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Page not found!"))
		}
	})
	http.ListenAndServe(":80", nil)
}

func removeHealthyFile() bool {
	var res bool
	err := os.Remove("healthy_file")
	if err == nil {
		res = true
	}
	return res
}

func createHealthyFile() bool {
	var res bool
	_, err := os.Create("healthy_file")
	if err == nil {
		res = true
	}
	return res
}

func checkHealth() bool {
	if _, err := os.Stat("healthy_file"); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
