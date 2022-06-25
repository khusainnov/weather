package repository

type Weather interface {
}

type Repository struct {
	Weather
}

func NewRepository() *Repository {
	return &Repository{}
}
