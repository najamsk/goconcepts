package service

import (
	"fmt"
	"reflect"
	"salty/data"
	"testing"
)

func TestGetAll(t *testing.T) {
	//arrange
	r := data.NewRepoMock()

	//act
	p := r.GetAll()

	//assert
	if len(p) == 0 {
		t.Errorf("Should have more then 0 records")
	}
}

func TestListBirthDays(t *testing.T) {
	//arrange
	r := data.NewRepoMock()
	s := NewBirthday(r)

	type TestCase struct {
		description string
		input       Dob
		want        []data.Person
	}

	// People12Dec := []data.Person{
	// 	{FirstName: "Eddie", LastName: "Tubbs", Birthday: "2021/12/11"},
	// }

	// PeopleEmpty := []data.Person{}

	People28FebWithNormalYear := []data.Person{
		{FirstName: "Mark", LastName: "Curry", Birthday: "1988/02/29"},
		{FirstName: "Silly", LastName: "Goose", Birthday: "1996/02/29"},
		{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"},
	}
	People28FebWithLeapYear := []data.Person{
		{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"},
	}
	// NonLeapYear29Feb := []data.Person{}

	tcs := []TestCase{
		// {"11 dec birthdays", Dob{Year: 2022, Month: 12, Day: 11}, People12Dec},
		// {"empty date", Dob{Year: 2022, Month: 12, Day: 33}, PeopleEmpty},
		{"28FebWithNormalYear", Dob{Year: 2022, Month: 02, Day: 28, IsLeap: IsLeapYear(2022)}, People28FebWithNormalYear},
		{"28FebWithLeapYear", Dob{Year: 2020, Month: 02, Day: 28, IsLeap: IsLeapYear(2020)}, People28FebWithLeapYear},
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