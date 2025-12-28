package main

import (
	"context"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/prompts"
)

type CodeReviewer struct {
	llm      llms.Model
	template *prompts.PromptTemplate
}

func NewCodeReviewer() (*CodeReviewer, error) {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-2.5-flash"))
	if err != nil {
		return nil, err
	}

	template := prompts.NewPromptTemplate(`
		You are an expert Go code reviewer. Analyze this Go code for:

		1. **Bugs and Logic Issues**: Potential runtime errors, nil pointer dereferences, race conditions
		2. **Performance**: Inefficient algorithms, unnecessary allocations, string concatenation issues
		3. **Style**: Go idioms, naming conventions, error handling patterns
		4. **Security**: Input validation, sensitive data handling

		Code to review:
		'''go
		{{.code}}
		'''

		File: {{.filename}}

		Provide specific, actionable feedback. For each issue:
		- Explain WHY it's a problem
		- Show HOW to fix it with code examples
		- Rate severity: Critical, Warning, Suggestion

		Focus on the most important issues first.
		`, []string{"code", "filename"})

	return &CodeReviewer{
		llm:      llm,
		template: &template,
	}, nil
}

func (cr *CodeReviewer) ReviewFile(ctx context.Context, filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, filename, content, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parsing Go file: %w", err)
	}

	prompt, err := cr.template.Format(map[string]any{
		"code":     string(content),
		"filename": filename,
	})
	if err != nil {
		return fmt.Errorf("formatting prompt: %w", err)
	}

	response, err := cr.llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	})
	if err != nil {
		return fmt.Errorf("generating review: %w", err)
	}

	fmt.Printf("\n=== Review for %s ===\n", filename)
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println(response.Choices[0].Content)
	fmt.Println(strings.Repeat("=", 80))

	return nil
}

func main() {
	ctx := context.Background()
	file := flag.String("file", "", "Go file to review")
	flag.Parse()

	reviewer, err := NewCodeReviewer()
	if err != nil {
		log.Fatal(err)
	}

	if err := reviewer.ReviewFile(ctx, *file); err != nil {
		log.Fatal(err)
	}
}
