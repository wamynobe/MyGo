package main

import "testing"

type addTest struct {
	arg1, expected string
}

var addTests = []addTest{
	addTest{"0000010101111001", "15"},
	addTest{"0101111100000101", "01"},
	addTest{"0111011001110101", "23"},
	addTest{"0010110101111001", "45"},
	addTest{"011110110100010101111111", "678"},
	addTest{"010001010111111101111101", "789"},
	addTest{"00000101", "1"},
}

func TestConvert(t *testing.T) {
	for _, test := range addTests {
		if output := ElectronicScreen(test.arg1); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}

}
