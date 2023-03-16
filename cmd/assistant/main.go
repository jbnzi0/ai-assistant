package main

import (
	"fmt"
	"log"

	"github.com/jbnzi0/virtual-assistant/internal/audio"
	"github.com/jbnzi0/virtual-assistant/internal/chatgpt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	transcription := audio.CreateTranscription("../../assets/temp.m4a")
	answer := chatgpt.Chat(transcription)

	transcriptionId := audio.ConvertTextToSpeech(answer)
	fmt.Println(transcription, answer, transcriptionId)
}
