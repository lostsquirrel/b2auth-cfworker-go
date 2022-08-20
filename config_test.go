package main

import (
	"fmt"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	os.Setenv("CF_WORKER_API_TOKEN", "a")
	os.Setenv("CF_ACCOUNT", "b")
	c := BuildFromEnv()

	fmt.Println(c)
	// for y := range reflect.VisibleFields(x) {
	// 	fmt.Println(y)
	// }
}
