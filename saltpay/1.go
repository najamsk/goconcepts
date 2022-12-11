package main

import (
	"fmt"
	"log"
	"salty/data"
	"salty/service"
)

func main() {

	//get filename from cli flags
	// fileName := parseFileFlag()
	repo, err := data.NewRepo()
	if err != nil {
		log.Printf("Cant create repo because of error: %v \n", err)
		return
	}
	srv := service.NewBirthday(repo)

	// res := srv.ListAll()
	// for _, v := range res {
	// 	fmt.Printf("person : %#v \n", v)
	// }

	// resB := srv.ListTodayBirthdays()
	fDate := service.Dob{Year: 2021, Month: 12, Day: 10}
	resB := srv.ListBirthdays(fDate)
	for _, v := range resB {
		fmt.Printf("birthday : %#v \n", v)
	}
}