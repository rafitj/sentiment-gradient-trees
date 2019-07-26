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
	emotions        = []string{"anger", "disgust", "fear", "joy", "sadness"}
)

// Emotion struct definition
type Emotion struct {
	anger, disgust, fear, joy, sadness string
}

func verifyNLU() *naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1 {
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

// EntityIdentification function via IBM NLU
func EntityIdentification(text string) string {
	naturalLanguageUnderstanding := verifyNLU()

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
	return string(b)
}

// EmotionAnalysis function via IBM NLU
func EmotionAnalysis(text string) Emotion {
	naturalLanguageUnderstanding := verifyNLU()
	// targets := []string{"apples", "oranges"}

	response, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			Text: &text,
			Features: &naturallanguageunderstandingv1.Features{
				Emotion: &naturallanguageunderstandingv1.EmotionOptions{
					// Targets: targets,
				},
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	result := naturalLanguageUnderstanding.GetAnalyzeResult(response)
	b, _ := json.MarshalIndent(result, "", "   ")

	fields := []string{"emotion", "document", "emotion"}
	unwrapped := unWrapJSON(b, fields)
	emotion := createEmotionStruct(unwrapped, emotions)

	return emotion
}

func unWrapJSON(data []byte, fields []string) []byte {
	for _, field := range fields {
		var rawJSON map[string]interface{}
		json.Unmarshal(data, &rawJSON)
		res := rawJSON[field]
		data, _ = json.MarshalIndent(res, "", "   ")
	}
	return data
}

func createEmotionStruct(data []byte, emotions []string) Emotion {
	var list []string
	for _, emotion := range emotions {
		var rawJSON map[string]interface{}
		json.Unmarshal(data, &rawJSON)
		res := rawJSON[emotion]
		var data, _ = json.MarshalIndent(res, "", "   ")
		list = append(list, string(data))
	}
	return Emotion{
		anger:   list[0],
		disgust: list[1],
		fear:    list[2],
		joy:     list[3],
		sadness: list[4],
	}
}
