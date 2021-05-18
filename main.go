package main

import (
	"time"

	"github.com/barisvelioglu/firstapp/greetings"
)

func main() {
	greetings.Greet("Hello World!!!!!!!")
	time.Sleep(10 * time.Second)
}
