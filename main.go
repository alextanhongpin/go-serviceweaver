package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/metrics"
)

var reverseCounter = metrics.NewCounter(
	"reverse_count",
	"The number of times Reverser.Reverse has been called",
)

//go:generate weaver generate .
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

	// Get a client to the Reverser component.
	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "Hello %s!\n", reversed)
	})
	http.Serve(lis, nil)
}

type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

type reverserOptions struct {
	Name string
}

type reverser struct {
	weaver.Implements[Reverser]
	weaver.WithConfig[reverserOptions]
}

// Init is where you setup infrastructure like storage etc.
func (r *reverser) Init(ctx context.Context) error {
	name := r.Config().Name
	logger := r.Logger()
	nameLogger := logger.With("name", name)
	nameLogger.Info("initialized with config")

	return nil
}

func (r *reverser) Reverse(_ context.Context, s string) (string, error) {
	reverseCounter.Add(1)
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes), nil
}
