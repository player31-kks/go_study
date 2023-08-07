package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneChan := make(chan bool)
	go readUserInput(doneChan)
	<-doneChan
	close(doneChan)

}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Invalid input", false
	}
	_, msg := isPrime(numToCheck)
	return msg, false
}

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("============")
	fmt.Println("Enter a number to check if it is prime or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("->")
}

func isPrime(n int) (bool, string) {
	// ...
	if n == 0 || n == 1 {
		return false, "0 or 1 is not prime"
	}
	if n < 0 {
		return false, "negative number is not prime"
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, "not prime"
		}
	}
	return true, "prime"
}
