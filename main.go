// All our code must belong to a PACKAGE
// The first statement in Go file must be PACKAGE
package main

// import built-in package in Go
import "fmt"

// ENTRYPOINT: Go need to know where to start our program
// ENTRYPOINT is a MAIN function
// meaning that for ONE Go application, we only need ONE MAIN function bc we only have ONE ENTRYPOINT
func main() {
	fmt.Println("Hello world")
}
