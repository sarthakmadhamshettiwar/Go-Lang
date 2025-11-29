package main
import "fmt"

var sum_with_goroutines int = 0;
func increment_with_goroutines(){
    sum_with_goroutines++;
}

var sum_without_goroutines int = 0;
func increment_without_goroutines(){
    sum_without_goroutines++;
}

func main() {
    for i:=0; i<1000; i++ {
        go increment_with_goroutines();
        increment_without_goroutines();
    }
    fmt.Printf("Sum without goroutines: %d\n", sum_without_goroutines);
    fmt.Printf("Sum with goroutines: %d\n", sum_with_goroutines);
}


/*
output:
Sum without goroutines: 1000
Sum with goroutines: 980
*/
