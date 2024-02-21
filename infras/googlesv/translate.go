package googlesv

import (
	"context"
	translator "github.com/Conight/go-googletrans"
	"go-clean-architecture-translate/services"
	"go-clean-architecture-translate/services/entities"
)

type googleTranslateAPI struct {
}

func (g googleTranslateAPI) Translate(ctx context.Context, originalText string, source string, destination string) (entities.Translation, error) {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	trans := translator.New(conf)

	result, err := trans.Translate(originalText, source, destination)
	if err != nil {
		return entities.Translation{}, err
	}

	return entities.NewTranslation(originalText, source, destination, result.Text), nil
}

func NewGoogleTranslateAPI() services.IGoogleService {
	return googleTranslateAPI{}
}
