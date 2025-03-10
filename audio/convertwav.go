package shazam

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// THIS CODE WAS DIRECTLY COPIED FROM SEEK_TUNE AS I HAD DIDNT WANT TO WASTE TIME ON THIS SIMPLE CODE

func ConvertToWave(inputFilePath string, channels int) (wavFilePath string, err error) {
	_, err = os.Stat(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if channels < 1 || channels > 1 {
		channels = 1
	}

	fileExt := filepath.Ext(inputFilePath)
	outputFile := strings.TrimSuffix(inputFilePath, fileExt) + ".wav"

	// Output file may already exists. If it does FFmpeg will fail as
	// it cannot edit existing files in-place. Use a temporary file.
	tmpFile := filepath.Join(filepath.Dir(outputFile), "tmp_"+filepath.Base(outputFile))
	defer os.Remove(tmpFile)

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", inputFilePath,
		"-c", "pcm_s16le",
		"-ar", "44100",
		"-ac", fmt.Sprint(channels),
		tmpFile,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to convert to WAV: %v, output %v", err, string(output))
	}

	// Rename the temporary file to the output file
	err = os.Rename(tmpFile, outputFile)
	if err != nil {
		return "", fmt.Errorf("failed to rename temporary file to output file: %v", err)
	}

	return outputFile, nil
}

//if alr wav

func ReformatWav(inputFilePath string, channels int) (reformatedFilePath string, err error) {
	if channels < 1 || channels > 1 {
		channels = 1
	}

	fileExt := filepath.Ext(inputFilePath)
	outputFile := strings.TrimSuffix(inputFilePath, fileExt)

	cmd := exec.Command(

		"ffmpeg",
		"-y",
		"-i", inputFilePath,
		"-c", "pcm_s16le",
		"-ar", "44100",
		"-ac", fmt.Sprint(channels),
		outputFile,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("fialed to convertto wav: %v, output %v", err, string(output))
	}
	return outputFile, nil
}
