package main

import "fmt"

func main() {
    fmt.Println("This program uses a loop to evaluate how many times the variable \n\"i\" will be looped through when it's compared to a value where \nit divides the number 200.")
	for i := 1; i < 200/i; i++ {
        eval := 200/i
		fmt.Printf("%b\t -\t %d\t -\t %x\t-\tComparison Value was %v\n", i, i, i, eval)
	}
}
