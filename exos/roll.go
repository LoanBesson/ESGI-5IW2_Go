package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("|", rand.Intn(6), "| |", rand.Intn(20), "| |", rand.Intn(100),"|")
}
