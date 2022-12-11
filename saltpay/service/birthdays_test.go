package service

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {
	//arrange
	r := NewRepoMock()

	//act
	p := r.GetAll()

	//assert
	if len(p) == 0 {
		t.Errorf("Should have more then 0 records")
	}
}

func TestListBirthDays(t *testing.T) {
	//arrange
	r := NewRepoMock()
	s := NewBirthday(r)

	type TestCase struct {
		description string
		input       Dob
		want        []Person
	}

	// People12Dec := []data.Person{
	// 	{FirstName: "Eddie", LastName: "Tubbs", Birthday: "2021/12/11"},
	// }

	PeopleEmpty := []Person{}

	People28FebWithNormalYear := []Person{
		{FirstName: "Mark", LastName: "Curry", Birthday: "1988/02/29"},
		{FirstName: "Silly", LastName: "Goose", Birthday: "1996/02/29"},
		{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"},
	}
	People29Feb := []Person{
		{FirstName: "Mark", LastName: "Curry", Birthday: "1988/02/29"},
		{FirstName: "Silly", LastName: "Goose", Birthday: "1996/02/29"},
	}
	People28FebWithLeapYear := []Person{
		{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"},
	}
	// NonLeapYear29Feb := []data.Person{}

	tcs := []TestCase{
		// {"11 dec birthdays", Dob{Year: 2022, Month: 12, Day: 11}, People12Dec},
		// {"empty date", Dob{Year: 2022, Month: 12, Day: 33}, PeopleEmpty},
		{"29FebWithLeapYear", Dob{Year: 2024, Month: 02, Day: 29, IsLeap: IsLeapYear(2024)}, People29Feb},
		{"28FebWithNormalYear", Dob{Year: 2022, Month: 02, Day: 28, IsLeap: IsLeapYear(2022)}, People28FebWithNormalYear},
		{"28FebWithLeapYear", Dob{Year: 2020, Month: 02, Day: 28, IsLeap: IsLeapYear(2020)}, People28FebWithLeapYear},
		{"Skipping invalid birthday", Dob{Year: 1982, Month: 02, Day: 22, IsLeap: IsLeapYear(1982)}, PeopleEmpty},
	}

	//act
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			got := s.ListBirthdays(tc.input)
			fmt.Printf("Date:%v:, Got: %v \n", tc.input, got)
			//assert
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("failed with date:%v, want:%v, got:%v \n", tc.input, tc.want, got)
			}
		})
	}

}

func TestComparingTwoSlices(t *testing.T) {
	want := []string{"n", "b"}
	got := []string{"n", "b"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want:%#v, got:%#v \n", want, got)
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

func TestBirthayStringToDobFormat(t *testing.T) {
	tcs := []struct {
		description string
		input       string
		want        error
	}{
		{"empty dob string", "", ErrParsingDate},
		{"jibberish string", "adfasdf", ErrParsingDate},
		// {"2022/12/11", "2022/12/11", nil, Dob{Year: 2022, Day: 11, Month: 12}},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			_, got := birthdayDobFormat(tc.input)
			if got != tc.want {
				t.Errorf("input:%v, want:%v, got:%v \n", tc.input, tc.want, got)
			}
		})
	}

	tcsP := []struct {
		description string
		input       string
		want        Dob
	}{
		{"valid date as string", "2022/12/11", Dob{Year: 2022, Day: 11, Month: 12}},
	}

	for _, tc := range tcsP {
		t.Run(tc.description, func(t *testing.T) {
			got, _ := birthdayDobFormat(tc.input)
			if got != tc.want {
				t.Errorf("input:%v, want:%v, got:%v \n", tc.input, tc.want, got)
			}
		})
	}

}
