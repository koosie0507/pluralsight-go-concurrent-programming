package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/koosie0507/pluralsight-go-concurrent-programming/data"
)

var cache = map[int]data.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i:=0; i<10; i++ {
		id := rnd.Intn(3) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDb(id, m); ok {
				fmt.Println("from db")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (data.Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDb(id int, m *sync.RWMutex) (data.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range data.Books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return data.Book{}, false
}