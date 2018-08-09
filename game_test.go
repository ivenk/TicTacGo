package main


import "testing"

var checkStateTests = []struct {
	description string
	object ticktacktoe
	expected int
}{
	{"Empty playing field", ticktacktoe{[9]int{0,0,0,0,0,0,0,0,0}, 0}, 0},
}

func TestCheckState(t *testing.T) {

}