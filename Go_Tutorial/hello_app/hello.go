// simple hello world program
package main //package name same as the main funtion

import (
	"Go_Tutorial/greetings" //importing internal dependencies
	"fmt"
	"log"

	"rsc.io/quote" //importing external dependencies
)

// Excercise 1 - simply print "Hello World!!"
func hello() {
	fmt.Println("Priting hello fn  -->")
	fmt.Println("Hello World!")
	fmt.Println()
}

// to show how to import local packages
// add the funtion call from the greetings package
// go mod tidy to add the package
// go mod edit replace to replace the "Go_Web/greetings=../greetings" path
func hello1() {
	fmt.Println("Priting hello1 fn  -->")
	fmt.Println("Hello World!")
	fmt.Println(quote.Go()) // use dependency funtion
	fmt.Println()
}

// To print greeting to the named person - uses greetings/Hello funtion
func hello2() {
	// Get a greeting message and print it.
	fmt.Println("Priting hello2 fn  -->")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	fmt.Println()
}

// To add error management to the check the correct name passing
func hello3() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	fmt.Println("Priting hello3 fn  -->")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello1("")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println()

	fmt.Println("Priting hello4 fn  -->")
	message1, err1 := greetings.Hello2("Shinchan") // print random greeting message
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(message1)
	fmt.Println()

	// print random greeting to random people froma list
	fmt.Println("Priting hello5 fn  -->")
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err2 := greetings.Hellos(names)
	if err2 != nil {
		log.Fatal(err2)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

// main funtion - entry point to the package
func main() {
	hello()
	hello1()
	hello2()
	hello3()
}
