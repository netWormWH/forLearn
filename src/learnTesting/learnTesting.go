package learTesting


type ForTest struct{
	
	dic map[string]int
}
func NewForTest()*ForTest{
	res:=new(ForTest)
	res.dic=make(map[string]int)
	return res
}

func (self *ForTest)SetKey(key string ,val int){
	self.dic[key]=val;
}
func (self *ForTest)Remove(key string){
	delete(self.dic,key)
}
func (self *ForTest)Has(key string)bool{
	_,ok:=self.dic[key]
	return ok
}

