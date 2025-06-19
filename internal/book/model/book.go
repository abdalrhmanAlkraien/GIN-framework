package book

import category "web/internal/category/model"

type Book struct {
	name          string
	numberOfPaper int
	category      category.Category
}

func (b Book) GetName() string {
	return b.name
}

func (b *Book) SetName(name string) {
	b.name = name
}

func (b Book) GetNumberOfPaper() int {
	return b.numberOfPaper
}

func (b *Book) SetNumberOfBook(numberOfBook int) {

	b.numberOfPaper = numberOfBook
}

func (b Book) getCategory() category.Category {

	return b.category
}

func (b *Book) setCategory(category category.Category) {

	b.category = category
}
