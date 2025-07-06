package main

import (
	"fmt"
	"log"
	"os"

	googleAI "github.com/Lazypupz/gozam/gemini-go"
	api "github.com/Lazypupz/gozam/serv/apis"
	"github.com/Lazypupz/gozam/src/fingerprint"
	"github.com/Lazypupz/gozam/src/spectrogram"
	"github.com/Lazypupz/gozam/src/wav"
	"github.com/joho/godotenv"
)

func main() {
	/*spec := spectrogram.CreateSpec()
	outputPath := "spectrogram.png"
	peaks, err := spectrogram.SpectrogramToImg(spec, outputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("spectrogram image saved to", outputPath)

	for i, val := range peaks {
		for j := range val {
			fmt.Println(peaks[i][j])
		}
	}
	*/

	audio_file := "C:/Users/anton/Documents/GitHub/gozam/src/spectrogram/Recording.wav"
	output := wav.GetWavData(audio_file)
	wav.ReformatWav(audio_file, uint(output))
	spec := spectrogram.CreateSpectrogram()
	convertedSpec, err := spectrogram.SpectrogramToImg(spec, "image.png")
	peaks := spectrogram.ExtractPeaks(convertedSpec)
	if err != nil {
		fmt.Print("Error", err)
	}
	fpData, songID := fingerprint.GenFingerPrint(peaks)

	server := api.NewServer(":8080")
	go server.Run()

	go func() {
		err = fingerprint.SaveFingerprintToDB(fpData, songID)
		if err != nil {
			fmt.Println("Error saving fingerprint to DB:", err)
		} else {
			fmt.Println("Fingerprint saved to DB successfully")
		}
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading env file")
		}

		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is not set in env")
		}
		text, songErr := googleAI.Get_Song_Recommendation(apiKey, "Dont Stop me now", "Queen")
		if songErr != nil {
			log.Fatal("Didnt work lil bro:", songErr)
		}
		fmt.Println(text)
	}()
	select {} // keeps the go routine running
	// i hate go routines //  and db
	//need to fix createhash

}
