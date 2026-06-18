package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Rune declarations
	var letter rune = 'A'

	// Type inference
	var urduLetter = 'ا'

	// Short variable declaration
	emoji := '😀'

	fmt.Println("=== Rune Examples ===")

	fmt.Printf("Letter: %c\n", letter)
	fmt.Printf("Urdu Letter: %c\n", urduLetter)
	fmt.Printf("Emoji: %c\n", emoji)

	fmt.Println("\n=== Unicode Values ===")

	fmt.Printf("A = %d\n", letter)
	fmt.Printf("ا = %d\n", urduLetter)
	fmt.Printf("😀 = %d\n", emoji)

	fmt.Println("\n=== Type Checking ===")

	fmt.Printf("Type of letter: %T\n", letter)

	fmt.Println("\n=== Rune vs String ===")

	char := 'G'
	text := "G"

	fmt.Printf("char type: %T\n", char)
	fmt.Printf("text type: %T\n", text)

	fmt.Println("\n=== String Iteration Using Runes ===")

	word := "Hello"

	for _, ch := range word {
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	fmt.Println("\n=== Unicode String Iteration ===")

	greeting := "你好"

	for _, ch := range greeting {
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	fmt.Println("\n=== Byte Count vs Rune Count ===")

	fmt.Println("Bytes in 'Hello':", len("Hello"))
	fmt.Println("Characters in 'Hello':", utf8.RuneCountInString("Hello"))

	fmt.Println("Bytes in '你好':", len("你好"))
	fmt.Println("Characters in '你好':", utf8.RuneCountInString("你好"))

	fmt.Println("\n=== Rune to String Conversion ===")

	fmt.Println(string(letter))
	fmt.Println(string(emoji))
}