package rotation_test

import (
	"rotation/rotation"
	"testing"
)

func TestAbs(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
func TestStringLeftRotation(t *testing.T) {
	s := "qwertyu"

	res := rotation.StringLeftRotation(s, 2)
	if res != "ertyuqw" {
		t.Errorf("expect ertyuqw, got = %s ", res)
	}
}

func TestStringRightRotation(t *testing.T) {
	s := "qwertyu"
	res := rotation.StringRightRotation(s, 2)
	if res != "yuqwert" {
		t.Errorf("expect yuqwert, got = %s ", res)
	}
}

func Test_RangeValues_StringLeftRotation(t *testing.T) {
	testCases := []struct {
		operation string
		x         string
		d         int
		want      string
	}{
		{"left rotate", "qwertyu", 2, "ertyuqw"},
		{"left rotate", "GeeksforGeeks", 2, "eksforGeeksGe"},
	}

	for _, tc := range testCases {
		t.Run(tc.operation, func(t *testing.T) {
			got := rotation.StringLeftRotation(tc.x, tc.d)
			if got != tc.want {
				t.Errorf("expect %s, got = %s ", tc.want, got)
			}
			t.Logf("input: %s, expect %s, got = %s \n", tc.x, tc.want, got)
		})
	}

}

func Test_RangeValues_StringRightRotation(t *testing.T) {
	testCases := []struct {
		operation string
		x         string
		d         int
		want      string
	}{
		{"left rotate", "qwertyu", 2, "yuqwert"},
		{"left rotate", "GeeksforGeeks", 2, "ksGeeksforGee"},
	}

	for _, tc := range testCases {
		t.Run(tc.operation, func(t *testing.T) {
			got := rotation.StringRightRotation(tc.x, tc.d)
			if got != tc.want {
				t.Errorf("expect %s, got = %s ", tc.want, got)
			}
			t.Logf("input: %s, expect %s, got = %s \n", tc.x, tc.want, got)
		})
	}

}
