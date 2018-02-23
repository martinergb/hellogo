package main

import (
	"fmt"
	"runtime"

	"github.com/rickb777/date"
)

func main() {
	defer func() {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}()

	d := date.Max()
	fmt.Println(d)
}
