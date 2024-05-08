package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	openaiApiKey := getEnvVar("OPENAI_API_KEY")
	openaiBaseUrl := getEnvVar("OPENAI_BASE_URL")

	llm, err := openai.New(openai.WithToken(openaiApiKey), openai.WithBaseURL(openaiBaseUrl))
	if err != nil {
		log.Fatal(err)
	}

	prompt := "你是谁"
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}

// getEnvVar 获取环境变量的值
func getEnvVar(key string) string {
	value := ""
	if _, ok := os.LookupEnv(key); ok {
		value = os.Getenv(key)
	}
	return value
}
