package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	// Open the current directory
	dirHandle, err := os.Open(dir)
	if err != nil {
		fmt.Printf("Error opening directory: %v\n", err)
		return
	}
	defer dirHandle.Close()

	// Get all files in the directory
	fileInfos, err := dirHandle.Readdir(-1)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Loop through each file
	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() {
			fileName := fileInfo.Name()
			if strings.HasSuffix(fileName, ".mov") {
				// Convert .mov file to .mp4
				err := convertMovToMp4(fileName)
				if err != nil {
					fmt.Printf("Error converting %s: %v\n", fileName, err)
					continue
				}

				// Remove the original .mov file
				err = os.Remove(fileName)
				if err != nil {
					fmt.Printf("Error removing %s: %v\n", fileName, err)
				} else {
					fmt.Printf("Converted and removed: %s\n", fileName)
				}
			}
		}
	}
}

func convertMovToMp4(filename string) error {
	// Convert MOV to MP4
	cmd := exec.Command("ffmpeg", "-i", filename, "-c:v", "libx264", "-b:v", "2M", "-c:a", "aac", "-b:a", "128K", "-movflags", "+faststart", strings.TrimSuffix(filename, ".mov")+".mp4")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
