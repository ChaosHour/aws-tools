// read from a file and print the contents and run a command against the file
package main

// import the necessary packages
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// runHost function
func runHost(host string) {
	// run the command
	cmd := exec.Command("host", host)
	// read the output
	output, err := cmd.Output()
	// if there is an error, print it
	if err != nil {
		fmt.Println(err)
	}
	// print the output
	fmt.Println(string(output))
}

// main function
func main() {
	// open the file
	file, err := os.Open("host.txt")
	// if there is an error, print it
	if err != nil {
		fmt.Println(err)
	}
	// make a scanner to read the file
	scanner := bufio.NewScanner(file)
	// loop through the file
	for scanner.Scan() {
		// get the line
		line := scanner.Text()
		// print the line
		fmt.Println(line)
		// run the host command
		runHost(line)
	}
}
