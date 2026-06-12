package work

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET"{
			time.Sleep(4*time.Second)
			fmt.Fprintf(w, "4s pass")
		}else {
			fmt.Fprintf(w, "method not allowed")
		}
	})

	http.ListenAndServe("127.0.0.1:8085", mux)
}

func RunClient(){
	ctx := context.Background()
	ctxWtihT, _ := context.WithTimeout(ctx, 3*time.Second)
	resq, err := http.NewRequestWithContext(ctxWtihT, "GET", "http://127.0.0.1:8085/timeout", nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{
	}
	res, err := client.Do(resq)
	if err != nil {
		log.Fatalln(err)
	} else {
		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	defer res.Body.Close()
}