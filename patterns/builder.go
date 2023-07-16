package patterns

import (
	"fmt"
)

type Scitation interface {
	Book(book string) Scitation
	Page(page int) Scitation
	Text(text string) Scitation
	Get() string
}

type builder struct {
	book string
	page int
	text string
}

func (b *builder) Book(book string) Scitation {
	b.book = book
	return b
}

func (b *builder) Page(page int) Scitation {
	b.page = page
	return b
}

func (b *builder) Text(text string) Scitation {
	b.text = text
	return b
}

func (b *builder) Get() string {
	return fmt.Sprintf("%s %d %s", b.book, b.page, b.text)
}

func NewScitation() Scitation {
	return &builder{
		book: "",
		page: 0,
		text: "",
	}
}