package main

import (
	"fmt"

	// "github.com/johndelavega/go-lib"
	golib "github.com/johndelavega/go-lib"
)

func main() {
	fmt.Println("cmd/golib v0.0.1")
	fmt.Println(fmt.Sprintf("golib.FuncTest(): %s\n", golib.FuncTest()))
}
