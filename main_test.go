package main

import (
	"context"
	"testing"

	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/weavertest"
)

func TestReverse(t *testing.T) {

	ctx := context.Background()
	root := weavertest.Init(ctx, t, weavertest.Options{})

	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		t.Fatal(err)
	}

	got, err := reverser.Reverse(ctx, "hello")
	if err != nil {
		t.Fatal(err)
	}

	if exp := "olleh"; exp != got {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
