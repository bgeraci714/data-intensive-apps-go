package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(getTimestamp())
	fmt.Println("My favorite number is", rand.Intn(10))
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
