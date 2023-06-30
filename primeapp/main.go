package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	intro()
	done := make(chan bool)
	go readUserInput(os.Stdin, done)
	<-done
	close(done)
}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	//check the see if user wants to quit
	scanner.Scan()
	if scanner.Text() == "q" {
		return "", true
	}

	//convert type to int
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Enter a whole number", false
	}

	_, msg := isPrime(i)

	return msg, false
}

func readUserInput(in io.Reader, doneCh chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		msg, done := checkNumber(scanner)
		if done {
			doneCh <- true
			return
		}
		fmt.Println(msg)
		prompt()
	}
}

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("Enter a number, Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("> ")
}

func isPrime(num int) (bool, string) {
	// 0 and 1 are not prime in definition
	if num == 0 || num == 1 {
		return false, fmt.Sprintf("number (%d) is not prime", num)
	}

	// Negative numbers are not prime in definition
	if num < 0 {
		return false, fmt.Sprintf("negative number (%d) is not prime", num)
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false, fmt.Sprintf("number (%d) is not prime, its divisible by (%d)", num, i)
		}
	}
	return true, fmt.Sprintf("number (%d) is prime!!", num)
}
