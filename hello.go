package main

import (
	"fmt"

	"github.com/antonioeatgoat/go-hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf(morestrings.ReverseRunes("Hello, world"))
	fmt.Printf(cmp.Diff("Hello World", "Hello Go"))
}
