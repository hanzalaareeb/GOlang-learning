package main

// main function is only needed in the main package not needed oother packagese
import (
	"first_test/pkg/class01" // Import package Helperfunction
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("dead world!")
	var IntNum int = 23232
	fmt.Println(IntNum)
	var FloNum float32 = 9.006
	fmt.Println(FloNum)
	var myString string = `what the ffff`
	PrintlnValue(myString)
	var num1 int = 32
	var num2 int = 6
	var string1, result1, result2 = inteigerDivision(num1, num2)
	fmt.Println(string1, result1, result2)
	giveMeStringLen(myString)
	var result3, result4, result5 = tryingNew(53, 23, "class")
	fmt.Println(result3, "|", result4, "|", result5)
	fmt.Println(class01.HelperFunction())
}

func PrintlnValue(value string) {
	fmt.Println(value)
}

func inteigerDivision(int1 int, int2 int) (string, int, int) { // wtf is now know how to commit in this fucking lang
	var temp int = int1 / int2
	var temp2 int = int1 % int2
	return "wtf is ans", temp, temp2
}

func giveMeStringLen(string1 string) {
	fmt.Println(utf8.RuneCountInString(string1))
}

func tryingNew(num1 int, num2 int, str string) (string, int, int) {
	var temp int
	var temp2 int
	temp = num1 / num2
	temp2 = num1 % num2
	return str, temp, temp2
}
