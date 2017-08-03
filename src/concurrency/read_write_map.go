package main

import (
	"sync"
	"time"
	"fmt"
	"strconv"
)

func main() {
	mutex := &sync.Mutex{}

	//store := Store{}

	store := make(map[string]int)
	c := make(chan []string)

	store["a"] = 1
	store["b"] = 1
	store["c"] = 1
	store["d"] = 1

	for {
		time.Sleep(time.Duration(5) * time.Second)
		go func() {
			chanData := make([]string, 2)
			chanData[0] = "a"
			chanData[1] = strconv.Itoa(store["a"] +1)
			c <- chanData
		}()


		go func() {
			for k, v := range store {
				if k == "b" {
					mutex.Lock()
					store["b"] = store["b"] + 1
					chanData := <- c
					i, _ := strconv.Atoi(chanData[1])
					store[chanData[0]] = i
					mutex.Unlock()
				}
				fmt.Println("test", k, v)
			}
		}()
		//store.updating = false
	}

}


func updateRequest(st map[string]int, mu *sync.Mutex) {

	//if !st.updating {
	mu.Lock()
	st["a"] = st["a"] + 1
	fmt.Println("updateRequest", st["a"])
	mu.Unlock()
	//}


}

