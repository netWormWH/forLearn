package learTesting

import "sync"

type ForTest struct{
	mu *sync.Mutex
	dic map[string]int
}
func NewForTest()*ForTest{
	res:=new(ForTest)
	res.dic=make(map[string]int)
	res.mu=new(sync.Mutex)
	return res
}

func (self *ForTest)SetKey(key string ,val int){
	self.mu.Lock()
	defer self.mu.Unlock()
	self.dic[key]=val;
}
func (self *ForTest)Remove(key string){
	self.mu.Lock()
	defer self.mu.Unlock()
	delete(self.dic,key)
}
func (self *ForTest)Has(key string)bool{
	self.mu.Lock()
	defer self.mu.Unlock()
	_,ok:=self.dic[key]
	return ok
}

