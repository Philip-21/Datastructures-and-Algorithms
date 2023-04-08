package main

import (
	"fmt"
	"strconv"
	"time"
)

//Calculate the year given the date of birth and age

func Year() {
	var t time.Time

	res := t.Format("2000-12-21")
	date, err := strconv.Atoi(res)
	if err != nil {
		return
	}

	age := 21

}

//Create a program that calculates the average weight of 5 people.
func Weight() {
	w1 := 29
	w2 := 120
	w3 := 44
	w4 := 56
	w5 := 35
	avg := (w1 + w2 + w3 + w4 + w5) / 5
	fmt.Println(avg)

}
func main() {
	Weight()
}
