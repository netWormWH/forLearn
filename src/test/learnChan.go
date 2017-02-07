package main

import (
	"time"
	"fmt"
)

func main() {
	ch:=make(chan int,0)
	
	go func(ch chan int){
		for {
			ch<-0
			time.Sleep(1*time.Second)
		}
	}(ch)
	
	for {
		val,ok:=<-ch
		if ok{
			fmt.Println(val)
		}else{
			break
		}
	}
}
