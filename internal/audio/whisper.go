package audio

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func CreateTranscription(filepath string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	payload := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filepath,
	}

	res, err := client.CreateTranscription(ctx, payload)

	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		panic("Transcription error with OpenAI API")
	}

	return res.Text
}
