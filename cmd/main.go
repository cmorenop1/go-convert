package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cmorenop1/go-converter/internals/components"
)

type Stepper struct {
	CurrentStep int
}

const (
	inputStep = iota
	outputStep
	executeStep
	quit
)

func main() {
	components.Welcome()
	stepper := NewStepper()
	reader := bufio.NewReader(os.Stdin)

	var inputFile string
	var outputFile string

	for {
		switch stepper.CurrentStep {
		case inputStep:
			fmt.Println("\n[Input (.mov) filename]")
			inputSting, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				continue
			}
			inputFile = strings.TrimSpace(inputSting)
			stepper.Next()
		case outputStep:
			fmt.Println("\n[Output (.mp4) filename]")
			outputString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				continue
			}
			outputFile = strings.TrimSpace(outputString)
			stepper.Next()
		case executeStep:
			fmt.Println("\n 🔨 Converting...")
			err := convertMovToMp4(inputFile, outputFile)
			if err != nil {
				fmt.Printf("Conversion error: %v\n", err)
			} else {
				fmt.Println("✅ Conversion successful!")
			}
			stepper.Next()
		case quit:
			fmt.Println("\n ✨ Done !!")
			os.Exit(0)
		}
	}
}

func NewStepper() Stepper {
	return Stepper{
		CurrentStep: 0,
	}
}

func (s *Stepper) Next() {
	s.CurrentStep++
}

func convertMovToMp4(inputFile, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-c:v", "copy", "-c:a", "aac", "-strict", "experimental", outputFile)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
