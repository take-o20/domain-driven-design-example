package domain

type Author struct {
	id   string
	name string
}

func (author *Author) ChangeName(newName string) {
	author.name = newName
}

type AuthorID string

func (*AuthorID) Generate() string {
	return ""
}

func NewAuthor() Author {
	return Author{}
}
