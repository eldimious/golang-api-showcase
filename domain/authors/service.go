package authors

// AuthorService defines author service behavior.
type AuthorService interface {
	CreateAuthor(*Author) (*Author, error)
	ReadAuthor(id int) (*Author, error)
	ListAuthors() ([]Author, error)
}

// Service struct handles author business logic tasks.
type Service struct {
	repository AuthorRepository
}

func (svc *Service) CreateAuthor(author *Author) (*Author, error) {
	return svc.repository.CreateAuthor(author)
}

func (svc *Service) ReadAuthor(id int) (*Author, error) {
	return svc.repository.ReadAuthor(id)
}

func (svc *Service) ListAuthors() ([]Author, error) {
	return svc.repository.ListAuthors()
}

// NewService creates a new service struct
func NewService(repository AuthorRepository) *Service {
	return &Service{repository: repository}
}
