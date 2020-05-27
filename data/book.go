package data

import "fmt"

// Book keeps information about a published book.
type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n",
		b.Title, b.Author, b.YearPublished,
	)
}

// Books contains our in-memory books database.
var Books = []Book{
	Book{
		ID:            1,
		Title:         "The Hitchhiker's Guide To The Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 2001,
	},
	Book{
		ID:            2,
		Title:         "Sapiens: A Brief History of Humankind",
		Author:        "Yuval Noah Harari",
		YearPublished: 2015,
	},
	Book{
		ID:            3,
		Title:         "Organisational Mastery: The product development blueprint for executive leaders",
		Author:        "Luis Gonçalves",
		YearPublished: 2019,
	},
}
