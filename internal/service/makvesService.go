package service

import (
	"context"

	"github.com/andy-ahmedov/makves/internal/domain"
)

type MakvesRepository interface {
	Get(ctx context.Context, ids []int64) ([]domain.Makves, error)
}

type MakvesService struct {
	repo MakvesRepository
}

func NewMakvesService(repo MakvesRepository) *MakvesService {
	return &MakvesService{
		repo: repo,
	}
}

func (t *MakvesService) GetData(ctx context.Context, ids []int64) ([]domain.Makves, error) {
	return t.repo.Get(ctx, ids)
}
