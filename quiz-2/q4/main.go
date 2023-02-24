//manually outputting hexadecimal numbers from 1 to f
//Using WaitGroups
package main


import (
	"fmt"
	"sync"
)
//function prints the hexadecimal alphabet from a to f
//has a pointer to wautgroup object as parameter
func alphabet(wg *sync.WaitGroup) {
defer wg.Done()//done method of waitgroups that signals its comppletion
//for loop to iterate over letter
for char := 'a'; char <= 'f'; char++ {
	fmt.Printf("%c" , char)//print letter
	fmt.Println();//prints line 
}
}
//functions prints the decimal numbers less than nine
//has a pointer tp waitgroup onject as parameter
func number(wg * sync.WaitGroup) {
	//for loop to iterate over the numberes less than 9
	defer wg.Done()//done method of waitgroup to singnal its completion
	for num := 1; num <=9; num++ {
		fmt.Printf("%d", num)
		fmt.Println()
	}

}
func main() {
var wg sync.WaitGroup//create waitgroup object to synchronize the goroutine to complete
wg.Add(2)//two waitgroup counters to wait for the two goroutine to complete


go alphabet(&wg)//runs the alphabet goroutine witht waitgroup as parameter
go number(&wg)//runs the number goroutine witht waitgroup as parameter
wg.Wait()//wai for both goroutine to complete before exiting


}
