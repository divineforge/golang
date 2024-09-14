package main

import (
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

func main() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Println("Hello, World!")
		fmt.Println(quote.Go())
		fmt.Println(quote.Go())
		fmt.Println(quote.Go())
		fmt.Println(quote.Go())
		fmt.Println(quote.Glass())
		fmt.Println(quote.Opt())
	}

}
