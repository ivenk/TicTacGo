package main

import (
	"testing"
	"fmt"
)

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
	for _, test := range checkStateTests {
		actual := test.object.CheckState()
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("%q : \n CheckState(%q): expected %q, actual %q", test.description,test.object.field, test.expected, actual)
		}
	}
}