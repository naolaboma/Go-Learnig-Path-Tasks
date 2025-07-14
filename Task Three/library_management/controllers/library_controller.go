package controllers

import (
	"bufio"
	"fmt"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	service services.LibraryManager
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		service: services.NewLibraryService(),
	}
}

func (lc *LibraryController) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			lc.addBook(reader)
		case 2:
			lc.removeBook(reader)
		case 3:
			lc.borrowBook(reader)
		case 4:
			lc.returnBook(reader)
		case 5:
			lc.listAvailableBooks()
		case 6:
			lc.listBorrowedBooks(reader)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (lc *LibraryController) addBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID: ")
	idInput, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idInput))

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := services.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	lc.service.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (lc *LibraryController) removeBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to remove: ")
	idInput, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idInput))

	lc.service.RemoveBook(id)
	fmt.Println("Book removed successfully!")
}

func (lc *LibraryController) borrowBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to borrow: ")
	bookIDInput, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDInput))

	fmt.Print("Enter member ID: ")
	memberIDInput, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

	err := lc.service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (lc *LibraryController) returnBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to return: ")
	bookIDInput, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDInput))

	fmt.Print("Enter member ID: ")
	memberIDInput, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

	err := lc.service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (lc *LibraryController) listAvailableBooks() {
	books := lc.service.ListAvailableBooks()
	fmt.Println("\nAvailable Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) listBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter member ID: ")
	memberIDInput, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

	books := lc.service.ListBorrowedBooks(memberID)
	fmt.Printf("\nBooks borrowed by member %d:\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
