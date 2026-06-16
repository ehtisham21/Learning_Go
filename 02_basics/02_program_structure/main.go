// Overall Structure of a Go File

package main

import "fmt"

const AppName = "Learning Go"

var version = "1.0"

func main() {
	fmt.Println(AppName)
	printVersion()
}

func printVersion() {
	fmt.Println(version)
}
