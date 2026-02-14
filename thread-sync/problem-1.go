// two threads
// thread 1 => 1, 2, 3, 4, 5 ....
// thread 2 => A, B, C, D, ....
// overall => A 1 B 2 C 3 .... Z 26
// doubts clarification: https://gemini.google.com/share/afbd6eea1e44

package main

import (
	"fmt"
	"sync"
)



func printChars(wg *sync.WaitGroup, numChannel, charChannel chan bool) {
	defer wg.Done()
	for i := 'A'; i <= 'Z'; i++ {
		<-charChannel
		fmt.Printf("%c ", i)
		numChannel <- true

		// Only signal if we aren't at the end of the alphabet
		// if i < 'Z' {
		// 	numChannel <- true
		// }
	}
}

func printNums(wg *sync.WaitGroup, numChannel, charChannel chan bool) {
	defer wg.Done()
	for i := 1; i <= 26; i++ {
		<-numChannel
		fmt.Printf("%d ", i)

		// Only signal if we aren't at i=26 (last integer)
		if i < 26 {
			charChannel <- true
		}
	}

}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	numChannel := make(chan bool)
	charChannel := make(chan bool)

	go printChars(&wg, numChannel, charChannel)
	go printNums(&wg, numChannel, charChannel)

	charChannel <- true

	wg.Wait()

	close(numChannel)
	close(charChannel)

}
