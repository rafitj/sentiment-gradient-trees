package nlp

import (
	"encoding/json"

	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

type ConceptResponse struct {
	Concepts []Concepts `json:"concepts"`
}
type Concepts struct {
	Text            string  `json:"text"`
	Relevance       float64 `json:"relevance"`
	DbpediaResource string  `json:"dbpedia_resource"`
}

// ConceptIdentification function via IBM NLU
func ConceptIdentification(text string) []string {
	naturalLanguageUnderstanding := VerifyNLU()

	limit := int64(1)

	response, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			Text: &text,
			Features: &naturallanguageunderstandingv1.Features{
				Concepts: &naturallanguageunderstandingv1.ConceptsOptions{
					Limit: &limit,
				},
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	result := naturalLanguageUnderstanding.GetAnalyzeResult(response)
	b, _ := json.MarshalIndent(result, "", "   ")
	conceptsList := createConceptsList(b)
	return conceptsList
}

func createConceptsList(data []byte) []string {
	var concepts ConceptResponse
	var list []string
	if err := json.Unmarshal(data, &concepts); err != nil {
		panic(err)
	}
	conceptsList := concepts.Concepts

	for _, word := range conceptsList {
		list = append(list, word.Text)
	}

	return list
}
