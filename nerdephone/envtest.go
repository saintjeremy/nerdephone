package main

import (
	"fmt"
	"os"
)

func main() {
	path, exists := os.LookupEnv("PATH")

	if exists {

	fmt.Print(path)
	}
}
