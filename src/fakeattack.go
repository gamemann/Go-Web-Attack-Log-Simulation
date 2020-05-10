package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// The file name we will be opening.
	fileName := "fake_attack.log"

	// Check for command line file name.
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	// Open file for appending.
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	// Create safe hosts array.
	safeHosts := []string{"92.119.148.4", "92.119.148.10", "92.119.148.15", "167.50.90.2", "127.0.0.1", "192.168.50.4", "192.168.50.6"}

	// Create malicious hosts array.
	maliciousHosts := []string{"84.10.30.4", "76.50.40.3", "123.12.40.9", "232.5.100.200", "157.90.1.2", "44.32.20.2", "90.5.253.54", "3.9.8.2", "60.6.32.19", "9.80.30.4"}

	// Count, limit, and timeout variables \o/.
	var count int64 = 0
	var limit int64 = 3000
	var timeout int64 = 500

	// Check for command line limit.
	if len(os.Args) > 2 {
		limit, err = strconv.ParseInt(os.Args[2], 10, 64)

		if err != nil {
			log.Fatal(err)
		}
	}

	// Check for command line timeout.
	if len(os.Args) > 3 {
		timeout, err = strconv.ParseInt(os.Args[3], 10, 64)

		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		// If count is higher than limit, deaconate program.
		if count > limit {
			// Fuck and you.
			fmt.Print("Ending program. Count limit reached.\n")

			// BREAK DIS BITCH.
			break
		}

		// Get random number.
		chance := rand.Intn(10)

		// Select which IP with "0.0.0.0" being default (though, it should never get this since we pick from the pool, xd).
		ip := "0.0.0.0"

		if chance > 1 {
			// Malicious
			ip = maliciousHosts[rand.Intn(len(maliciousHosts))]
		} else {
			// Safe Host.
			ip = safeHosts[rand.Intn(len(safeHosts))]
		}

		// Let's write to the file.
		if _, err := file.Write([]byte(ip + " /index.php GET Chrome/V7.6.3.4\n")); err != nil {
			log.Fatal(err)
		}

		// Increase count.
		count++

		// Print dat out.
		fmt.Print("Writing to file with host \"", ip, "\". Count ", count, "\n")

		// Wait 500 msecs.
		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}

	// Close the file.
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
