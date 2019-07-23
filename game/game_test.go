package game

import (
	"fmt"
	"testing"
)

var checkStateTests = []struct {
	description string
	object      Ticktacktoe
	expected    int
}{
	{"Empty playing field", Ticktacktoe{[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 0}, 0},
	{"Draw", Ticktacktoe{[9]int{1, 2, 1, 2, 2, 1, 1, 1, 2}, 9}, 3},
	{"Victory player one", Ticktacktoe{[9]int{1, 2, 2, 1, 0, 0, 1, 0, 0}, 5}, 1},
	{"Victory player two", Ticktacktoe{[9]int{2, 1, 1, 2, 0, 0, 2, 0, 0}, 5}, 2},
}

var makeMoveTests = []struct {
	description string
	object      Ticktacktoe
	move        int
	expected    Ticktacktoe
}{
	{"First move", Ticktacktoe{[9]int{}, 0}, 0, Ticktacktoe{[9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}, 1}},
	{"Any move", Ticktacktoe{[9]int{1, 2, 1, 2, 0, 0, 0, 0, 0}, 4}, 4, Ticktacktoe{[9]int{1, 2, 1, 2, 1, 0, 0, 0, 0}, 5}},
}

func TestCheckState(t *testing.T) {
	for _, test := range checkStateTests {
		actual := test.object.CheckState()
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("%q : \n CheckState(%q): expected %q, actual %q", test.description, test.object.field, test.expected, actual)
		}
	}
}

func TestMakeMove(t *testing.T) {
	for _, test := range makeMoveTests {
		actual := test.object.MakeMove(test.move)
		if actual != test.expected {
			t.Fatalf("%q :\n MakeMove(%q): expected field %q, actual field %q", test.description, test.move, test.expected, actual)
		}

		if actual.moveCounter != test.expected.moveCounter {
			t.Fatalf("%q :\n MakeMove(%q): expected move count %q, actual move count %q", test.description, test.move, test.expected.moveCounter, actual.moveCounter)
		}
	}
}
