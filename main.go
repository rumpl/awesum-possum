package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Check if API key is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable not set")
		fmt.Println("Please set your OpenAI API key: export OPENAI_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	// Get the question from command line arguments
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: awesum-possum <question>")
		os.Exit(1)
	}

	// Combine all arguments into a single question
	question := strings.Join(args, " ")

	// Create OpenAI client
	client := openai.NewClient(apiKey)

	// Create request
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
		Stream: true,
	}

	// Create streaming completion
	ctx := context.Background()
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("Error creating completion stream: %v\n", err)
		os.Exit(1)
	}
	defer stream.Close()

	// Process the streaming response
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			os.Exit(1)
		}

		content := response.Choices[0].Delta.Content
		fmt.Print(content)
	}
	fmt.Println() // Add a newline at the end
}