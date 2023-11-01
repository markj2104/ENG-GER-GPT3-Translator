package main

import (
	"fmt"
	"os"
)

func main() {
	// Retrieve the API key from an environment variable
	apiKey := os.Getenv("GPT3_API_KEY")

	if apiKey == "" {
		fmt.Println("API key not set. Please set the GPT3_API_KEY environment variable.")
		return
	}

	client := gpt3.NewClient(apiKey)

	// English text to be translated
	englishText := "Hello, how are you?"

	// Create a prompt to translate the English text to German
	prompt := fmt.Sprintf("Translate the following English text to German: \"%s\"", englishText)

	// Request translation from GPT-3
	response, err := client.Completions.Create(
		gpt3.CompletionRequest{
			Prompt:      prompt,
			Temperature: 0.7,
			MaxTokens:   30, // You can adjust the number of tokens as needed
		},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Extract the translated text from GPT-3's response
	translatedText := response.Choices[0].Text

	// Print the translation
	fmt.Printf("English: %s\n", englishText)
	fmt.Printf("German: %s\n", translatedText)
}
