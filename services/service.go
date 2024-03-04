package services

import (
	"context"
	"go-clean-architecture-translate/services/entities"
)

type service struct {
	translateRepository ITranslateRepository
	googleService       IGoogleService
}

func NewService(translateRepository ITranslateRepository, googleService IGoogleService) ITranslateUseCase {
	return service{
		translateRepository: translateRepository,
		googleService:       googleService,
	}
}

func (s service) Translate(ctx context.Context, originalText string, source string, destination string) (entities.Translation, error) {
	oldTrans, err := s.translateRepository.GetTranslation(ctx, originalText, source, destination)

	if err == nil {
		return oldTrans, nil
	}

	newTrans, err := s.googleService.Translate(ctx, originalText, source, destination)

	if err != nil {
		return entities.Translation{}, err
	}

	err = s.translateRepository.InsertTranslation(ctx, newTrans)

	if err != nil {
		return entities.Translation{}, err
	}

	return newTrans, nil
}

func (s service) FetchHistories(ctx context.Context) ([]entities.Translation, error) {
	return s.translateRepository.FetchHistories(ctx)
}
