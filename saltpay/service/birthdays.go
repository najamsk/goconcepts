package service

import (
	"fmt"
	"salty/data"
)

type Birthdays struct {
	repo *data.Repo
}

func NewBirthday(repo *data.Repo) *Birthdays {
	return &Birthdays{repo: repo}
}

func (s *Birthdays) ListAll() []data.Person {
	fmt.Println("list people from service")
	return s.repo.GetAll()
}
