package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("=== Numeric Type Conversion ===")

	age := 25
	price := 10.99

	ageAsFloat := float64(age)
	priceAsInt := int(price)

	fmt.Println("age:", age)
	fmt.Println("age as float64:", ageAsFloat)
	fmt.Println("price:", price)
	fmt.Println("price as int:", priceAsInt)

	fmt.Println("\n=== Mixed Type Operation ===")

	total := float64(age) + price
	fmt.Println("float64(age) + price:", total)

	fmt.Println("\n=== String to Int ===")

	textAge := "25"

	convertedAge, err := strconv.Atoi(textAge)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String age:", textAge)
		fmt.Println("Converted int age:", convertedAge)
	}

	fmt.Println("\n=== Int to String ===")

	number := 100

	numberText := strconv.Itoa(number)

	fmt.Println("Number:", number)
	fmt.Println("Number as string:", numberText)

	fmt.Println("\n=== String to Float ===")

	textPrice := "10.99"

	convertedPrice, err := strconv.ParseFloat(textPrice, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String price:", textPrice)
		fmt.Println("Converted float price:", convertedPrice)
	}

	fmt.Println("\n=== Float to String ===")

	floatValue := 99.999

	floatText := strconv.FormatFloat(floatValue, 'f', 2, 64)

	fmt.Println("Float value:", floatValue)
	fmt.Println("Float as string:", floatText)

	fmt.Println("\n=== String to Bool ===")

	textBool := "true"

	convertedBool, err := strconv.ParseBool(textBool)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String bool:", textBool)
		fmt.Println("Converted bool:", convertedBool)
	}

	fmt.Println("\n=== Bool to String ===")

	isActive := true

	boolText := strconv.FormatBool(isActive)

	fmt.Println("Bool value:", isActive)
	fmt.Println("Bool as string:", boolText)

	fmt.Println("\n=== Rune and Byte Conversion ===")

	letter := 'A'
	b := byte(65)

	fmt.Println("Rune to string:", string(letter))
	fmt.Println("Byte to string:", string(b))

	fmt.Println("\n=== String and Byte Slice Conversion ===")

	message := "Hello"

	bytesData := []byte(message)

	fmt.Println("String:", message)
	fmt.Println("String to bytes:", bytesData)
	fmt.Println("Bytes back to string:", string(bytesData))

	fmt.Println("\n=== Type Checking ===")

	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of ageAsFloat: %T\n", ageAsFloat)
	fmt.Printf("Type of numberText: %T\n", numberText)
	fmt.Printf("Type of convertedBool: %T\n", convertedBool)
}
