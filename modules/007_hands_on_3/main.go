package main

import (
	"fmt"

	"rsc.io/quote/v3"
	quoteV2 "rsc.io/quote/v2"
)

func main() {
	fmt.Println(quote.GlassV3())
	fmt.Println(quoteV2.OptV2())

}