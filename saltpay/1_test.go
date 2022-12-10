package main

import (
	"fmt"
	"salty/data"
	"salty/service"
	"strings"
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

func translateDateFormat(d string) string {
	return strings.ReplaceAll(d, "/", "-")
}
func TestTranslateDateFormat(t *testing.T) {
	// fmt.Println("test pass")
	i := "1986/03/28"
	expected := "1986-03-28"
	got := translateDateFormat(i)
	// currentTime := time.Now()
	// cf := currentTime.Format("2006-02-01")

	// fmt.Println("MM-DD-YYYY : ", cf)
	if got != expected {
		t.Errorf("date format change failed expected:%s, got:%s \n", expected, got)
	}

}

type Dob struct {
	Day    int
	Month  string
	Year   int
	IsLeap bool
}

func IsLeapYear(y int) bool {

	//Thirty days hath September,
	//April, June and November;
	//February has twenty eight alone
	//All the rest have thirty-one
	//Except in Leap Year, that's the time
	//When February's Days are twenty-nine

	// convert int to Time - use the last day of the year, which is 31st December
	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	days := year.YearDay()

	if days > 365 {
		return true
	} else {
		return false
	}
}

func returnTodayInDobFormat() Dob {
	currentTime := time.Now()
	d := Dob{
		Day:    currentTime.Day(),
		Month:  currentTime.Month().String(),
		Year:   currentTime.Year(),
		IsLeap: IsLeapYear(currentTime.Year()),
	}
	return d
}
func TestTodayDateComponnents(t *testing.T) {

	currentTime := time.Now()
	// cf := currentTime.Format("2006-02-01")
	expected := Dob{
		Day:    currentTime.Day(),
		Month:  currentTime.Month().String(),
		Year:   currentTime.Year(),
		IsLeap: IsLeapYear(currentTime.Year()),
	}
	got := returnTodayInDobFormat()
	fmt.Printf("got:%v \n", got)
	fmt.Printf("expected:%v \n", expected)
	if got != expected {
		t.Errorf("date format change failed expected:%v, got:%v \n", expected, got)
	}
}

func TestLeapYears(t *testing.T) {
	if IsLeapYear(1998) == true {
		t.Errorf("failed with year:%v \n", 1998)
	}
	if IsLeapYear(1997) == true {
		t.Errorf("failed with year:%v \n", 1997)
	}
	if IsLeapYear(1995) == true {
		t.Errorf("failed with year:%v \n", 1995)
	}
	if IsLeapYear(1996) == false {
		t.Errorf("failed with year:%v \n", 1995)
	}

}
