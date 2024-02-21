package services

import (
	"context"
	"go-clean-architecture-translate/services/entities"
)

type ITranslateUseCase interface {
	Translate(ctx context.Context, originalText string, source string, destination string) (entities.Translation, error)
	FetchHistories(ctx context.Context) ([]entities.Translation, error)
}

type ITranslateRepository interface {
	GetTranslation(ctx context.Context, originalText string, source string, destination string) (entities.Translation, error)
	FetchHistories(ctx context.Context) ([]entities.Translation, error)
	InsertTranslation(ctx context.Context, translation entities.Translation) error
}

type IGoogleService interface {
	Translate(ctx context.Context, originalText string, source string, destination string) (entities.Translation, error)
}
