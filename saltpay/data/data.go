package data

import (
	"encoding/json"
	"io/ioutil"
)

type Person struct {
	FirstName string
	LastName  string
	Birthday  string
}

type Repo struct {
	People []Person
}

func NewRepo(filepath string) (*Repo, error) {
	p, err := setupData(filepath)
	if err != nil {
		return nil, err
	}

	return &Repo{People: p}, nil
}

func (r *Repo) GetAll() []Person {
	return r.People
}

func setupData(filepath string) ([]Person, error) {
	people := []Person{}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		// log.Printf("Erorr while reading from %s: %v \n", filepath, err)
		return people, err
	}

	type BirthDaysData [][]string
	var lines BirthDaysData
	// unmarshall it
	err = json.Unmarshal(data, &lines)
	if err != nil {
		// log.Printf("Erorr while parsing data form %s: %v \n", filepath, err)
		return people, err
	}

	// transform into our data.Person object
	for _, v := range lines {
		// fmt.Printf("record: %#v \n", v)
		p := Person{FirstName: v[1], LastName: v[0], Birthday: v[2]}
		people = append(people, p)
	}
	return people, nil
}
