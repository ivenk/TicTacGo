package main


import "testing"

var checkStateTests = []struct {
	description string
	object ticktacktoe
	expected int
}{
	{"Empty playing field", ticktacktoe{[9]int{0,0,0,0,0,0,0,0,0}, 0}, 0},
	{"Draw", ticktacktoe{[9]int {1, 2 , 1, 2, 2, 1, 1, 1, 2}, 9}, 3},
	{"Victory player one", ticktacktoe{ [9] int {1, 2 , 2, 1 , 0 , 0, 1, 0 , 0}, 5}, 1 },
	{"Victory player two", ticktacktoe{[9] int {2, 1 , 1, 2 , 0, 0, 2, 0, 0}, 5}, 2 },

}

func TestCheckState(t *testing.T) {

}