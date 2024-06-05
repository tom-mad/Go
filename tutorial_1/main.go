package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!")

	var intNum uint16 = 10
	fmt.Println(intNum)

	var floatNum float64 = 12.63242
	fmt.Println(floatNum)

	var result float64 = floatNum + float64(intNum)
	fmt.Println(result)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)
	fmt.Println(len(myString)) // number of bytes
	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	// myConst = "change illegal"
}
