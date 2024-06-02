package usecase

import (
	"github.com/take-o20/domain-driven-design-example/domain"
	"github.com/take-o20/domain-driven-design-example/domain/factory"
	"github.com/take-o20/domain-driven-design-example/domain/repository"
)

type AuthorService struct {
	authorFactory    factory.AuthorFactory
	authorRepository repository.AuthorRepository
}

func NewAuthorService() AuthorService {
	return AuthorService{}
}

func (as *AuthorService) CreateAuthor() {
	author := as.authorFactory.NewAuthor()
	author.ChangeName("hogehoge")
	as.authorRepository.Save(author)
}

func (as *AuthorService) FindAuthor(id string) domain.Author {
	return domain.Author{}
}
