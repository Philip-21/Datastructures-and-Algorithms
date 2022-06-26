//additional examples on working with channels and goroutines

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//channels transfer data between different go routines
//channels provide a means for diffent goroutines to communicate with each other

func main() {
	//creating the channel
	//(1) //unbuffered channel
	dataChan := make(chan int)
	go func() {
		dataChan <- 3387 //sending data into the channel
	}()
	n := <-dataChan          //receiving data from the channel
	fmt.Printf("n= %d\n", n) //prints 3387

	//(2) buffered channel
	nextChan := make(chan int, 2)
	nextChan <- 129
	nextChan <- 354
	v := <-nextChan
	fmt.Printf("v=%d\v", v) //prints 129345

	//(3)
	//using for loops in channels
	makeChan := make(chan int)
	go func() { //this runs on a background thread
		for i := 0; i < 1000; i++ {
			makeChan <- i //sending data into the channel
		}
		close(makeChan) //the go routine terminates well without a deadluck
	}()
	//receiving data from the channel
	for n := range makeChan { //this runs on the main thread
		fmt.Printf("n= %d\n", n)
	}

	//(4)
	//working with wait grups
	//wait groups wait for a collection of goroutines to finish
	//using wait groups in channels
	newChan := make(chan int)
	go func() {
		wg := sync.WaitGroup{}    //using a wait group on each iteration of the loop(each go routine per iteration )
		for i := 0; i < 50; i++ { //for each iteration it creates a new goroutine
			wg.Add(1)
			//initializing another goroutine to make it faster
			go func() {
				defer wg.Done()
				result := DoWork() //a random integer of 100
				newChan <- result  //sending result into the channel
			}()
			wg.Wait() //wait until all the iterations are done
			close(newChan)
		}
	}()

}

//function for example 4
func DoWork() int {
	time.Sleep(time.Second) //this stops the  execution of the goroutine for a second the returns the integer which will be used in the goroutine
	return rand.Intn(100)
}
