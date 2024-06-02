package repository

import "github.com/take-o20/domain-driven-design-example/domain"

type AuthorRepository interface {
	Save(domain.Author)
	Delete(domain.Author)
}
