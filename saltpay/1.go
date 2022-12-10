package main

import (
	"flag"
	"fmt"
	"log"
	"salty/data"
	"salty/service"
)

func main() {

	//get filename from cli flags
	fileName := parseFileFlag()
	repo, err := data.NewRepo(fileName)
	if err != nil {
		log.Printf("Cant create repo because of error: %v \n", err)
		return
	}
	srv := service.NewBirthday(repo)

	res := srv.ListAll()
	for _, v := range res {
		fmt.Printf("person : %#v \n", v)
	}
}

func parseFileFlag() string {
	iFile := flag.String("file", "data2.json", "provide file to read")
	flag.Parse()
	// if *iFile == "" {
	// 	fmt.Println("no file provided to read from")
	// 	return ""
	// }
	fmt.Printf("will read file: %s \n", *iFile)
	return *iFile
}
