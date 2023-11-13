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
	token := `eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..r-F_HzvA7APrx0Qv.Jz-59M7VEy39w4iVaZB-rW8aRvbU6xUKPgIMkydzMGmNj0zTJDmV1JNkbzJkBBKIR3iAal_DqSROcCPypQ_hFRUjbtWxa-PPjTPjcEoh-28w6D-tvnBMfLbFMBNCZ_39Jh9S2kR46mb3SBBTv_9o6bFzIzZIvdO7S7ndckAfikzbJi_1rnRSV1hecwtWQJgd3U_5TL9cvWS2p85avZxVmHXE-6kakcGjXVEwqyqnoshQP3G9KWmskzc0ZqXkoQwajhmgV1qrOGsm5y7xTFRfPCV7UJcigPGAOardA3xFLySo75ZYVBrePi33Yhx6N3trKCs_GQDJ1FWfNSm1q5i6RS19DVVWhGfXJO4wSxpUY38XKhf5WdaJSz4fc3B8ma-abI7t2hlhpxFVXW5nUkHDaMafLomvlDAWCUv6kWUyx2IQKdNi0XW6emmTT8uamrhmH8bkiEkKr_s7XqE8P_byBqmLzkyy4YwHXgskSZFuH8utleGQ-H0QW_FpGcPl8DCnk2OFe4VWBY1YInvRn98AnAUeHVZX9c3UxggWy1TWvwrQ_eMA2dQlPDW0E2nb76vappYHIOxOyeiDL4N5RQLioNxE8mgM3l7cf5xT1VJuqQt1KV8QAm4j_ROHZxoZL2gv-X5TldF5gUofnumKsaGh5DPpo2FzxcC_rp_M6llhWbADf2KFuIKB9qzM-pQetDf_LryXL1mTjKpnxkA65NBYWHOKf8lP5xdhDDfRSW3SA3UQrOmEux7_6jWr-HU9yKdyzsFJP6VrGZTGHcEnXCMlfh7W5v9XpAjOzmIRTzwOudnWqv6Q5H1fJVPEsgjAHXlGgOTQbjn17EBHeNKbQGoJhD6xmF7gBhVotAxVtF2LFvNwl_qeG2dNTbYTxbt6ImrI5Ap3dKfnjJEGcTcPCVyg104lOGge63iYmYdvU3iW0CBd…`
	cfValue := "GjeM0ksLecmSZIW"

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
		chatgpt.WithDebug(true),
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
