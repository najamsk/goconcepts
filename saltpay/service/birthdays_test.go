package service

import (
	"salty/data"
	"testing"

	"golang.org/x/exp/slices"
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

	LeapYear29Feb := []data.Person{
		{FirstName: "Mark", LastName: "Curry", Birthday: "1988/02/29"},
		{FirstName: "Silly", LastName: "Goose", Birthday: "1996/02/29"},
	}
	// NonLeapYear29Feb := []data.Person{}

	tcs := []TestCase{
		// {"11 dec birthdays", Dob{Year: 2022, Month: 12, Day: 11}, People12Dec},
		// {"empty date", Dob{Year: 2022, Month: 12, Day: 33}, PeopleEmpty},
		{"29feb", Dob{Year: 2022, Month: 02, Day: 28, IsLeap: IsLeapYear(2022)}, LeapYear29Feb},
	}

	//act
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			got := s.ListBirthdays(tc.input)
			//assert
			if !slices.Equal(got, tc.want) {
				t.Errorf("failed with date:%v, want:%v, got:%v \n", tc.input, tc.want, got)
			}
		})
	}

}
