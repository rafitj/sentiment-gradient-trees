package nlp

import (
	"encoding/json"

	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

type KeywordsResponse struct {
	Keywords []Keywords `json:"keywords"`
}

type Keywords struct {
	Text      string    `json:"text"`
	Sentiment Sentiment `json:"sentiment"`
	Relevance float64   `json:"relevance"`
	Emotions  Emotions  `json:"emotions"`
}
type KeywordsList struct {
	Text []string
}

// CollectKeywords function via IBM NLU
func CollectKeywords(text string) []string {
	naturalLanguageUnderstanding := VerifyNLU()

	sentiment := true
	emotion := true
	limit := int64(3)

	response, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			Text: &text,
			Features: &naturallanguageunderstandingv1.Features{
				Keywords: &naturallanguageunderstandingv1.KeywordsOptions{
					Sentiment: &sentiment,
					Emotion:   &emotion,
					Limit:     &limit,
				},
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	result := naturalLanguageUnderstanding.GetAnalyzeResult(response)
	b, _ := json.MarshalIndent(result, "", "   ")
	keywordsList := createKeywordsList(b)
	return (keywordsList)
}

func createKeywordsList(data []byte) []string {
	var keywords KeywordsResponse
	var list []string
	if err := json.Unmarshal(data, &keywords); err != nil {
		panic(err)
	}
	keywordsList := keywords.Keywords

	for _, word := range keywordsList {
		list = append(list, word.Text)
	}

	return list
}
