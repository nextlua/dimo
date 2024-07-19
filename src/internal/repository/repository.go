package repository

type Repository struct {
	DB any
}

func NewRepository(db any) *Repository {
	return &Repository{DB: db}
}
