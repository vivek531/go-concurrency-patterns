package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	outChannel := make(chan int, 10)
	go func() {
		for _, n := range nums {
			outChannel <- n
		}
		close(outChannel)
	}()
	return outChannel
}

func sq(in <-chan int) <-chan int {
	outChannel := make(chan int, 10)
	go func() {
		for n := range in {
			outChannel <- n * n
		}
		close(outChannel)
	}()
	return outChannel
}

func main() {
	nums := []int{2, 3, 4, 5}
	inputChannel := sliceToChannel(nums)
	finalChannel := sq(inputChannel)
	for n := range finalChannel {
		fmt.Println(n)
	}
}
