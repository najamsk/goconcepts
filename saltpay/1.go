package main

import (
	"fmt"
	"log"
	"salty/data"
	"salty/service"
	"time"
)

func main() {

	repo, err := data.NewRepo()
	if err != nil {
		log.Fatalf("Cant create repo because of error: %v \n", err)
	}
	srv := service.NewBirthday(repo)

	// resB := srv.ListTodayBirthdays()
	today := time.Now()

	fDate := service.Dob{Year: today.Year(), Month: int(today.Month()), Day: today.Day()}
	resB := srv.ListBirthdays(fDate)
	for _, v := range resB {
		fmt.Printf("birthday : %#v \n", v)
	}
}
