package main

import (
	"time"
)

func main() {
	defer println("in main")
	go func() {
		defer func() {
			println("in goroutine")
			if err := recover(); err != nil {
				println(err)
				panic("")
			}
		}()
		panic("")
	}()

	time.Sleep(1 * time.Second)
}
