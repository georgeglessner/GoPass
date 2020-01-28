package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// Password must be at least 15 chars (up to 20 right now)
	// Must contain uppercase, lowercase, number, and special char

	// Values to work with
	var lower = "abcdefghijklmnopqrstuvwxyz"
	var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const special = "!@#$%&"

	// https://golang.org/pkg/math/rand/
	// Use the Seed function to initialize the default Source
	// if different behavior is required for each run.
	rand.Seed(time.Now().UnixNano())

	// Generate random number
	n := strconv.Itoa(rand.Intn(100))

	// Initialize password letter string
	s := ""

	// Generate random string of letters
	for i := range lower {
		if i < 12 {
			r := rand.Intn(25)
			if r%2 == 0 {
				// pick random lowercase
				r := rand.Intn(25)
				s += lower[r : r+1]
			} else {
				// pick random uppercase
				s += upper[r : r+1]
			}
		}
	}

	// Put it all together
	var p = s + special[rand.Intn(len(special)):] + n

	// Display results
	fmt.Println("New Password:", p)
	fmt.Println("Length:", len(p))

}
