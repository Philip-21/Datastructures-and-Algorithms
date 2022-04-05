package main //package name
import (
	"fmt"     //this package allows us to format strings
	"strconv" // this package converts an int to a string
)

// naming variables in go thenparsing into a function  ,starting with capital letter it will be exported from the packge ,lowercase letter imported to the package

//(A)variable types in go

// multiple variabletypes
var (
	firstname     = "Philip"
	lastname      = "Obiora"
	role          = "mL engineer"
	id        int = 21
)
var b float32 = 34.22 //declaring variable types at package level

func main() { // the main() is an entry point to my application
	fmt.Println(firstname, lastname, role, id)
	fmt.Printf("%v,%T", b, b)

	//declaring a variable type,//then assigning it in a function
	var i int = 42 // used when go doesnt have enough information to assign the type we want to assign to it
	fmt.Println(i)
	var j int //used when declaring a variable but not ready to initialize it yet, used when the compiler guesses wrong
	j = 56
	fmt.Println(j)
	k := 22                   //this method is commonly used
	fmt.Printf("%v,%T", k, k) //another print method (variable,var type,assigned values,assigned value)
	var m float32 = 35        //declaring 35 as afloating point number
	fmt.Printf("%v,%T", m, m)

	//working with acronyms, acronyms can be in block letters for esy readability
	var URL string = "https://google.com"
	fmt.Println(URL)
	var HTTP string = "kaggle.org"
	fmt.Println(HTTP)

	// convert from one variable type to another
	// (1)float to int
	var y float32 = 45.67
	fmt.Printf("%v,%T\n", y, y)
	var z int
	z = int(y) //y is converted to an integer
	fmt.Printf("%v,%T\n", z, z)
	// (2) from int to string
	var s int = 34
	fmt.Printf("%v,%T\n", s, s)
	var name string
	name = strconv.Itoa(s)
	fmt.Printf("%v,%T\n", name, name)

	//(B) Boolean data in go
	// comparison operator to get the boolean values (True,false)
	e := 12 == 12
	c := 12 == 13
	fmt.Printf("%v,%T\n", e, e)
	fmt.Printf("%v,%T\n", c, c)

	//(C) the bit operator (binary representation)
	l := 5 //0101
	o := 4 //0100
	fmt.Println(l & o)
	fmt.Println(l | o)
	fmt.Println(l ^ o)
	fmt.Println(l &^ o)

	//(D) CONSTANTS
	//(naming constants)
	const myConst = "golang"
	fmt.Printf("%v,%T\n", myConst, myConst)
	//enumerated constants
	const (
		a = iota // iota is used when creating enumrated const
		b
		d
	)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", d)
	// using bitshifting to set booleans ina sigle byte
	const (
		isAdmin   = 1 << iota //setting the value to 1 bitshifted by iota (00000001)=1
		isStudent             //(00000010)=2

		canCreateUser //(00000100)=4
		canDeleteUser //(00001000)=8
		canMonitorAccounts
	) // 8 bits make a byte soeach of the constants will occupy 1 location in a byte
	var roles byte = isAdmin | canCreateUser | canDeleteUser   //using or statement if at least one of the values are true
	fmt.Printf("%b\n", roles)                                  //to get the result that encodes the roles in binary format
	fmt.Printf("IS Admin? %v\n", isAdmin&roles == isAdmin)     // getting the boolean statement of the admin
	fmt.Printf("IS Student? %v", isStudent&roles == isStudent) //getting the booolean statement of student

	// (E) ARRAYS AND SLICES //COLLECTION TYPES
	//(1)ARRAYS
	// arrays form the basis of slices
	// parsing book titles into an array of 4 characters which is the sizeof the array
	title := [4]string{"the trenches", "the birches", "swindlers", "cannons"}
	fmt.Printf("book titles: %v", title)
	// defining index of arrays
	var students [4]string
	students[0] = "Patrick"
	students[1] = "Ammy"
	students[2] = "Poly"
	students[3] = "Paul"
	fmt.Printf("list of students are : %v\n", students)
	fmt.Printf("third student is: %v\n", students[3])         //studenst with third index
	fmt.Printf("number of students are: %v\n", len(students)) //getting the number of items in the array or size of the array
	// arrays of arrays with matrices
	// an array of 4 sizes and 3 integers in each sizes
	var matrix [4][3]int = [4][3]int{[3]int{2, 1, 9}, [3]int{1, 1, 1}, [3]int{1, 1, 5}, [3]int{2, 3, 4}}
	fmt.Println(matrix)
	// array of 3 sizes and 3 integers in each size using the index
	var identitymatrix [3][3]int
	identitymatrix[0] = [3]int{3, 4, 5}
	identitymatrix[1] = [3]int{2, 1, 4}
	identitymatrix[2] = [3]int{16, 7, 8}
	fmt.Println(identitymatrix)
	//assigning arrays
	r := [...]int{1, 4, 5}
	q := r
	q[2] = 8 //Assigning 8 to the 2nd index of r
	fmt.Println(r)
	fmt.Println(q)

	//(2)SLICES
	sl := [8]int{2, 3, 4, 55, 66, 7, 12, 34}
	fmt.Printf("length : %v\n", len(sl))
	fmt.Printf("capacity : %v\n", cap(sl))
	cl := []int{1, 2, 3, 4, 5, 76, 23, 89, 56, 40, 100}
	cb := cl[:]   //slice all elements
	df := cl[3:]  // slice from 4th element
	db := cl[2:8] // slice from the 3rd element to 8th element
	fmt.Println(cl)
	fmt.Println(cb)
	fmt.Println(df)
	fmt.Println(db)
	// applying append () in slicing
	cv := []int{}
	fmt.Println(cv)
	fmt.Printf("length: %v\n", len(cv))
	fmt.Printf("capacity: %v\n", cap(cv))
	cv1 := append(cv, []int{2, 3, 4, 5}...)
	fmt.Println(cv1)
	fmt.Printf("length: %v\n", len(cv1))
	//stack operations in slicing
	stack := []int{61, 982, 763, 246, 985}
	stack1 := stack[:len(stack)-1] //getting the first 4 elements
	fmt.Println(stack1)

	//(F)MAPS AND STRUCTS
	//(1)MAPS (maps are used for matching a key type to a value type ), when working with maps or slices you dont need pointers, they both have internal pointers to their underlined data
	fansattendance := map[string]int{
		"Leeds":     34458888,
		"Westham":   2456678765,
		"Norwich":   3451768,
		"Chelsea":   4546434323598,
		"Liverpool": 343523412376,
		"Mancity":   767533223476,
		"Arsenal":   45463262465,
	} //key type are the club names , value types are the integers
	fmt.Println(fansattendance)
	//adding another key type to the map
	fansattendance["Manutd"] = 15469087523
	fmt.Println(fansattendance["Manutd"])
	//deleting a key from the map
	delete(fansattendance, "Norwich")
	fmt.Println(fansattendance)
	//(2)STRUCT
	// struct gathers information together related to a concept any tpe of data can be stored in a struct
	type footballer struct {
		name      string
		number    int
		positions []string
	}
	Profile := footballer{
		name:   "perry",
		number: 3,
		positions: []string{
			"center back",
			"defensive midfielder",
			"left back",
			"wing back",
		},
	}
	fmt.Println(Profile)
	//getting a paricular field
	fmt.Println(Profile.number)
	fmt.Println(Profile.positions[1])
	//another method of parsing struct
	Doctor := struct{ name string }{name: "Dennis Falker"}
	fmt.Println(Doctor)
	// adding another name to the doctor field
	anotherDoctor := Doctor
	anotherDoctor.name = "Elliot Barnes"
	fmt.Println(anotherDoctor)
	//EMBEDDINGS in STRUCTS
	//go supports composition through embedding
	type Animal struct {
		Name   string
		Origin string
	}
	type Bird struct {
		Animal   //embedding animal struct into Bird struct
		SpeedKPH float32
		CanFly   bool
	}
	bd := Bird{}
	bd.Name = "Eagle"
	bd.Origin = "America"
	bd.SpeedKPH = 20
	bd.CanFly = true
	fmt.Println(bd)

	//(G)IF statements
	examscore := 65
	testscore := 45
	if testscore < examscore {
		fmt.Println("average score")
	}
	if testscore > examscore {
		fmt.Println("total score")
	}
	if testscore == examscore {
		fmt.Println("no score yet")
	}
	fmt.Println(testscore < examscore, testscore > examscore, testscore != examscore)
	// the >,<,<=,>=work for numeric types not string types,== or !=work for string types
	// combining logical tests together using logical operators
	number := 45
	guess := 37
	unit := 19
	if guess < 1 || guess > 100 { //using an or logic in the if statement
		fmt.Println("figures less than 100")
	}
	if guess > 30 && guess < 100 { //using & logic in if statement
		fmt.Println("guess intervals")
	}
	if guess < number {
		fmt.Println("low figure")
	}
	if guess == number {
		fmt.Println("equal")
	}
	fmt.Println(number <= guess, number >= guess, number != guess) //gets the boolean values ofthe print statements
	//using else statement
	if number < unit {
		fmt.Println("this is the unit")
	} else if number > unit {
		fmt.Println("number")
	} else {
		fmt.Println("no result")
	}

	//(H)SWITCH STATEMENTS
	//this is a shorter way to 	write a sequence of if else statement
	switch 1 { //this switch expression(1) prints out the case 1condition
	case 1:
		fmt.Println("first expression")
	case 2:
		fmt.Println("second expression")
	case 3:
		fmt.Println("3RD expression")
	}
	id := 50
	switch {
	case id < 40:
		fmt.Println("less than 40")
	case id <= 50:
		fmt.Println("fifty")
		fallthrough //this enables the nextcase to be executed
	case id > 40:
		fmt.Println("greater than forty")

	}

	//LOOPING
	//(I) FOR STATEMENTS
	for x := 0; x < 5; x = x + 2 { // incrementing it by 2 (0,2,4)
		fmt.Println(x)
	}
	//increamenting two items
	for u, w := 0, 0; u < 5; u, w = u+1, w+1 { // incrementing by 1
		fmt.Println(u, w)
	}
	for z := 0; z < 5; z++ {
		fmt.Println(z)
		if z%2 == 0 { //odd number that will give remainders willpass through this statement
			z = z / 2
		} else {
			z = 2*z + z //for even nmbers
		}
	}
	//using continue statement
	for y := 0; y < 10; y++ {
		if y%2 == 0 {
			continue //used to determine if you want to process a record or not
		}
		fmt.Println(y)
	}
	//nested loop
loop:
	for s := 1; s <= 3; s++ {
		for q := 3; q <= 7; q++ {
			fmt.Println(q * s)
			if q*s >= 4 {
				break loop
			}
		}
	}
	// for range loop
	sv := []int{1, 2, 3, 4, 5}
	for item := range sv {
		fmt.Println(item)
	}

	// (J)POINTERS
	//pointers hold the memory address of the data but not the data itself
	//pointers passes references to values and records within your program
	ab := [5]int{2, 5, 9, 1, 6}
	bc := &ab //bc points to the same data as ab
	bc[3] = 4 //assigning 4 to the 3rd index to bc which would also assign 4 to ab in the 3rd index due to the pointer
	fmt.Println(ab)
	fmt.Println(bc)
	var fg int = 23
	var bb *int = &fg //creatig a pointer btwn bb to fg
	fmt.Println(bb)   //this prints out the memeory location of fg
	fmt.Println(*bb)  //this gets the value 23
	*bb = 12
	fmt.Println(fg) //prints 12 cause bb and fg are pointers to each other
	//Handling variables when assigned to one another
	gy := map[string]string{"menu": "salad", "drink": "vodka"}
	er := gy
	fmt.Println(gy, er)
	gy["menu"] = "pancake"
	fmt.Println(er) // er brings out the changed data of gy (pancake)
}
