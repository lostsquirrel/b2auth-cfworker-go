package main

import (
	"fmt"
	"testing"
)

func TestListKVNS(t *testing.T) {
	cfg := BuildFromEnv()
	kvns := cfg.ListKVNS()
	fmt.Println(kvns)
}
