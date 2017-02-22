package main

import (
	c "github.com/mattn/anko/builtins"
	"github.com/mattn/anko/vm"
	"time"
	"fmt"
)

type A struct {
	B int
}

func main() {
	//tmp:="aaa23232aaa"
	
	s:=`
	var time=import("time")
	func foo(a,b){
		println("foo",a,b)
		return a|b
	}
	func foo2(s){
		i=0
		
		for {
			i++
			println(s,i)
			time.Sleep(toDuration(1*time.Second))
		}
	}
	println(foo(2,4))
	#go foo2("A:")
	#go foo2("B:")
	go foo(1,2)
	
	func foo3(){
		t=time.Now().UnixNano()
		n=10
		for i=0;i<10000000;i++{
			#n+=i
		}
		println((time.Now().UnixNano()-t)/time.Millisecond,n)
	}
	foo3()
	#time.Sleep(30*time.Second)
	#println("over",time.Second)
	`

	
	ttt:=new(A)
	ttt.B=111
	
	foo3:=func(){
		t:=time.Now().UnixNano()
		n:=0
		for i:=0;i<10000000;i++{
			n+=i
		}
		println(float64(time.Now().UnixNano()-t)/float64(time.Millisecond),n)
	}
	foo3()
	env:=vm.NewEnv()
	c.LoadAllBuiltins(env)
	
	//env.Define("time",time)
	go func() {
		res, err := env.Execute(s)
		fmt.Println(res, err)
	}()
	//time.Millisecond
	time.Sleep(30*time.Second)
}
