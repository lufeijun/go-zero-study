package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := timeout(doBadthing)

	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(2 * time.Second)
}

func doBadthing(done chan bool) {
	fmt.Println("doBadthing start")

	defer func() {
		fmt.Println("doBadthing end")
	}()
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) (err error) {

	fmt.Println("timeout start")

	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
	case <-time.After(time.Millisecond):
		err = errors.New("timeout")
	}

	fmt.Println("timeout end")
	return

}
