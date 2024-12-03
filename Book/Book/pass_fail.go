// Page 69
// pass_fail notifies whether the user has passed the exam
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Request a value from the user.
	fmt.Print("Enter a grade: ")
	//Create a "buffered reader" to receive text from the keyboard.
	reader := bufio.NewReader(os.Stdin)
	// Returns the text entered by the user before pressing the Enter key. Будет прочитан весь текст до руны новой строки.
	//input, _ := reader.ReadString('\n') // An empty identifier is used as a "placeholder" for the error code
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err) //Report an error and stop the program.
	}

	// Removing the newline character from the input string
	input = strings.TrimSpace(input)
	//The arguments pass the converted string "input" and the number of precision bits for the result is 64.
	// (количество битов точности для результата 64). And the return value is "grade" (float64)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err) //Report an error and stop the program.
	}
   
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
