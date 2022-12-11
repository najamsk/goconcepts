package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrParsingDate = constError("Cant parse the birthday string")
)

type Person struct {
	FirstName string
	LastName  string
	Birthday  string
}
type Birthdays struct {
	repo IRepo
}

type IRepo interface {
	GetAll() []Person
}

func NewBirthday(repo IRepo) *Birthdays {
	return &Birthdays{repo: repo}
}

func (s *Birthdays) ListAll() []Person {
	fmt.Println("list people from service")
	return s.repo.GetAll()
}

func (s *Birthdays) ListBirthdays(forDate Dob) []Person {
	fmt.Println("list people with birthdays from service")
	party := []Person{}
	// tDob := todayDobFormat()
	p := s.repo.GetAll()

	for _, v := range p {
		chk := isYourBirthdayForDate(v, forDate)
		if chk {
			party = append(party, v)
		}
		// fmt.Printf("name: %s, birthday:%#v \n", v.FirstName, birthdayDobFormat(v.Birthday))
	}
	return party
}

func isYourBirthdayForDate(p Person, forDate Dob) bool {
	//if person year is leap and having feb as a month and date 29
	//	then check if today year is also leap
	//	else normal checking of month and date
	pDob, err := birthdayDobFormat(p.Birthday)
	if err != nil {
		log.Printf("error parsing %s %s birthday with error:%v", p.FirstName, p.LastName, err)
		return false
	}
	tDob := forDate

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

func birthdayDobFormat(d string) (Dob, error) {
	dd := strings.Split(d, "/")
	yy, err := strconv.Atoi(dd[0])
	if err != nil {
		return Dob{}, ErrParsingDate
	}
	days, err := strconv.Atoi(dd[2])
	if err != nil {
		return Dob{}, ErrParsingDate
	}
	mm, err := strconv.Atoi(dd[1])
	if err != nil {
		return Dob{}, ErrParsingDate
	}

	dob := Dob{Year: yy, Month: mm, Day: days, IsLeap: IsLeapYear(yy)}
	return dob, nil
}
