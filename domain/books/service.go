package books

// BookService defines book service behavior.
type BookService interface {
	CreateBook(*Book) (*Book, error)
	ReadBook(int, int) (*Book, error)
	ListBooks(int) ([]Book, error)
}

// Service struct handles book business logic tasks.
type Service struct {
	repository BookService
}

func (svc *Service) CreateBook(book *Book) (*Book, error) {
	return svc.repository.CreateBook(book)
}

func (svc *Service) ReadBook(id int, authorId int) (*Book, error) {
	return svc.repository.ReadBook(id, authorId)
}

func (svc *Service) ListBooks(authorId int) ([]Book, error) {
	return svc.repository.ListBooks(authorId)
}

// NewService creates a new service struct
func NewService(repository BookRepository) *Service {
	return &Service{repository: repository}
}
