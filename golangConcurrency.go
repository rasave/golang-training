package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() { // goroutine for alphabet

		for ch := 'A'; ch < 'Z'; ch += 2 {
			letter <- true
			fmt.Print(string(ch))
			fmt.Print(string(ch + 1))
			<-number

		}
		close(letter)
	}()

	go func() {
		start := 1
		for {
			number <- true

			fmt.Print(start)
			fmt.Print(start + 1)
			start += 2

			_, ok := <-letter
			if ok == false {
				break
			}
		}
		wg.Done()
	}()
	////////// inside the main()

	<-number

	wg.Wait()
	fmt.Print("\n")

}

