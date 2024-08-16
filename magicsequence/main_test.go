package main

import "testing"

func testCase(t *testing.T, input MagicSequence, expectedResult uint64) {
	actualResult := getHash(input)
	if actualResult != expectedResult {
		t.Fail()
	}
}

func TestSample1(t *testing.T) {
	testCase(
		t,
		MagicSequence{n: 7, a: 7, b: 7, c: 12, x: 1, y: 20},
		0,
	)
}

func TestSample2(t *testing.T) {
	testCase(
		t,
		MagicSequence{n: 7, a: 16, b: 1, c: 15, x: 1, y: 14},
		1,
	)
}

func TestSample3(t *testing.T) {
	testCase(
		t,
		MagicSequence{n: 7, a: 2, b: 5, c: 6, x: 1, y: 19},
		8,
	)
}

func TestSample4(t *testing.T) {
	testCase(
		t,
		MagicSequence{n: 1e6, a: 1e8 + 456, b: 1e7 + 7673465, c: 1e6 + 847238, x: 1e7 + 8765, y: 1e7 + 34343},
		3064624,
	)
}
