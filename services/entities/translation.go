package entities

type Translation struct {
	OriginalText string `json:"original_text"`
	Source       string `json:"source"`
	Destination  string `json:"destination"`
	ResultText   string `json:"result_text"`
}

func NewTranslation(originalText string, source string, destination string, resultText string) Translation {
	return Translation{
		OriginalText: originalText,
		Source:       source,
		Destination:  destination,
		ResultText:   resultText,
	}
}

func (t *Translation) SetResultText(resultText string) {
	t.ResultText = resultText
}
