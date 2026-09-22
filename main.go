package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	chatgpt "github.com/chatgp/chatgpt-go"
)

func main() {
	// New ChatGPT client
	token := os.Getenv("CHATGPT_SESSION_TOKEN")
	cfValue := os.Getenv("CF_CLEARANCE")

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: token,
		},
		{
			Name:  "cf_clearance",
			Value: cfValue,
		},
	}

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(false),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)

	// Get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an English word or phrase: ")
	englishText, _ := reader.ReadString('\n')
	englishText = strings.TrimSpace(englishText)

	// Send the English text to ChatGPT for translation
	stream, err := cli.GetChatStream(englishText)
	if err != nil {
		log.Fatalf("get chat stream failed: %v\n", err)
	}

	var translatedText string
	for text := range stream.Stream {
		log.Printf("stream text: %s\n", text.Content)
		translatedText = text.Content
	}

	if stream.Err != nil {
		log.Fatalf("stream closed with error: %v\n", stream.Err)
	}

	log.Printf("English: %s\n", englishText)
	log.Printf("German: %s\n", translatedText)
}
