package nlp

import (
	"encoding/json"

	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

type EmotionResponse struct {
	Targets  []Targets `json:"targets"`
	Document Document  `json:"document"`
}

type Targets struct {
	Text    string   `json:"text"`
	Emotion Emotions `json:"emotion"`
}
type Document struct {
	Emotion Emotions `json:"emotion"`
}

// EmotionAnalysis function via IBM NLU
func EmotionAnalysis(text string) map[string]Emotions {
	naturalLanguageUnderstanding := VerifyNLU()
	// targets := []string{"apples", "oranges"}
	keywordsList := CollectKeywords(text)
	response, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			Text: &text,
			Features: &naturallanguageunderstandingv1.Features{
				Emotion: &naturallanguageunderstandingv1.EmotionOptions{
					Targets: keywordsList,
				},
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	result := naturalLanguageUnderstanding.GetAnalyzeResult(response)
	b, _ := json.MarshalIndent(result, "", "   ")
	unwrappedJSON := UnwrapJSON(b, []string{"emotion"})
	emotionsMap := createEmotionStruct(unwrappedJSON)

	return emotionsMap
}

func createEmotionStruct(data []byte) map[string]Emotions {
	var emotionResponse EmotionResponse
	var emotionsMap = make(map[string]Emotions)
	if err := json.Unmarshal(data, &emotionResponse); err != nil {
		panic(err)
	}
	emotionResponseTargets := emotionResponse.Targets
	emotionResponseDoc := emotionResponse.Document
	emotionsMap["document"] = emotionResponseDoc.Emotion

	for _, target := range emotionResponseTargets {
		emotionsMap[target.Text] = target.Emotion
	}

	return emotionsMap
}
