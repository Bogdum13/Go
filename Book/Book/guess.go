// guess is a game in which the player has to guess a random number.

package main

// Page 89

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	// Generate a random number of minutes between 1 and 100
	target := rng.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	//Create a "buffered reader" to receive text from the keyboard.
	reader := bufio.NewReader(os.Stdin)

	// Declaring the "success" variable before the loop so that it remains in scope after exiting the loop
	success := false

	//The "guesses" variable stores the number of remaining attempts
	for guesses := 10; guesses > 0; guesses-- {
		fmt.Println("You have", guesses, "guess left.")

		// Request a value from the user.
		fmt.Print("Make a guess: ")
		// Returns the text entered by the user before pressing the Enter key. Будет прочитан весь текст до руны новой строки.
		//input, _ := reader.ReadString('\n') // An empty identifier is used as a "placeholder" for the error code
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err) //Report an error and stop the program.
		}

		// Removing the newline character from the input string
		input = strings.TrimSpace(input)
		//The arguments pass the converted string "input" and the number in integer format
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			// If the player's guess is less than the hidden number, report it
			fmt.Println("Oops. Your guess was LOW\n\n")
		} else if guess > target {
			// If the player's guess is more than the hidden number, report it
			fmt.Println("Oops. Your guess was HIGH\n\n")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	// If the player has NOT guessed correctly (if the "success" variable contains false), display a loss message
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}
