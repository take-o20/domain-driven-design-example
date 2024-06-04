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

func (as *AuthorService) GetAuthor(id string) (domain.Author, error) {
	return as.authorRepository.SearchAuthor(id)
}

func (as *AuthorService) DeleteAuthor(id string) error {
	return nil
}

func (as *AuthorService) UpdateAuthorName(id string, newAuthorName string) error {
	targetAuthor, err := as.GetAuthor(id)
	if err != nil {
		return err
	}
	targetAuthor.ChangeName(newAuthorName)
	return as.authorRepository.Save(targetAuthor)
}
