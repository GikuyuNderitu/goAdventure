package main

import "fmt"

func main() {
    for i := 0; i<= 200; i++{
        fmt.Printf("%d\t%b\t%x\t%q\t\n", i, i, i, i)
    }
}
