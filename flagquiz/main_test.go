package main

import "testing"

func testCase(t *testing.T, input, expectedResult []string) {
	actualResult := findBestFlags(input)
	if len(actualResult) != len(expectedResult) {
		t.Fail()
	}
	for i, ar := range actualResult {
		if expectedResult[i] != ar {
			t.Fail()
		}
	}
}

func TestSample1(t *testing.T) {
	testCase(
		t,
		[]string{
			"Green stripe, black stripe, yellow",
			"Red stripe, black stripe, yellow",
			"Red stripe, black stripe, white",
			"Red stripe, green stripe, yellow",
		},
		[]string{
			"Red stripe, black stripe, yellow",
		},
	)
}

func TestSample2(t *testing.T) {
	testCase(
		t,
		[]string{
			"Black, white, pink, shrubbery",
			"Black, white, red, shrubbery",
			"Pink, white, red, shrubbery",
			"Black, pink, red, shrubbery",
		},
		[]string{
			"Black, white, red, shrubbery",
		},
	)
}

func TestTwoResults(t *testing.T) {
	testCase(
		t,
		[]string{
			"Green stripe, black stripe, yellow",
			"Red stripe, black stripe, yellow",
			"Red stripe, black stripe, yellow",
			"Red stripe, black stripe, white",
			"Red stripe, green stripe, yellow",
		},
		[]string{
			"Red stripe, black stripe, yellow",
			"Red stripe, black stripe, yellow",
		},
	)
}
