package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// Password must be at least 15 chars
	// Must contain uppercase, lowercase, number, and special char

	// Values to work with
	const lower = "abcdefghijklmnopqrstuvwxyz"
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const special = "!@#$%&"

	// https://golang.org/pkg/math/rand/
	// Use the Seed function to initialize the default Source
	// if different behavior is required for each run.
	rand.Seed(time.Now().UnixNano())

	// Generate random number for password
	n := strconv.Itoa(rand.Intn(100))

	// Initialize password string
	p := ""

	// Assign values to determine where to place
	// number and special char in the password
	nr := rand.Intn(11)
	sr := rand.Intn(11)

	// Generate password
	// 13 is used to ensure min. of 15 chars
	for i := 0; i < 13; i++ {
		r := rand.Intn(25)

		if i == nr {
			// Insert number
			p += n
		}
		if i == sr {
			// Insert special char
			p += special[rand.Intn(len(special)):]

		}
		if r%2 == 0 {
			// Insert random lowercase
			r := rand.Intn(25)
			p += lower[r : r+1]
		} else {
			// Insert random uppercase
			p += upper[r : r+1]
		}
	}

	// Display results
	fmt.Println("New Password:", p)
	fmt.Println("Length:", len(p))

}
