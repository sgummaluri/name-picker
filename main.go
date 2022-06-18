package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Enter the list of names. Key in '--done' when you are done.")
	var names_slice []string

	reader := bufio.NewReader(os.Stdin)

	// Seed the random number generator - https://pkg.go.dev/math/rand#Seed
	rand.Seed(time.Now().UnixNano())

	name, err := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\r\n")

	for name != "--done" && err == nil {
		names_slice = append(names_slice, name)
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error reading input")
			return
		}
		name = strings.TrimSuffix(name, "\r\n")
	}

	randomIndex := rand.Intn(len(names_slice))
	fmt.Printf("The chosen one is %s", names_slice[randomIndex])
}
