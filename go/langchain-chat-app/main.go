package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-2.5-flash"))
	if err != nil {
		log.Fatal(err)
	}

	chatMemory := memory.NewConversationBuffer()
	conversationChain := chains.NewConversation(llm, chatMemory)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Chat Application Started (type 'quit' to exit)")
	fmt.Println("---------------------------------------")

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		result, err := chains.Run(ctx, conversationChain, input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("AI: %s\n\n", result)
	}
}
