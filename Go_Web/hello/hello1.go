package main

import (
	"fmt"

	"rsc.io/quote" // import external hosted repo - go mod tidy
)

func hello() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Go()) // use dependency funtion
}
