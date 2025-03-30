package wav

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Record(outputFile string) {

	//if windows
	cmd := exec.Command(
		"ffmpeg",
		"-f",
		"dshow",
		"-i",
		"audio=Microphone Array",
		"-t",
		"5",
		outputFile,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Recording saved to:", outputFile)
}

func GetWavData(inputFile string) int {
	cmd := exec.Command(
		"ffprobe",
		"-v",
		"error",
		"-select_streams",
		"a:0",
		"-show_entries",
		"stream=channels",
		"-of",
		"csv=p=0",
		inputFile,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error (getwavdata)", err)
		return 0
	}
	// trims all white space
	stringedOutput := strings.TrimSpace(string(output))

	// coverts to int
	cNum, cErr := strconv.Atoi(stringedOutput)
	if cErr != nil {
		fmt.Println("Error (getwavdataint)", cErr)
		return 0
	}

	return cNum
	// --todo doesnt output anything, doesnt cause errors also
	// nvm it works i just forgot to print the output
}
