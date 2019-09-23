package grammarexercise

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

var once sync.Once
var client *translate.Client
var ctx = context.Background()

func initializeClient(ctx context.Context) {
	var err error
	client, err = translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

// TranslateText use Google translate API to translate a string
func TranslateText(text string, fromLanguage string, toLanguage string) string {
	// Creates a client, only once
	once.Do(func() { initializeClient(ctx) })

	if client == nil {
		log.Fatalf("No client initialized")
	}

	// Sets the target language.
	target, err := language.Parse(toLanguage)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Sets the source language.
	source, err := language.Parse(fromLanguage)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{text}, target, &translate.Options{Source: source})
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	return translations[0].Text
}
