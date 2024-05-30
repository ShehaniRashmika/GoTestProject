// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	//strings
// 	var testOne string = "Monkey"
// 	var testTwo = "Rabbit"
// 	//This := could be done only within a function not outside a function
// 	testThree := "Cow"
// 	var testFour string

// 	fmt.Println(testOne, testTwo, testFour, testThree)
// 	//testFour is printed as an empty string like this Monkey Rabbit  Cow

// 	//nums
// 	var testNum1 int8 = 25
// 	var testNum2 = 30
// 	//This := could be done only within a function not outside a function

// 	testNum3 := 40
// 	var testNum4 int
// 	//	testNum4 is printed as zero like this 25 30 0 40
// 	var testNum5 uint = 4  // only positive values
// 	var testNum6 int8 = 14 //only values ranging from -128 to +127

// 	fmt.Println(testNum1, testNum2, testNum4, testNum3, testNum5, testNum6)

// 	//float

// 	var numOne float32 = 12.3
// 	var score float64 = 123456765432123456.78
// 	scoreTwo := 123.3
// 	fmt.Println(numOne, score, scoreTwo)

// 	//arrays
// 	var animals [2]string = [2]string{"rabbit", "dog"}

// 	fruits := [3]string{"apple", "orange", "peach"}
// 	var vegetables = [3]string{"cabbage", "onion", "cucumber"}
// 	age := [3]int{1, 2, 4}

// 	fmt.Println(animals, fruits, vegetables, age)
// 	fmt.Println(len(animals), len(fruits), len(vegetables), len(age))

// 	greetings := "Hello everyone good morning !"

// 	fmt.Println(strings.Contains(greetings, "Hello"))
// 	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
// 	fmt.Println(strings.Index(greetings, "oo"))

// 	names := []string{"ashani", "shehani", "bhashi", "ishi"}

// 	for index, value := range names {

// 		if index == 1 {
// 			fmt.Println("continuing at pos", index)
// 			continue
// 		}
// 		fmt.Printf("the name at index %v is %v \n", index, value)
// 	}

// }
