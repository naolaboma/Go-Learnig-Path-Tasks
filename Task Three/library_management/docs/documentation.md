# Library Management System Documentation

## Overview
This is a console-based library management system implemented in Go. It allows users to manage books and track borrowing status.

## Features
1. Add new books to the library
2. Remove books from the library
3. Borrow books (changes status to "Borrowed")
4. Return books (changes status back to "Available")
5. List all available books
6. List all books borrowed by a specific member

## Data Structures
- **Book**: Contains ID, Title, Author, and Status fields
- **Member**: Contains ID, Name, and a slice of borrowed Books

## Error Handling
The system handles several error cases:
- Trying to borrow a non-existent book
- Trying to borrow an already borrowed book
- Trying to return a book that wasn't borrowed by the member
- Trying to return a book to a non-existent member

## Usage
Run the program and follow the on-screen menu options to perform various operations.