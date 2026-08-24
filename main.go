package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	patterns := []struct {
		Name string
		Re   string
	}{
		{"Credit Card", `4[0-9]{3}-?[0-9]{4}-?[0-9]{4}-?[0-9]{4}`},
		{"SSN", `[0-9]{3}-[0-9]{2}-[0-9]{4}`},
		{"Email", `[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-z]+`},
		{"Phone", `\+?[0-9]{10,13}`},
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: dlp-scanner <file>")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	found := 0
	ln := 0
	for scanner.Scan() {
		ln++
		line := scanner.Text()
		for _, p := range patterns {
			re, err := regexp.Compile(p.Re)
			if err != nil {
				continue
			}
			if re.MatchString(line) {
				fmt.Printf("  [LINE %d] %s detected: %s\n", ln, p.Name, strings.TrimSpace(line[:min(len(line), 60)]))
				found++
			}
		}
	}
	fmt.Printf("\nDLP Scan: %d potential data leaks in %d lines\n", found, ln)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}