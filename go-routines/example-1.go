package main
import (
    "fmt"
    // "sync"
    )

func goroutines_printer(cnt int){
    for i:=0; i<cnt; i++ {
        fmt.Printf("Go! : %d\n", i);
    }
}

func normal_printer(cnt int){
    for i:=0; i<cnt; i++ {
        fmt.Printf("Hello: %d\n", i);
    }
}
func main() {
    go goroutines_printer(10);
    normal_printer(10);
    
}

/*
output: 
Hello: 0
Hello: 1
Hello: 2
Hello: 3
Hello: 4

without WGs the main function will complete its execution as soon as normal_printer is completed without actually waiting for goroutines_printer() to complete itself
*/
