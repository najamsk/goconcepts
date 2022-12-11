package main

import (
	"salty/data"
	"salty/service"
	"testing"
	"time"
)

func TestReadingFromJsonFile(t *testing.T) {
	repo, err := data.NewRepo()
	if err != nil {
		t.Fatalf("cant get repo with error :%#v \n", err)
	}
	srv := service.NewBirthday(repo)
	people := srv.ListAll()
	t.Logf("response is :%#v \n", people)

	if len(people) == 0 {
		t.Errorf("people list is empty")
	}

	// fmt.Println("test pass")
}

func IsLeapYear(y int) bool {

	// convert int to Time - use the last day of the year, which is 31st December
	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	days := year.YearDay()

	if days > 365 {
		return true
	} else {
		return false
	}
}
func TestLeapYearTable(t *testing.T) {
	tcs := []struct {
		description string
		input       int
		want        bool
	}{
		{"2024 is leap year", 2024, true},
		{"2000 is leap year", 2000, true},
		{"2004 is leap year", 2004, true},
		{"1996 is leap year", 1996, true},
		{"2022 is not leap year", 2022, false},
		{"1995 is not leap year", 1995, false},
		{"1997 is not leap year", 1997, false},
		{"1998 is not leap year", 1998, false},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			got := IsLeapYear(tc.input)
			if got != tc.want {
				t.Errorf("failed with year:%v, want:%v, got:%v \n", tc.input, tc.want, got)
			}
		})

	}
}
