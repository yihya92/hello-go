package main

import (
	"fmt"

	"github.com/yihya92/hello-go/greet"
	"github.com/yihya92/hello-go/mathutil"
)

func main() {
	fmt.Println("Hello")
	sum := mathutil.Add(1, 2)
	message := greet.Greet("Developer")

	fmt.Println(sum)
	fmt.Println(message)

}
