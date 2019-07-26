package main

import (
	"fmt"

	"github.com/rafitj/sentiment-gradient-trees/nlp"
)

func main() {
	text := "Hey what's up mom, I'm having a pretty bad day today. I don't know why I keep getting bullied I just hate school so much."
	fmt.Println(text)
	response := nlp.EmotionAnalysis(text)
	fmt.Println(response)
}
