package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	tChan := make(chan int)
	go func() {
		_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		tChan <- 1
	}()
	result := <-tChan
	fmt.Println(result)
}
