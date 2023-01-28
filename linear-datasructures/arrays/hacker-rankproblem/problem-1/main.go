package main

/*The for keyword is followed by three statements, separated by semicolons. These statements are:
The initialization statement: i, j := 0, len(a)-1.
This statement is executed before the loop begins. It initializes the variables i and j to 0 and len(a)-1, respectively.
The loop condition: i < j. This statement is evaluated at the beginning of each iteration of the loop.
If it evaluates to true, the loop continues. If it evaluates to false, the loop terminates.
The post-iteration statement: i, j = i+1, j-1. This statement is executed at the end of each iteration of the loop.
It increments the value of i by 1 and decrements the value of j by 1.
The loop iterates as long as the loop condition (i < j) is true. At the end of each iteration,(conditional statement)
the loop updates the values of i and j using the post-iteration statement (i, j = i+1, j-1).
This causes the loop to eventually terminate when i and j meet or cross each other.*/
import (
	"fmt"
	"os"
)

func reverseArray() ([]int, error) {
	a := []int{1, 4, 3, 2}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println(a[0], a[1], a[2], a[3])
	return a, nil

}

func Errcheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	reverseArray()

}
