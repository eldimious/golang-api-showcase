package authors

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	domain "github.com/eldimious/golang-api-showcase/domain/authors"
	booksDomain "github.com/eldimious/golang-api-showcase/domain/books"
	domainErrors "github.com/eldimious/golang-api-showcase/domain/errors"
)

const (
	createError = "Error in creating new author"
	readError   = "Error in finding author in the database"
	listError   = "Error in getting authors from the database"
)

// Store struct manages interactions with authors store
type Store struct {
	db        *gorm.DB
	booksRepo booksDomain.BookRepository
}

// New creates a new Store struct
func New(db *gorm.DB, booksRepo booksDomain.BookRepository) *Store {
	db.AutoMigrate(&Author{})

	return &Store{
		db:        db,
		booksRepo: booksRepo,
	}
}

func (s *Store) CreateAuthor(author *domain.Author) (*domain.Author, error) {
	entity := toDBModel(author)

	if err := s.db.Create(entity).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(entity), nil
}

func (s *Store) ReadAuthor(id int) (*domain.Author, error) {
	result := &Author{}

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

func (s *Store) ListAuthors() ([]domain.Author, error) {
	var results []Author

	if err := s.db.Find(&results).Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
		return nil, appErr
	}

	var authors = make([]domain.Author, len(results))

	for i, element := range results {
		authors[i] = *toDomainModel(&element)
	}

	return authors, nil
}
