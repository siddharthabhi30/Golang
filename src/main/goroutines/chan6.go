package main

import (
	"fmt"
	"time"
)

func main() {
	// jobs := make(chan int, 100)
	// results := make(chan int, 100)

	// //multicore processing ...is what we can see here

	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	// for i := 0; i < 10; i++ {
	// 	jobs <- i
	// }
	// //close(jobs)

	// //it should be 10 ..else deadlock...we havenot closed this channel
	// for j := 0; j < 10; j++ {
	// 	fmt.Println(<-results)
	// }

	//one()
	three()
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- n
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func one() {
	job := make(chan int)

	go two(job)
	job <- 10
	close(job)

	//think like ..main thread is waiting and as we all know..main thread after receiving terminates
	//after main thread terminates....all the go workers are jobless
	time.Sleep(time.Millisecond * 3000)

	//close(job)
	//range always searches...which is bad ..if the channel is not closed

}

func two(job chan int) {
	// for msh := range job {
	//
	// }
	//time.Sleep(time.Millisecond * 3000)
	//fmt.Println("sss")
	msh := <-job
	fmt.Println(msh)
}

//we know receiver has to wait.....demo  receiver really deadlocks if it cant find value
//receiver should not look for value even in buffeted channel ..if the value is not there..or no possiblity to find the value..or
//if receiver exactly knows that there are X values in the channel.......
//main thread should be mostly rewceive or blocking..it cant have non blocking ..else all
//thread will lose their job
//you can concurrently send value to channel and receiver getting it,,,but at some point we should close the
//channel from sender POV
//receiver may wait for channel to send value even if it is long ...but channel's sender end shouldn't be open
//for infinite time..eg three()..main fucn is waiting for 10 second ..not for infinite

// receiving from close channel results in zero
// receiving from open buffet channel ...deadlock
// if the receiver knows ....length of the channel..then there is no problem. even if the channel is not closed
// compiler wait if some thread maay send something to this channel......if one thread is in sleep and
// there is a possibility that it may send some data ..then compiler waits before sending deadlock
// above thing is may a bug
// wait group in chan7......it is like delaying the main method ....for all
// rountined to finish
// if a function running on a thread is calling two threads....and if func ends ..thread ends....
// the parent func exibits main thread behaviour
func three() {
	chh := make(chan int)

	go func() {
		close(chh)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			chh <- i
		}

		//close(chh)
	}()

	go func() {
		time.Sleep(time.Millisecond * 10000)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			//chh <- i
		}
	}()

	go func() {
		time.Sleep(time.Millisecond * 100000)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			//chh <- i
		}
	}()
	// for n := range chh {
	// 	fmt.Println(n)
	// }

	for i := 0; i < 12; i++ {
		fmt.Println(<-chh)
	}

}
