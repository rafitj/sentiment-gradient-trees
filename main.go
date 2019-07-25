package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

func main() {
	naturalLanguageUnderstanding, naturalLanguageUnderstandingErr := naturallanguageunderstandingv1.
		NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
			URL:       os.Getenv("IBM_NLP_URL"),
			Version:   "2019-07-12",
			IAMApiKey: os.Getenv("IBM_NLP_API_KEY"),
		})
	if naturalLanguageUnderstandingErr != nil {
		panic(naturalLanguageUnderstandingErr)
	}

	html := "<html><head><title>Fruits</title></head><body><h1>Apples and Oranges</h1><p>I love apples! I don't like oranges.</p></body></html>"
	targets := []string{"apples", "oranges"}

	response, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			HTML: &html,
			Features: &naturallanguageunderstandingv1.Features{
				Emotion: &naturallanguageunderstandingv1.EmotionOptions{
					Targets: targets,
				},
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	result := naturalLanguageUnderstanding.GetAnalyzeResult(response)
	b, _ := json.MarshalIndent(result, "", "   ")
	fmt.Println(string(b))
}
