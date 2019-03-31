package main

import(
	"testing"
	"fmt"
)

func TestPrint(t *testing.T){
	// t.SkipNow()
	res := Print1to20()
	fmt.Println("hey")
	if res != 210{
		t.Errorf("Wrong result of Print1to20")
	}
}

func TestPrint2(t *testing.T){
	res := Print1to20()
	res++
	if res !=211 {
		t.Errorf("Test Print2 failed")
	}
}

func TestAll(t *testing.T){
	t.Run("TestPrint", TestPrint)
	t.Run("TestPrint2", TestPrint2)
}

func TestMain(m *testing.M){
	fmt.Println("Test begins...")
	m.Run() 
}

// func aaa(n int) int{
// 	for n > 0{
// 		n--
// 	}
// 	return n
// }

// go test -bench=.
func BenchmarkAll(b * testing.B){
	for n := 0; n < b.N; n++ {
		Print1to20()
		// aaa(n)
	}
}