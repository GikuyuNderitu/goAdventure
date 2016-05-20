package main

import (
	"fmt"

	"github.com/GikuyuNderitu/goAdventure/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("Use mirror to read my name, which is below.")
	fmt.Println("\n\b", stringutil.Reverse(stringutil.MyName))
}
