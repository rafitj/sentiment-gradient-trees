package main

import (
	"fmt"

	"github.com/rafitj/sentiment-gradient-trees/nlp"
)

func main() {
	response := nlp.EmotionAnalysis("Hey what's up Alexa, I'm have a pretty bad day.")
	fmt.Println(response)
}
