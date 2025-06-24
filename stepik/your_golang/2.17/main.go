package main

import (
	"fmt"
	"strings"
)

func main() {
	// var message string = " Go - это не просто язык, это СТИЛЬ ЖИЗНИ! "
	var message = "Код — это как шутка. Если ты его объясняешь, значит, он плохой!"

	fmt.Println(strings.TrimSpace(message))
	fmt.Println(strings.TrimSpace(strings.ToLower(message)))
	fmt.Println(strings.Contains(message, "Go"))
}
