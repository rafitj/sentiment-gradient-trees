package nlp

import (
	"encoding/json"
	"os"

	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

var (
	IBM_NLP_URL     = os.Getenv("IBM_NLP_URL")
	IBM_NLP_VERSION = "2019-07-12"
	IAM_API_KEY     = os.Getenv("IBM_NLP_API_KEY")
	EmotionsList    = []string{"anger", "disgust", "fear", "joy", "sadness"}
)

// Sentiment Score struct
type Sentiment struct {
	Score float64 `json:"score"`
	Label string  `json:"label"`
}

// Emotion struct
type Emotions struct {
	Sadness float64 `json:"sadness"`
	Joy     float64 `json:"joy"`
	Fear    float64 `json:"fear"`
	Disgust float64 `json:"disgust"`
	Anger   float64 `json:"anger"`
}

// VerifyNLU func for IAM NLU auth
func VerifyNLU() *naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1 {
	naturalLanguageUnderstanding, naturalLanguageUnderstandingErr := naturallanguageunderstandingv1.
		NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
			URL:       IBM_NLP_URL,
			Version:   IBM_NLP_VERSION,
			IAMApiKey: IAM_API_KEY,
		})
	if naturalLanguageUnderstandingErr != nil {
		panic(naturalLanguageUnderstandingErr)
	}
	return naturalLanguageUnderstanding
}

// UnwrapJSON func for general json parse
func UnwrapJSON(data []byte, fields []string) []byte {
	for _, field := range fields {
		var rawJSON map[string]interface{}
		json.Unmarshal(data, &rawJSON)
		res := rawJSON[field]
		data, _ = json.MarshalIndent(res, "", "   ")
	}
	return data
}
