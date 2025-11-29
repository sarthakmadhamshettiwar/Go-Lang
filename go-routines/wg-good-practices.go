package main
import (
    "fmt"
    "sync"
    )

func goroutines_printer(cnt int, wg *sync.WaitGroup){
    for i:=0; i<cnt; i++ {
        func (){
            wg.Add(1);
            defer wg.Done();
            fmt.Printf("Go! : %d\n", i);
        }()
    }
}

func normal_printer(cnt int){
    for i:=0; i<cnt; i++ {
        fmt.Printf("Hello: %d\n", i);
    }
}
func main() {
    var wg sync.WaitGroup;
    go goroutines_printer(5, &wg);
    normal_printer(5);
    
    wg.Wait();
    
}

// there are two possible outputs for this case: 
// case 1: 
/*
Hello: 0
Hello: 1
Hello: 2
Hello: 3
Hello: 4

Q. Why this happened even when the WGs were used? 
A: the goroutine was made when goroutines_printer() was called, lets say it took 10ms to spin it up. 
In those 10ms the normal_printer() completed its entire execution.
Now since goroutines_printer goroutine was yet to be spunned up the WG was empty!
And the control reached line 27 with empty WG, and thus it didn't waited for goroutines_printer

Best Practice: Always call wg.Add() in the parent thread (outside the goroutine) BEFORE you spawn the goroutine.
*/

// case 2: 
/*
in the above case what if it took more than 10ms to finish the execution of normal_printer? 
Output and error like this: 
Hello: 0
Hello: 1
Hello: 2
Hello: 3
Hello: 4
Go! : 0
Go! : 1
Go! : 2
Go! : 3
Go! : 4
panic: sync: WaitGroup is reused before previous Wait has returned

goroutine 1 [running]:
sync.(*WaitGroup).Wait(0x5?)
	/usr/local/go/src/sync/waitgroup.go:120 +0x74
main.main()
	/tmp/FuAlkAgMpj/main.go:30 +0x6f
exit status 2


The Crash Sequence
Here is the race condition that causes the panic:

Loop 1: wg.Add(1) (Count is 1). Code prints. wg.Done() (Count is 0).

Main Thread: wg.Wait() sees the count is 0. It thinks, "Great! All work is finished. I will release the block and exit."

Loop 2: At the exact same moment, your loop continues and calls wg.Add(1) again.

The Conflict: Wait() is trying to return because "Work is done (0)", but Add() is screaming "Wait, I have more work!". Go panics to prevent unpredictable behavior.
*/
