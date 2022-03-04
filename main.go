package main

import (
	"fmt"
	"sync"
)

func corruptedData() {
	wg := sync.WaitGroup{}
	N := 10
	arr := []int{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			arr = append(arr, i)
			wg.Done()
		}()
	}

	wg.Wait() // wait until all routines are done

	fmt.Println(arr) // [10, 10, 10...]
}

func correctDataShadowing() {
	wg := sync.WaitGroup{}
	N := 10
	arr := []int{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		i := i // add this line here to shadow variable

		go func() {
			arr = append(arr, i)
			wg.Done()
		}()
	}

	wg.Wait() // wait until all routines are done

	fmt.Println(arr) // array with 0-9 values in random order
}

func correctDataArg() {
	wg := sync.WaitGroup{}
	N := 10
	arr := []int{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func(i int) { // add argument here
			arr = append(arr, i)
			wg.Done()
		}(i) // and here
	}

	wg.Wait() // wait until all routines are done

	fmt.Println(arr) // array with 0-9 values in random order
}

func main() {
	corruptedData()
	correctDataShadowing()
	correctDataArg()
}
