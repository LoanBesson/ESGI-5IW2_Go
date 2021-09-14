package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Nous sommes le", time.Now().Day(), time.Now().Month().String() + ", il est", time.Now().Hour(), "h", time.Now().Minute())
}
