package main

import (
	"bufio"
	"fmt"
	"os"
)

// Check for any error
func check(err error) {
	if err != nil {
		fmt.Printf("Error happened: %s \n", err)
		os.Exit(1)
	}
}

// function to create default folders such as Images, Music, Docs, Others, Videos
func createDefaultFolders(targetFolder string) {}

// function to Organize folders
func organizeFolders(targetFolder string) {}

func main() {
	// Get the user input - target folder needs to be organized
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Which folder do you want to organize? - ")
	scanner.Scan()

	targetFolder := scanner.Text()

	// Check the folder exists or not
	_, err := os.Stat(targetFolder)
	if os.IsNotExist(err) {
		fmt.Printf("Folder %s does not exist.\n", targetFolder)
		os.Exit(1)
	} else {
		createDefaultFolders(targetFolder)

		organizeFolders(targetFolder)
	}
}
