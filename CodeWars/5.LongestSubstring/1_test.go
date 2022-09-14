package main

import "testing"

type TestCase struct {
	str  string
	want int
}

func Test_LongesSubstring(t *testing.T) {
	testCases := []TestCase{
		{str: "abcabcbb", want: 3},
		{str: "abcaijklm", want: 5},
		{str: "bbbb", want: 1},
		{str: "pwwkew", want: 3},
	}

	for _, test := range testCases {
		got := LongestSub(test.str)
		if got != test.want {
			t.Errorf("string:%s, got=%d, want=%d fails \n", test.str, got, test.want)
		}
		t.Logf("string:%s, got=%d, want=%d passed \n", test.str, got, test.want)
	}

}
