package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error happened:", err)
		os.Exit(1)
	}
}

// Get all books
func getBooks() (books []Book) {
	bookBytes, err := ioutil.ReadFile("./books.json")
	checkError(err)

	err = json.Unmarshal(bookBytes, &books)
	checkError(err)

	return books
}

// Save books to books.json file
func saveBooks(books []Book) error {
	// converting into bytes for writing into a file
	bookBytes, err := json.Marshal(books)
	checkError(err)

	err = ioutil.WriteFile("./books.json", bookBytes, 0644)
	return err
}

// Get all the books logic
func handleGetBooks(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	// Checking for all or id
	if !*all && *id == "" {
		fmt.Println("subcommand --all or --id needed")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// if for all return all books
	if *all {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t ImageUrl \n")

		for _, book := range books {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.Id, book.Title, book.Author, book.Price, book.ImageUrl)
		}
	}

	// if for id return the book with the id
	// throw error if book is not found
	if *id != "" {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t ImageUrl \n")
		// to check a book exist or not
		var bookExist bool
		for _, book := range books {
			bookExist = true
			if book.Id == *id {
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.Id, book.Title, book.Author, book.Price, book.ImageUrl)
			}
		}
		// if no book found with mentioned id throws error
		if !bookExist {
			fmt.Println("Book not found")
		}
	}
}

// add or update a book logic
func handleAddBook(addCmd *flag.FlagSet, id, tittle, author, price, image_url *string, addNewBook bool) {
	addCmd.Parse(os.Args[2:])

	if *id == "" || *tittle == "" || *author == "" || *price == "" || *image_url == "" {
		fmt.Println("Please provide book id, title, author, price and image URL")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	books := getBooks()
	var newBook Book
	var bookExist bool

	// Checking for add or update
	if addNewBook {
		newBook = Book{*id, *tittle, *author, *price, *image_url}
		books = append(books, newBook)
	} else {
		for i, book := range books {
			if book.Id == *id {
				bookExist = true
				books[i] = Book{*id, *tittle, *author, *price, *image_url}
			}
		}

		// if no book found with mentioned id throws an error
		if !bookExist {
			fmt.Println("Book not found")
			os.Exit(1)
		}
	}

	err := saveBooks(books)
	checkError(err)

	fmt.Println("Book added successfully")
}

func handleDeleteBook(deleteCmd *flag.FlagSet, id *string) {
	deleteCmd.Parse(os.Args[2:])

	if *id == "" {
		fmt.Println("Please provide book --id")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}

	books := getBooks()
	var bookExist bool

	for i, book := range books {
		if book.Id == *id {
			bookExist = true
			books = append(books[:i], books[i+1:]...)
		}
	}

	// if no book found with mentioned id throws an error
	if !bookExist {
		fmt.Println("Book not found")
		os.Exit(1)
	}

	err := saveBooks(books)
	checkError(err)

	fmt.Println("Book deleted successfully")
}
