package main

import (
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	// NOTE:
	// following line works if you supply env variable while running go program eg:
	// foo=bar go run 1.go
	fmt.Printf("reading env variable from shell %v \n", os.Getenv("foo"))

	//NOTE:
	//now we can use "github.com/joho/godotenv" to load .env files and just print variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("loading env using pkg throws err:%v \n", err)
	}
	//now at this point we should have dev variable loaded from env files
	fmt.Printf("dev=%v \n", os.Getenv("dev"))
	fmt.Printf("code=%v \n", os.Getenv("code"))

}
