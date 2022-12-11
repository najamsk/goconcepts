package main

import (
	"salty/data"
	"salty/service"
	"testing"
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
