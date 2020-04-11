package main

import "fmt"

func sum(s []int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2])
	<-c
	fmt.Println("ssss")
}
