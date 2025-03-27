package googleAI

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/genai"
)

type SongRecommendation struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Reason string `json:"reason"`
}

type ResponseData struct {
	RecommendedSongs []SongRecommendation `json:"recommended_songs"`
}

func Get_Song_Recommendation(apiKey string, songName, artistName string) ([]SongRecommendation, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// -------------------------------------------------- //
	prompt := fmt.Sprintf(`
	I am building a music recommendation system. Given a song name and artist, recommend 5 similar songs based on lyrical themes, mood, tempo, and genre.
	- Input: %s - %s
	- Output format: JSON with song recommendations.
	`, songName, artistName)
	// --------------------------------------------------- //

	result, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash-exp", genai.Text(prompt), nil)
	if err != nil {
		log.Fatal(err)
	}

	textResult := result.Candidates[0].Content.Parts[0].Text

	var responseData ResponseData
	err = json.Unmarshal([]byte(textResult), &responseData)
	if err != nil {
		return nil, err
	}

	return responseData.RecommendedSongs, nil
}

/*
func debugPrint[T any](r *T) {
	response, err := json.MarshalIndent(*r, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
*/
