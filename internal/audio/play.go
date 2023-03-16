package audio

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SpeechRequest struct {
	Voice   string   `json:"voice"`
	Content []string `json:"content"`
}

type SpeechResponse struct {
	Status          string `json:"status"`
	TranscriptionId string `json:"transcriptionId"`
	Error           string `json:"error"`
}

func ConvertTextToSpeech(text string) string {
	endpoint := "https://play.ht/api/v1/convert"
	payload := SpeechRequest{
		Voice:   "en-US-MichelleNeural",
		Content: []string{text},
	}

	data, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", os.Getenv("PLAY_API_KEY"))
	req.Header.Add("X-User-ID", os.Getenv("PLAY_USER_ID"))
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var result SpeechResponse
	json.Unmarshal(body, &result)

	return result.TranscriptionId
}

func FetchAudioFile(transcriptionId string) {

}
