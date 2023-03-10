package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	// Initializes the Service Weaver application. It also returns a
	// weaver.Instance, which we assign to root.
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}

	// Returns a network listener, similar to net.Listen.
	// With Service Weaver, listeners are named.
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hello listener available on %v\n", lis)

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!\n", r.URL.Query().Get("name"))
	})
	http.Serve(lis, nil)
}
