package authorModel

import category "web/internal/category/model"

type Author struct {
	name     string
	age      int
	category category.Category
}

func (a Author) getName() string {

	return a.name
}

func (a *Author) setName(name string) {

	a.name = name
}

func (a *Author) setCategory(category category.Category) {

	a.category = category
}

func (a Author) getCategory() category.Category {
	return a.category
}
