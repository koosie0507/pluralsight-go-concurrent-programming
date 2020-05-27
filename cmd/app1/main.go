package main

import (
	"fmt"

	"github.com/koosie0507/pluralsight-go-concurrent-programming/data"
)

func main() {
	for _, book := range data.Books {
		fmt.Println(book.String())
	}
}
