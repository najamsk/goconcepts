package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Task struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var t Task
	started := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(100)
	sem := make(chan bool, 10)
	for i := 0; i < 100; i++ {
		// i := i // this will make i captuerd inside go-routine for each loop iteration
		// fmt.Println(runtime.NumGoroutine())
		sem <- true
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-sem
			}()
			res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i))
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
				log.Fatal(err)
			}
			fmt.Println(t.Title)
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(started))
}
