package goroutine

import (
	"runtime"
	"fmt"
)

func say(s string){
	for i := 0 ; i < 10; i++{
		runtime.Gosched() //使线程休息一下。再进行处理
		fmt.Println(s)
	}
}