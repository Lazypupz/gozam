package main

import (
	api "github.com/Lazypupz/gozam/serv"
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
	/*
		audio_file := "C:/Users/anton/Documents/GitHub/gozam/src/spectrogram/Recording.wav"
		output := wav.GetWavData(audio_file)
		wav.ReformatWav(audio_file, uint(output))
		spec := spectrogram.CreateSpectrogram()
		specAgain, err := spectrogram.SpectrogramToImg(spec, "image.png")
		peaks := spectrogram.ExtractPeaks(specAgain)
		if err != nil {
			fmt.Print("Error", err)
		}
		fpData, songID := fingerprint.GenFingerPrint(peaks)
		err = fingerprint.SaveFingerprintToDB(fpData, songID)
		if err != nil {
			fmt.Println("Error saving fingerprint to DB:", err)
		} else {
			fmt.Println("Fingerprint saved to DB successfully")
		}
	*/
	//need to fix createhash

	/*
		err := godotenv.Load()
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
		fmt.Println(text) */

	server := api.NewServer(":8080")
	server.Run()

}
