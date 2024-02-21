package mysql

import (
	"context"
	"go-clean-architecture-translate/services"
	"go-clean-architecture-translate/services/entities"
	"gorm.io/gorm"
)

const tbName = "translations"

type mysqlRepo struct {
	db *gorm.DB
}

func (m mysqlRepo) GetTranslation(ctx context.Context, originalText string, source string, destination string) (entities.Translation, error) {
	var data entities.Translation
	if err := m.db.Table(tbName).Where("original_text = ? AND source = ? AND destination = ?", originalText, source, destination).First(&data).Error; err != nil {
		return entities.Translation{}, err
	}
	return data, nil
}

func (m mysqlRepo) FetchHistories(ctx context.Context) ([]entities.Translation, error) {
	var histories []entities.Translation
	if err := m.db.Table(tbName).Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

func (m mysqlRepo) InsertTranslation(ctx context.Context, translation entities.Translation) error {
	return m.db.Create(translation).Error
}

func NewMySQLRepo(db *gorm.DB) services.ITranslateRepository {
	return mysqlRepo{db: db}
}
