package main


import "testing"

var checkStateTests = []struct {
	description string
	object ticktacktoe
	expected int
}{
	{"Empty playing field", ticktacktoe{[9]int{0,0,0,0,0,0,0,0,0}, 0}, 0},
	{"Draw", ticktacktoe{[9]int {1, 2 , 1, 2, 2, 1, 1, 1, 2}, 9}, 3},
}

func TestCheckState(t *testing.T) {

}