package work

import "net/http"

func RunServer() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe("127.0.0.1:8085", )
}