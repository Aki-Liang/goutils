package main

import (
	"fmt"

	"github.com/Aki-Liang/goutils/csprng"
)

func main() {
	// Example: this will give us a 44 byte, base64 encoded output
	token, err := csprng.GenerateRandomString(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
	}
	fmt.Println(token)
}
