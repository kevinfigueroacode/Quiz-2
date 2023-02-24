// Printing full name using goroutine
package main

import (
	"fmt"
	"time"
)
//function will print first name
func firstname() {
	fmt.Println("Kevin")
}
//function will print second name
func secondname() {
	fmt.Println("Alexander")
}//function prints last name
func lastname() {
	fmt.Println("Figueroa")
}

func main() {
//go routines usually has go before the function
	go firstname()//run goroutines
	go secondname()
	go lastname()
	time.Sleep(time.Second)//this piece of code tells the main to delay the termination of the program
}
