package http

import (
	"GO-POT/core/state"
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	log.Println("Starting HTTP server on port 8080")

	state := &state.HTTPState{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s with User-Agent: %s and Method: %s", r.RemoteAddr, r.UserAgent(), r.Method)
		if r.URL.Path != "/" {
			log.Printf("[%s] Requested path: %s", r.RemoteAddr, r.URL.Path)
			http.Redirect(w, r, "/", 302)
			return
		}

		if len(r.URL.Query()) > 0 {
			log.Printf("[%s] Query parameters: %s", r.RemoteAddr, r.URL.RawQuery)
		}

		state.RemoteAddr = r.RemoteAddr
		state.UserAgent = r.UserAgent()
		state.Method = r.Method

		fmt.Fprint(w, "It works!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
