package service

import (
	"fmt"
	"salty/data"
	"strconv"
	"strings"
	"time"
)

type Birthdays struct {
	repo IRepo
}

type IRepo interface {
	GetAll() []data.Person
}

func NewBirthday(repo IRepo) *Birthdays {
	return &Birthdays{repo: repo}
}

// func NewBirthday(repo *data.Repo) *Birthdays {
// 	return &Birthdays{repo: repo}
// }

func (s *Birthdays) ListAll() []data.Person {
	fmt.Println("list people from service")
	return s.repo.GetAll()
}

func (s *Birthdays) ListTodayBirthdays() []data.Person {
	fmt.Println("list people with birthdays from service")
	party := []data.Person{}
	// tDob := todayDobFormat()
	p := s.repo.GetAll()

	for _, v := range p {
		chk := isYourBirthdayToday(v)
		if chk {
			party = append(party, v)
		}
		// fmt.Printf("name: %s, birthday:%#v \n", v.FirstName, birthdayDobFormat(v.Birthday))
	}
	return party
}

func isYourBirthdayToday(p data.Person) bool {
	//if person year is leap and having feb as a month and date 29
	//	then check if today year is also leap
	//	else normal checking of month and date
	pDob := birthdayDobFormat(p.Birthday)
	tDob := todayDobFormat()

	// fmt.Printf("checking:%v, dob: %#v \n", p, pDob)

	if pDob.IsLeap && pDob.Month == 2 && pDob.Day == 29 {
		//match with today is leapYear
		if tDob.Day == 29 && tDob.Month == 2 && tDob.IsLeap {
			//if tody is 29 feb due to leap year
			return true
		} else if tDob.Day == 28 && tDob.Month == 2 && !tDob.IsLeap {
			//if tody is 28 feb due not leap year
			return true
		} else {
			return false
		}
	} else if pDob.Month == tDob.Month && pDob.Day == tDob.Day {
		return true
	}

	return false
}

func IsLeapYear(y int) bool {
	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	days := year.YearDay()

	if days > 365 {
		return true
	} else {
		return false
	}
}

type Dob struct {
	Day    int
	Month  int
	Year   int
	IsLeap bool
}

func birthdayDobFormat(d string) Dob {
	dd := strings.Split(d, "/")
	yy, _ := strconv.Atoi(dd[0])
	days, _ := strconv.Atoi(dd[2])
	mm, _ := strconv.Atoi(dd[1])

	dob := Dob{Year: yy, Month: mm, Day: days, IsLeap: IsLeapYear(yy)}
	return dob
}

func todayDobFormat() Dob {
	currentTime := time.Now()
	d := Dob{
		Day:    currentTime.Day(),
		Month:  int(currentTime.Month()),
		Year:   currentTime.Year(),
		IsLeap: IsLeapYear(currentTime.Year()),
	}
	return d
}
