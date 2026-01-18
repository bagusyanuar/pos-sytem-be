package postgres

import (
	"context"

	"github.com/bagusyanuar/pos-sytem-be/internal/domain/repository"
	"gorm.io/gorm"
)

type pgUserRepo struct {
	DB *gorm.DB
}

// Find implements repository.UserRepository.
func (p *pgUserRepo) Find(ctx context.Context) error {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &pgUserRepo{
		DB: db,
	}
}
