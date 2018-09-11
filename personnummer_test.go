package main

import (
	"testing"
)

func TestFormatIsValid(t *testing.T) {
	var tables = []struct {
		input  string
		output bool
	}{
		{"01017092940", true},
		{"01017092789", true},
		{"01017000000", true},
		{"?1017000000", false},
		{"010170?0000", true},
		{"010179990??", true},
		{"0101712345", false},
		{"010170123456", false},
		{"01017f12345d", false},
		{"''u/2x1f|}", false},
		{"'`¥@ð¡đŋħ©ª", false},
		{"¡@£$½¥¥{[]}", false},
		{"@ł€®þ←↓→œππ", false},
		{"20008812345", false},
		{"20138812334", false},
		{"00017012345", false},
		{"32109912345", false},
	}
	for _, table := range tables {
		foo := formatIsValid(table.input)
		if foo != table.output {
			t.Errorf("Failed on test %s, got %t but expected %t\n",
				table.input,
				!table.output,
				table.output,
			)
		}
	}
}

func TestCalculateCtrlNumber(t *testing.T) {
	var foobar = []struct {
		input     [11]int32
		expected1 int32
		expected2 int32
	}{
		{[11]int32{0, 1, 0, 1, 7, 0, 2, 2, 0}, 6, 3},
		{[11]int32{2, 0, 0, 5, 8, 7, 5, 8, 9}, -1, -1},
	}
	for _, table := range foobar {
		actual1, actual2 := calculateCtrlNumber(table.input)
		if actual1 != table.expected1 || actual2 != table.expected2 {
			t.Errorf("Input %d, expected %d and %d, got %d and %d",
				table.input,
				table.expected1,
				table.expected2,
				actual1,
				actual2,
			)
		}

	}

}
