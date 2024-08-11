package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func createHLS(inputFile string, outputDir string, segmentDuration int) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	if segmentDuration == 0 {
		segmentDuration = 10
	}

	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", inputFile,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", strconv.Itoa(segmentDuration),
		"-hls_list_size", "0",
		"-f", "hls",
		fmt.Sprintf("%s/playlist.m3u8", outputDir),
	)

	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create HLS: %v\nOutput: %s", err, string(output))
	}

	return nil
}

func main() {
	inputFile := "./video/video.mp4"
	outputDir := "./output"
	segmentDuration := 10

	if err := createHLS(inputFile, outputDir, segmentDuration); err != nil {
		log.Fatalf("Error creating HLS: %v", err)
	}

	log.Println("HLS created successfully")
}
