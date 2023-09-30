package main

import (
	"fmt"
)

func main() {

	/**
	if not variable is not used, compile error happens
	*/

	// with data type
	var answer int = 10
	var sentence string = "hello golang"
	fmt.Println(answer)
	fmt.Println(sentence)

	// without data type
	var x = 10
	var stringInput = "golang newbie"
	fmt.Println(x)
	fmt.Println(stringInput)

	// implicit initialization of variable
	a := 10
	b := 20
	fmt.Println(a)
	fmt.Println(b)

	// multiple initialization of variable at a time
	var posX, posY, posZ int = 3, 4, 5
	string1, string2, string3 := "apple", "banana", "kiwi"
	// be careful, there are no ",". and ":"
	// ":" is only required for implicit initialization of variable
	var (
		apple  = "apple"
		banana = "banana"
		kiwi   = "kiwi"
	)

	fmt.Println(posX, posY, posZ)
	fmt.Println(string1, string2, string3)
	fmt.Println(apple, banana, kiwi)

	// constant variable
	const discount int = 10
	const price = 10
	const (
		tomato    = "tomato" // const variable can be not used
		green     = "green"
		newGgreen // if no value is assigned, previous value("green") is assigned
	)

	fmt.Println(green)
	fmt.Println(newGgreen)

}
