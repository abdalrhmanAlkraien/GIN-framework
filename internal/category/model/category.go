package category

type Category struct {
	name         string
	numberOfBook int
}

func (c *Category) setName(name string) {

	c.name = name
}

func (c Category) getName() string {

	return c.name
}

func (c *Category) setNumberOfBook() {

	c.numberOfBook += 1
}

func (c Category) getNumberOfBook() int {
	return c.numberOfBook
}
