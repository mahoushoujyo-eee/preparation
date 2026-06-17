package service

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func CheckClusterServiceHealth() {

}

func RequestWithTimeout(url string, method string) {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	resq, err := http.NewRequestWithContext(ctxWithTimeout, method, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{}
	res, err := client.Do(resq)
	if err != nil {
		log.Fatalln(err)
	} else {
		body, _ := io.ReadAll(res.Body)
		log.Println(string(body))
	}
	defer res.Body.Close()
}