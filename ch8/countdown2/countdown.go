package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10*time.Second):
		//Do nothing
	case <-abort:
		return
	}
	fmt.Println("launch!!!")
}