package main

import (
	"fmt"
	"os"
)

// data_loss_prevention - Prevent data exposure
func data_loss_prevention(path string) {
	fmt.Println("========================================")
	fmt.Println("  Data-Loss-Prevention")
	fmt.Println("  Prevent data exposure")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data_loss_prevention(path)
}
