package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	response, err := llms.GenerateFromSinglePrompt(ctx, llm, "Hello! How can you help me today?")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("AI: %s\n", response)
}
