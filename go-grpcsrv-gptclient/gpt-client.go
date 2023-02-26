package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GptCompletionClient(prompt []string) string {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	//var stop = []string{"人間:", "AI:"}
	//var stop []string = []string{"\n"}
	resp, err := client.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:           prompt,
		MaxTokens:        gpt3.IntPtr(206),
		Temperature:      gpt3.Float32Ptr(1),
		TopP:             gpt3.Float32Ptr(0.5),
		PresencePenalty:  0.6,
		FrequencyPenalty: 0.25,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.Choices[0].Text
}

func GptEditClient(input string, instruction string) string {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp, err := client.Edits(ctx, gpt3.EditsRequest{
		Model:       "text-davinci-edit-001",
		Input:       input,
		Instruction: instruction,
		Temperature: gpt3.Float32Ptr(0.7),
		TopP:        gpt3.Float32Ptr(0.3),
		N:           nil,
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("ctx:", ctx)
	return resp.Choices[0].Text
}
