package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// create a wait group
var wg sync.WaitGroup

func printRandom(i int) {
	defer wg.Done()
	r := rand.Intn(20000)
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println(i)

}

func main() {
	noofthreads := 10
	wg.Add(noofthreads)
	for i := 0; i < noofthreads; i++ {
		//create a thread
		go printRandom(i)
	}
	//wait for all the threads to finish
	wg.Wait()
}
