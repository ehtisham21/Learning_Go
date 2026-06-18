package main

import "fmt"

func main() {

	// Byte declaration
	var b byte = 65

	fmt.Println("=== Byte Example ===")

	fmt.Println("Byte value:", b)
	fmt.Printf("As character: %c\n", b)

	fmt.Println("\n=== Byte Type ===")

	fmt.Printf("Type of b: %T\n", b)

	fmt.Println("\n=== String to Bytes ===")

	text := "Go"

	data := []byte(text)

	fmt.Println("String:", text)
	fmt.Println("Bytes:", data)

	fmt.Println("\n=== Bytes to String ===")

	bytesData := []byte{71, 111}

	fmt.Println("String:", string(bytesData))

	fmt.Println("\n=== Accessing Bytes in a String ===")

	word := "Go"

	fmt.Println("word[0]:", word[0])
	fmt.Println("word[1]:", word[1])

	fmt.Printf("word[0] as character: %c\n", word[0])
	fmt.Printf("word[1] as character: %c\n", word[1])

	fmt.Println("\n=== Byte Count vs Actual Bytes ===")

	chineseWord := "你"

	fmt.Println("Byte count:", len(chineseWord))
	fmt.Println("Actual bytes:", []byte(chineseWord))

	fmt.Println("\n=== Byte Slice ===")

	message := "Hello"

	byteSlice := []byte(message)

	fmt.Println("Message:", message)
	fmt.Println("Byte Slice:", byteSlice)

	fmt.Println("\n=== Converting Byte Slice Back to String ===")

	fmt.Println(string(byteSlice))
}
