package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

func main() {
	if err := os.Setenv("GOPROXY", "https://goproxy.cn"); err != nil {
		log.Fatalf("failed to set $GOPROXY: %v", err)
	}

	println(os.Getenv("GONOPROXY"))

	http.ListenAndServe(":8080", goproxy.New())
}
