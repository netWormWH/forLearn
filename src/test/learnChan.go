package main

import (
	"time"
	"fmt"
)

func main() {
	ch:=make(chan int,0)
	
	go func(ch chan int){
		for i:=0;i<10;i++{
			ch<-i
			time.Sleep(1*time.Second)
		}
		close(ch) // close the channel
	}(ch)
	
	for {
		val,ok:=<-ch
		if ok{
			fmt.Println(val)
		}else{
			fmt.Println("channle is closed!")
			break
		}
	}
	fmt.Println("Program ready Exit!")
}
