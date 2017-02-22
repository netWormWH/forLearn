package main

import (
	"fmt"
	"encoding/gob"
	"bytes"
	"mf/utils"
	"encoding/json"
)

type SA struct {
	A int
	B string
	Arr []int
}





func main(){
	//tmp:=new(SA)
	//tmp.A=111
	//tmp.B="kdjf;lakjfkljasfkljasfkljadfjasf;ljasfljasdfjadfljas;lfkjasfkljasfklj"
	//tmp.Arr=make([]int,10)
	//tmp.Arr[9]=1
	//fmt.Println(len(tmp.Arr))
	//
	tmp:=map[string]interface{}{"F":"foo","A":[]interface{}{1,"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}}
	
	utils.CheckBegin()
	
	b1,e1:=testGob(tmp)
	utils.CheckEnd()
	fmt.Println(len(b1),e1)
	utils.CheckBegin()
	b2,e2:=testJson(tmp)
	
	utils.CheckEnd()
	
	fmt.Println(len(b2),e2)
}

func testGob(t interface{})(b []byte,err error){
	for i:=0;i<1000;i++ {
		var buff bytes.Buffer
		enc := gob.NewEncoder(&buff)
		err = enc.Encode(t)
		b = buff.Bytes()
	}
	return
}
func testJson(t interface{})(b []byte,err error){
	for i:=0;i<1000;i++ {
		b, err = json.Marshal(t)
	}
	return
}
