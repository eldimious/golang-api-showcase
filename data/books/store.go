package books

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	authorSchema "github.com/eldimious/golang-api-showcase/data/authors"
	domain "github.com/eldimious/golang-api-showcase/domain/books"
	domainErrors "github.com/eldimious/golang-api-showcase/domain/errors"
)

const (
	createError = "Error in creating new book"
	readError   = "Error in finding book in the database"
	listError   = "Error in getting books from the database"
)

// Store struct manages interactions with books store
type Store struct {
	db *gorm.DB
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	db.AutoMigrate(&Book{})

	return &Store{
		db: db,
	}
}

// AddAuthorFKConstraint creates a foreign key constraint for the author Id field
func (s *Store) AddAuthorFKConstraint() {
	s.db.Model(&Book{}).AddForeignKey("author_id", "authors(id)", "CASCADE", "CASCADE")
}

func (s *Store) CreateBook(book *domain.Book) (*domain.Book, error) {
	entity := toDBModel(book)

	if err := s.db.Create(entity).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(entity), nil
}

func (s *Store) ReadBook(id int) (*domain.Book, error) {
	result := &Book{}

	query := s.db.Where("id = ?", id).First(result)

	if query.RecordNotFound() {
		appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appErr
	}

	if err := query.Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(result), nil
}

func (s *Store) ListBooks(authorId int) ([]domain.Book, error) {
	var results []Book
	author := &authorSchema.Author{
		Id: authorId,
	}

	if err := s.db.Model(&author).Related(&results).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
		return nil, appErr
	}

	var books = make([]domain.Book, len(results))

	for i, element := range results {
		books[i] = *toDomainModel(&element)
	}

	return books, nil
}
