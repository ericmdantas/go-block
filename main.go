package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	sum := sha256.Sum256([]byte("shhhh\n"))
	fmt.Printf("%x", sum)
}