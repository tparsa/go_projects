package main

import (
	"encoding/csv"
	"os"
	"io"
	"strconv"
)

type City struct {
    name    string
    zipcode int
}

var states map[string][]City

func get_city_of_state(state string) []City {
	return states[state]
}

func get_city_zipcode(state string, city_name string) int {
	for _, city := range states[state] {
		if city.name == city_name {
			return city.zipcode
		}
	}
	return 0
}

func get_city_with_zipcode_state(state string, zipcode int) string{
	for _, city := range states[state] {
		if city.zipcode == zipcode {
			return city.name
		}
	}
	return ""
}

func main(){
	csvfile, _ := os.Open("summercamp-golang-instructions-dataset.csv")
	r := csv.NewReader(csvfile)
	states = make(map[string][]City)
	for {
		record, err := r.Read()
		if err == io.EOF{
			break
		}
		zipcode, _ := strconv.Atoi(record[2])
		states[record[0]] = append(states[record[0]], City{record[1], zipcode})
	}
}