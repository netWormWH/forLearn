package learTesting

import "testing"

func BenchmarkForTest(b *testing.B){
	obj:=NewForTest()
	
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj.SetKey("A", 100)
			obj.Remove("A")
			if obj.Has("A") {
				b.Fatal("Error")
			}
		}
	})
}