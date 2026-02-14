/*
The number passed to Add() must exactly match the number of times Done() is executed.
If Add > Done: Deadlock (as in your example).
If Done > Add: The program will panic immediately with a "negative WaitGroup counter" error.

If we just do wg.Add(2) in main() no error will be thrown.
*/

package main

import (
	"fmt"
	"sync"
)

// two threads
// thread 1 => 1, 2, 3, 4, 5 ....
// thread 2 => A, B, C, D, ....
// overall => 1, A, 2, B, 3, C, ...

func printChars(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c ", i)
	}
}

func printNums(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 26; i++ {
		fmt.Printf("%d ", i)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go printChars(&wg)
	go printNums(&wg)
	wg.Wait()
}

// Output: 1 2 3 4 5 6 7 8 9 10 A B C D E F G H 11 12 13 14 15 16 17 18 19 20 21  I J K L M N O P Q R S T U V W 22 23 24 25 26 X Y Z fatal error: all goroutines are asleep - deadlock!
