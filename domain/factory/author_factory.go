package factory

import "github.com/take-o20/domain-driven-design-example/domain"

type AuthorFactory interface {
	NewAuthor() domain.Author
}
