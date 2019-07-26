package nlp

import (
	"encoding/json"

	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

type EntityResponse struct {
	Entities []Entities `json:"entities"`
}

type Disambiguation struct {
	Subtype         []string `json:"subtype"`
	Name            string   `json:"name"`
	DbpediaResource string   `json:"dbpedia_resource"`
}
type Entities struct {
	Type           string         `json:"type"`
	Text           string         `json:"text"`
	Sentiment      Sentiment      `json:"sentiment"`
	Relevance      float64        `json:"relevance"`
	Disambiguation Disambiguation `json:"disambiguation"`
	Count          int            `json:"count"`
}

// EntityIdentification function via IBM NLU
func EntityIdentification(text string) []string {
	naturalLanguageUnderstanding := VerifyNLU()

	sentiment := true
	limit := int64(1)

	response, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			Text: &text,
			Features: &naturallanguageunderstandingv1.Features{
				Entities: &naturallanguageunderstandingv1.EntitiesOptions{
					Sentiment: &sentiment,
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
	entityList := createEntityList(b)
	return entityList
}

func createEntityList(data []byte) []string {
	var entities EntityResponse
	var list []string
	if err := json.Unmarshal(data, &entities); err != nil {
		panic(err)
	}
	entitiesList := entities.Entities

	for _, word := range entitiesList {
		list = append(list, word.Disambiguation.Name)
	}

	return list
}
