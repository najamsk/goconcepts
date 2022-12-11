package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type Dob struct {
	Day    int
	Month  string
	Year   int
	IsLeap bool
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

func translateDateFormat(d string) string {
	return strings.ReplaceAll(d, "/", "-")
}
