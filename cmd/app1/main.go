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
	cacheChan := make(chan data.Book)
	dbChan := make(chan data.Book)
	for i:=0; i<3; i++ {
		id := rnd.Intn(3) + 1
		wg.Add(3)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- data.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheChan)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- data.Book) {
			if b, ok := queryDb(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbChan)
		go func(cacheCh, dbCh <-chan data.Book, wg *sync.WaitGroup) {
			select{
			case cachedBook := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(cachedBook)
				<-dbCh
			case databaseBook := <-dbCh:
				fmt.Println("from db")
				fmt.Println(databaseBook)
			}
			wg.Done()
		}(cacheChan, dbChan, wg)
		time.Sleep(100 * time.Millisecond)
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
			return b, true
		}
	}
	return data.Book{}, false
}