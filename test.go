package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// fmt.Println(reverseString("hello"))

	jobs := make(chan int, 5)

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go Worker(i, jobs)
	}
	for j := 0; j <= 10; j++ {
		jobs <- j
	}

	go func() {
		defer close(jobs)
		defer wg.Wait()
	}()

}

func reverseString(s string) string {
	newval := []rune(s)

	i := 0
	j := len(newval) - 1
	for i < j {
		newval[i], newval[j] = newval[j], newval[i]
		i++
		j--
	}
	return string(newval)
}

func Worker(id int, jobs <-chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Worker ", id, "need to do ", job)
	}
}
