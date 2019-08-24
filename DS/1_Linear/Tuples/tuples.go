package main

import (
	"fmt"
	"time"
)

func getMessage() (string, string) {
	first := "Hello,"
	var second string
	hr := time.Now().Hour()
	if hr < 12 {
		second = "good morning"
	} else if hr > 12 && hr < 16 {
		second = "good afternoon"
	} else {
		second = "good evening"
	}
	return first, second
}
func main() {
	first, second := getMessage()
	fmt.Println(first, second)
}
