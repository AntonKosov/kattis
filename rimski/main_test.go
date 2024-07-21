package main

import "testing"

func testCase(t *testing.T, input, expectedResult string) {
	if actualResult := findMin(input); actualResult != expectedResult {
		t.Fail()
	}
}

func TestVI(t *testing.T) {
	testCase(t, "VI", "IV")
}

func TestVII(t *testing.T) {
	testCase(t, "VII", "VII")
}

func TestIII(t *testing.T) {
	testCase(t, "III", "III")
}

func TestXCIX(t *testing.T) {
	testCase(t, "XCIX", "XCIX")
}

func TestXXXI(t *testing.T) {
	testCase(t, "XXXI", "XXIX")
}

func TestXXXVI(t *testing.T) {
	testCase(t, "XXXVI", "XXXIV")
}

func TestXLIX(t *testing.T) {
	testCase(t, "XLIX", "XLIX")
}

func TestXLI(t *testing.T) {
	testCase(t, "XLI", "XLI")
}

func TestLXXVI(t *testing.T) {
	testCase(t, "LXXVI", "LXXIV")
}

func TestLI(t *testing.T) {
	testCase(t, "LI", "LI")
}

func TestLX(t *testing.T) {
	testCase(t, "LX", "XL")
}

func TestXCVI(t *testing.T) {
	testCase(t, "XCVI", "XCIV")
}

func TestLXI(t *testing.T) {
	testCase(t, "LXI", "XLI")
}

func TestLXXXI(t *testing.T) {
	testCase(t, "LXXXI", "LXXIX")
}

func TestLXIV(t *testing.T) {
	testCase(t, "LXIV", "XLIV")
}

func TestLIX(t *testing.T) {
	testCase(t, "LIX", "XLI")
}

func TestXI(t *testing.T) {
	testCase(t, "XI", "IX")
}

func TestXVI(t *testing.T) {
	testCase(t, "XVI", "XIV")
}

func TestXXI(t *testing.T) {
	testCase(t, "XXI", "XIX")
}

func TestLXIX(t *testing.T) {
	testCase(t, "LXIX", "XLIX")
}

func TestLXXIX(t *testing.T) {
	testCase(t, "LXXIX", "LXXIX")
}
