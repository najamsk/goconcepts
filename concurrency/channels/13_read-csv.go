package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type City struct {
	Name       string
	Population int
}

func main() {
	fmt.Println("working")
	f, err := os.Open("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	start := time.Now()
	rows := genRows(f)
	ur1 := upperCityName(rows)
	ur2 := upperCityName(rows)
	ur3 := upperCityName(rows)

	i := 0
	for c := range fanIn(ur1, ur2, ur3) {
		log.Println(c)
		i++
	}
	fmt.Println("rows =", i)
	fmt.Println("done=", time.Since(start))
	// fmt.Println(upperRows)
}

func fanIn(chans ...<-chan City) <-chan City {
	out := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, c := range chans {
		go func(city <-chan City) {
			for r := range city {
				out <- r
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func upperCityName(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			out <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
		}
		close(out)
	}()
	return out
}

func genRows(r io.Reader) chan City {
	out := make(chan City)
	go func() {
		reader := csv.NewReader(r)
		//reading first row which is actually header row so discard it
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		//reading all the data rows
		for {
			row, err := reader.Read()
			//if we reach file end break loop
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			populationInt, err := strconv.Atoi(row[9])
			out <- City{
				Name:       row[1],
				Population: populationInt,
			}
		}
		close(out)
	}()
	return out
}
