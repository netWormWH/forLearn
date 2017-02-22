package main

import (
	

	"fmt"

	"math/rand"
	"time"
)
var stringVolumes=[]string{"err","err","err","3","4","5","6","7","8","9","10","J","Q","K","A","2"}
var CARDS=[]Card{
	0x13,0x14,0x15,0x16,0x17,0x18,0x19,0x1a,0x1b,0x1c,0x1d,0x1e,0x1f,
	0x23,0x24,0x25,0x26,0x27,0x28,0x29,0x2a,0x2b,0x2c,0x2d,0x2e,0x2f,
	0x33,0x34,0x35,0x36,0x37,0x38,0x39,0x3a,0x3b,0x3c,0x3d,0x3e,0x3f,
	0x43,0x44,0x45,0x46,0x47,0x48,0x49,0x4a,0x4b,0x4c,0x4d,0x4e,0x4f,
	0x51,0x62}
var isSeedSet=false
func ShuffleCards(arr *[]Card)(err error){
	if  !isSeedSet{
		isSeedSet=true
		rand.Seed(time.Now().UnixNano())
	}
	l:=len(*arr)
	if l==0{
		err=fmt.Errorf("arr length=0")
	}else if l>1{
		for i,_:= range(*arr){
			j:=rand.Intn(l)
			(*arr)[i],(*arr)[j]=(*arr)[j],(*arr)[i]
		}
	}
	return
}
func main(){
	c:=make([]Card,len(CARDS))
	copy(c,CARDS)
	//fmt.Println(c)
	ShuffleCards(&c)
	//fmt.Println(c)
	
	arr:=MakeCheckArr(c[:20])
	fmt.Println(arr)
	res:=GetSS(&arr)
	fmt.Println(arr)
	fmt.Println(res)
	//res:=GetRepeat(&arr,4)
	//fmt.Println(arr)
	//res1:=GetRepeat(&arr,3)
	//fmt.Println(arr)
	//res2:=GetRepeat(&arr,2)
	//fmt.Println(arr)
	//fmt.Println(res,res1,res2)
	//
}


type Card int

func (self Card)Color()int{
	//1-4 (5 6)
	return (int(self)>>4)&0x0f
}
func (self Card)Volume()int{
	//2-14 (0 0)
	return int(self)&0x0f
}
func (self Card)PrintVolume()string{
	v:=self.Volume()
	if v>=0&&v<len(stringVolumes){
		return stringVolumes[v]
	}
	return "ErrVolume"
}
func (self Card)String()(s string){
	//♡ ♢ ♤ ♧ ♣ ♦ ♥ ♠ ♕♛ ⓈⒷ
	switch self.Color() {
	case 1:
		s="♢"+self.PrintVolume()
	case 2:
		s="♧"+self.PrintVolume()
	case 3:
		s="♡"+self.PrintVolume()
	case 4:
		s="♤"+self.PrintVolume()
	case 5:
		s="Ⓢ"
	case 6:
		s="Ⓑ"
	}
	return
}

type Cards struct {
	all []Card
	T []int  //[typeCode]   //typeCode=0x|t|num|c|v|    //t 1.one 2.two 3.three 4.sequence 5.bomb(four) 6.bigBomb
	
	//typeCode new  0x t,num,v
}
type Data struct {
	C [][]int
}




func MakeCheckArr(cards []Card)(res []int) {
	res = make([]int, 16)
	for _, c := range cards {
		res[c.Volume()]++
	}
	return
}

func GetTypeCode(t,num,val int)int{
	return t*10000+num*100+val
}
//-------

func GetRepeat(arr *[]int,num int)(res []int){
	res=make([]int,0)
	for c,n:=range *arr{
		if n>=num{
			//res=append(res,(num<<8)|(num<<4)|c)
			if num==4{
				res=append(res,GetTypeCode(num+1,num,c))
			}else{
				res=append(res,GetTypeCode(num,num,c))
			}
			
			
			(*arr)[c]-=num
		}
	}
	return
}
//out:: from best choice


//follow:: find all (current cardType first 's) types ,compare best choice,Preferred result.




func GetSS(arr *[]int)(res []int){
	res=[]int{}
	if (*arr)[7]==0 &&(*arr)[10]==0{
		return
	}
	fmt.Println("getSS::",*arr)
	num:=0
	tmpCard:=0
	for i:=3;i<15;i++{
		if (*arr)[i]>0{
			tmpCard=i
			num++
		}else{
			if num>=5{
				//tmp=append(tmp,GetTypeCode(4,num,i))
				for j:=i-num+1;j<i;j++{
					(*arr)[j]--
				}
				return append(GetSS(arr),GetTypeCode(4,num,tmpCard))
			}
			num=0
		}
	}
	if num>=5{
		for j:=tmpCard-num+1;j<=tmpCard;j++{
			(*arr)[j]--
		}
		return append(GetSS(arr),GetTypeCode(4,num,tmpCard))
	}
	return res
}