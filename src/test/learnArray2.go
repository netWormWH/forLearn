package main

import "fmt"

var arr=[...]string{
	12:"3",
	1:"a",
	10:"b",
}

func main() {
	
	fmt.Println(len(arr))
	for i,v:=range  arr{
		fmt.Println(i,v)
	}
	
}
