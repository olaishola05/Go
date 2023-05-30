package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOutTime time.Time
	checkInTime  time.Time
}

type BookEntry struct {
	totalCopies int
	totalLended int
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string

		if audit.checkInTime.IsZero() {
			returnTime = "Not returned yet"
		} else {
			returnTime = audit.checkInTime.String()
		}

		fmt.Println(member.name, ":", title, "at", audit.checkOutTime.Local().String(), "and", returnTime)
	}
}

func printMembersAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooksInfo(library *Library) {
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.totalCopies, "copies, ", book.totalLended, "lended")
	}
	fmt.Println()
}

func checkedOutBook(library *Library, member *Member, title Title) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not found or not part of the library")
		return false
	}

	if book.totalCopies == book.totalLended {
		fmt.Println("All copies of this book are lended")
		return false
	}

	book.totalLended++
	library.books[title] = book

	member.books[title] = LendAudit{
		checkOutTime: time.Now(),
	}

	return true
}

func returnedBook(library *Library, member *Member, title Title) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not found or not part of the library")
		return false
	}

	if book.totalLended == 0 {
		fmt.Println("No copies of this book are lended")
		return false
	}

	audit, found := member.books[title]

	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.totalLended--
	library.books[title] = book

	audit.checkInTime = time.Now()
	member.books[title] = audit

	return true
}

func main() {
	library := Library{
		members: make(map[Name]Member),
		books:   make(map[Title]BookEntry),
	}

	library.books["The Lord of the Rings"] = BookEntry{
		totalCopies: 3,
		totalLended: 0,
	}

	library.books["The Hobbit"] = BookEntry{
		totalCopies: 2,
		totalLended: 0,
	}

	library.books["Person of interest"] = BookEntry{
		totalCopies: 100,
		totalLended: 0,
	}

	library.books["The Richest Man in Babylon"] = BookEntry{
		totalCopies: 10,
		totalLended: 0,
	}

	library.members["John"] = Member{name: "John", books: make(map[Title]LendAudit)}
	library.members["Mary"] = Member{name: "Mary", books: make(map[Title]LendAudit)}
	library.members["Peter"] = Member{name: "Peter", books: make(map[Title]LendAudit)}

	fmt.Println("Initial Library books info:")
	printLibraryBooksInfo(&library)
	printMembersAudit(&library)

	member := library.members["John"]
	checkout := checkedOutBook(&library, &member, "The Lord of the Rings")
	fmt.Println("\nCheck out a book:")

	if checkout {
		printLibraryBooksInfo(&library)
		printMembersAudit(&library)
	}

	fmt.Println("\nReturn a book:")

	returned := returnedBook(&library, &member, "The Lord of the Rings")

	if returned {
		printLibraryBooksInfo(&library)
		printMembersAudit(&library)
	}

}
