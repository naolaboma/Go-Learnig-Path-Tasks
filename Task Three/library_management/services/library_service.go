package services

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author  string
	Status string
}

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}

type Library struct {
	books   map[int]Book
	members map[int]Member
}

func NewLibraryService() *Library {
	return &Library{
		books:   make(map[int]Book),
		members: make(map[int]Member),
	}
}

func (l *Library) AddBook(book Book) {
	book.Status = "Available"
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status != "Available" {
		return errors.New("book is not available")
	}

	member, exists := l.members[memberID]
	if !exists {
		// Create new member if not exists
		member = Member{ID: memberID, Name: fmt.Sprintf("Member %d", memberID)}
	}

	book.Status = "Borrowed"
	l.books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	// Find and remove the book from member's borrowed books
	found := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("this member hasn't borrowed this book")
	}

	book.Status = "Available"
	l.books[bookID] = book
	l.members[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []Book {
	var availableBooks []Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []Book {
	member, exists := l.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}