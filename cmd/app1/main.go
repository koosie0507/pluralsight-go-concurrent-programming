package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/koosie0507/pluralsight-go-concurrent-programming/data"
)

var cache = map[int]data.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
func main() {
	for i:=0; i<10; i++ {
		id := rnd.Intn(3) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDb(id); ok {
			fmt.Println("from db")
			fmt.Println(b)
			continue
		}
		fmt.Println("book ID ", id, " not found")
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (data.Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDb(id int) (data.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range data.Books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return data.Book{}, false
}