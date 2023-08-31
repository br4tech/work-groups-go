// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int) {
// 	fmt.Printf("Processo %d iniciado\n", id)

// 	time.Sleep(time.Second)
// 	fmt.Printf("Processo %d finalizado\n", id)
// }

// func main() {

// 	var wg sync.WaitGroup

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)

// 		i := i

// 		go func() {
// 			defer wg.Done()
// 			worker(i)
// 		}()
// 	}

// 	wg.Wait()

// }

//////////////////////////////////////////////////////////////////////////////////

// Channels

// package main

// import "fmt"

// func main() {

// 	messages := make(chan string)

// 	go func() { messages <- "ping" }()

// 	msg := <-messages
// 	fmt.Println(msg)
// }

/////////////////////////////////////////////////////////////////////////////////////////

// Mutex

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Container struct {
// 	mu       sync.Mutex
// 	counters map[string]int
// }

// func (c *Container) inc(name string) {

// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.counters[name]++
// }

// func main() {
// 	c := Container{

// 		counters: map[string]int{"a": 0, "b": 0},
// 	}

// 	var wg sync.WaitGroup

// 	doIncrement := func(name string, n int) {
// 		for i := 0; i < n; i++ {
// 			c.inc(name)
// 		}
// 		wg.Done()
// 	}

// 	wg.Add(3)
// 	go doIncrement("a", 10000)
// 	go doIncrement("a", 10000)
// 	go doIncrement("b", 10000)

// 	wg.Wait()
// 	fmt.Println(c.counters)
// }
