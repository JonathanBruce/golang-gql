package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jonathanbruce/graphql/resolver"
)

func handler(writer http.ResponseWriter, r *http.Request) {

}

func main() {
	var (
		addr              = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	root, err := resolver.NewRoot(c)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
}
