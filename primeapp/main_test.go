package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel, and io.Reader instance
	doneCh := make(chan bool)

	// creates a reference to the bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneCh)
	<-doneCh
	close(doneCh)
}
func Test_checkNumber(t *testing.T) {
	checkNumberTests := []struct {
		name     string
		input    io.Reader
		expected string
	}{
		{"empty", strings.NewReader(""), "Enter a whole number"},
		{"zero", strings.NewReader("0"), "number (0) is not prime"},
		{"one", strings.NewReader("1"), "number (1) is not prime"},
		{"two", strings.NewReader("2"), "number (2) is prime!!"},
		{"four", strings.NewReader("4"), "number (4) is not prime, its divisible by (2)"},
		{"negative", strings.NewReader("-1"), "negative number (-1) is not prime"},
		{"quit", strings.NewReader("q"), ""},
	}

	for _, tt := range checkNumberTests {
		reader := bufio.NewScanner(tt.input)
		res, _ := checkNumber(reader)
		if res != tt.expected {
			t.Errorf("%s: got %s, want %s", tt.name, res, tt.expected)
		}
	}
}

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		num      int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "number (7) is prime!!"},
		{"zero", 0, false, "number (0) is not prime"},
		{"one", 1, false, "number (1) is not prime"},
		{"negative", -1, false, "negative number (-1) is not prime"},
		{"not prime", 100, false, "number (100) is not prime, its divisible by (2)"},
	}

	for _, tt := range primeTests {
		result, msg := isPrime(tt.num)
		if result != tt.expected {
			t.Errorf("%s: got %t, want %t", tt.name, result, tt.expected)
		}

		if msg != tt.msg {
			t.Errorf("%s: got %s, want %s", tt.name, msg, tt.msg)
		}
	}
}
func Test_intro(t *testing.T) {
	// Save the copy of os.Stdout
	oldOut := os.Stdout

	// Create a read and write pipe
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// Set os.Stdout to the write pipe
	os.Stdout = w

	intro()

	// Close the write pipe
	_ = w.Close()

	// Reset os.Stdout to its original value
	os.Stdout = oldOut

	// Read the output from prompt func from our read pipe
	out, _ := io.ReadAll(r)

	// Perform assertions
	if !strings.Contains(string(out), "Enter a number, Enter q to quit.") {
		t.Errorf("got %s, want %s", string(out), "Enter a number, Enter q to quit.")
	}
}

func Test_prompt(t *testing.T) {
	// Save the copy of os.Stdout
	oldOut := os.Stdout

	// Create a read and write pipe
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// Set os.Stdout to the write pipe
	os.Stdout = w

	prompt()

	// Close the write pipe
	_ = w.Close()

	// Reset os.Stdout to its original value
	os.Stdout = oldOut

	// Read the output from prompt func from our read pipe
	out, _ := io.ReadAll(r)

	// Perform assertions
	if string(out) != "> " {
		t.Errorf("got %s, want %s", string(out), "> ")
	}
}
