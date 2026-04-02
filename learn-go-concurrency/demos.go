package main

import (
	"fmt"
	"sync"
	"time"
)

func Num01GoroutineDemo() {

	go func() {
		fmt.Println("hello world from goroutine")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("main done")
}

func Num02ChannelsBasicDemo() {
	ch := make(chan string)

	go func() {
		ch <- "data from goroutine pipeline"
	}()

	msg := <-ch
	fmt.Println("received:", msg)

}

func Num03ChannelsBufferedDemo() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func Num04WaitGroupDemo() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Println("worker", id, "done")
		}(i)
	}

	wg.Wait()
	fmt.Println("all worker one after wg wait")
}

func Num05MutexDemo() {
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()

			counter++

			fmt.Println("counter++", counter)
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("final counter:", counter)
}

func Num06SelectDemo() {
}

func Num07ChannelPatternsDemo() {
}

func Num08ContextDemo() {
}

func Num09RWMutexDemo() {
}

func Num10SyncOnceDemo() {
}

func Num11SyncMapDemo() {
}

func Num12AtomicDemo() {
}

func Num13WorkerPoolDemo() {
}

func Num14ErrGroupDemo() {
}

func Num15MemoryModelDemo() {
}

func Num16RaceDetectorDemo() {
}

func Num17SemaphoreDemo() {
}

func Num18GMPSchedulerDemo() {
}

func Num19BenchmarkingDemo() {
}
