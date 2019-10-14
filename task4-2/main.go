package main

import (
	"fmt"
	"sort"
)

type contact struct {
	name   string
	number int
}

type byNumber []*contact

func (c byNumber) Len() int           { return len(c) }
func (c byNumber) Less(i, j int) bool { return c[i].number < c[j].number }
func (c byNumber) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

type byName []*contact

func (c byName) Len() int           { return len(c) }
func (c byName) Less(i, j int) bool { return c[i].name < c[j].name }
func (c byName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func main() {
	contacts := []*contact{
		&contact{"Alex", 123},
		&contact{"Wtf", 42},
		&contact{"Bob", 321},
	}

	sort.Sort(byNumber(contacts))
	displayContacts("by number:", contacts)

	sort.Sort(byName(contacts))
	displayContacts("by name:", contacts)
}

func displayContacts(header string, contacts []*contact) {
	fmt.Println(header)
	for i, c := range contacts {
		fmt.Printf("  %d. %s %d\n", i+1, c.name, c.number)
	}
	fmt.Println()
}
