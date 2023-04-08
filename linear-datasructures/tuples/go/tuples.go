//tuples are data structures that group data into immutable  sequential collections.
package main

import "fmt"

///--the values in the tuples  must be returned after declaring the values in a function

//gets the series of m and returns tuple(int) of square of m
// and cube of m
func Series(m int) (int, int) {
	return m * m, m * m * m
}

//using errors in the  tuples
func powerSeries(g int) (int, int, error) {
	squ := g * g
	cub := squ * g
	return squ, cub, nil
}

//describing the tuples in the function
func IntSeries(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return square, cube
}

//The Main method calls intSeries method with 3 as a parameter.
// The square and cube values are returned from the method
func Main() {
	var square int
	var cube int
	square, cube = IntSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
}
