package main

import (
	"fmt"

	"github.com/sabubhatia/quotes"
)

func main() {
	for _, q := range quotes.Favs() {
		fmt.Println(q)
	}
	fmt.Println(quotes.Favs())
}