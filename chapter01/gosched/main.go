package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func showNumber(num int) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
	time.Sleep(time.Millisecond * 10)

}
func main() {
	runtime.GOMAXPROCS(2)
	iteration := 10

	for i := 0; i <= iteration; i++ {
		go showNumber(i)
	}

	fmt.Println("Goodbye!")
	runtime.Gosched() // Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.

}
