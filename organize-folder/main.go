package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Check for any error
func check(err error) {
	if err != nil {
		fmt.Printf("Error happened: %s \n", err)
		os.Exit(1)
	}
}

// function to create default folders such as Images, Music, Docs, Others, Videos
func createDefaultFolders(targetFolder string) {
	defaultFolders := []string{"Docs", "Images", "Music", "Videos", "Others"}

	for _, folder := range defaultFolders {
		_, err := os.Stat(filepath.Join(targetFolder, folder))
		if os.IsNotExist(err) {
			os.Mkdir(filepath.Join(targetFolder, folder), 0755)
		}
	}
}

// function to Organize folders
func organizeFolders(targetFolder string) {
	filesAndFolders, err := os.ReadDir(targetFolder)
	check(err)

	// to track how many files moved
	noOfFiles := 0

	for _, fileAndFolder := range filesAndFolders {
		// Check for files
		if !fileAndFolder.IsDir() {
			fileInfo, err := fileAndFolder.Info()
			check(err)

			// Get the file full path
			oldPath := filepath.Join(targetFolder, fileInfo.Name())
			fileExt := filepath.Ext(oldPath)

			switch fileExt {
			case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
				newPath := filepath.Join(targetFolder, "Images", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp3", ".wav", ".flac", ".aac", ".ogg":
				newPath := filepath.Join(targetFolder, "Music", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".doc", ".docx", ".txt", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx":
				newPath := filepath.Join(targetFolder, "Docs", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp4", ".mkv", ".avi", ".flv", ".mov", ".wmv", ".3gp", ".m4v":
				newPath := filepath.Join(targetFolder, "Videos", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			default:
				newPath := filepath.Join(targetFolder, "Others", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			}
		}
	}

	// Print how many files moved
	if noOfFiles > 0 {
		fmt.Printf("%v number of files moved \n", noOfFiles)
	} else {
		fmt.Println("No files moved")
	}
}

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
