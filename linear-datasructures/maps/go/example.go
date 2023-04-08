package main

import "fmt"

func main() {
	Maps1()
	Maps2()
}

/////---------------------------MAPS--------------------//
//Maps are used to keep track of keys that are types, such as integers, strings, float, double,
//pointers, interfaces, structs, and arrays

//(1)creating a map of key strings and value strings
var items = map[string]string{
	"Shop 1": "Phones",
	"Shop 2": "tablets",
	"Shop 3": "laptops",
}

//iterate the items in a loop, and print the keys and values in the code
var i string      //i =key string [string]
var values string //this the value  string
func Maps1() {
	for i, values = range items {
		fmt.Println("items", i, ":", values)
	}
}

//(2) Creating a map using the make method, specifying the key type and the value type.
func Maps2() {
	var products = make(map[int]string)
	products[100] = "shirts"
	products[32] = "Jerseys"
	products[12] = "Socks"
	products[87] = "Balls"
	products[10] = "Scarfs"

	fmt.Println(products)
	fmt.Println(products[12])

}
