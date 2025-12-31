package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"mp3/internal/organizer"
)

func main() {
	var inputDir string
	var outputDir string

	flag.StringVar(&inputDir, "input", ".", "Source directory to scan for MP3 files")
	flag.StringVar(&outputDir, "output", "output", "Destination directory to sort files into")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Printf("Scanning: %s\nTarget: %s\n", inputDir, outputDir)

	err := filepath.Walk(inputDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return nil // Continue walking
		}

		if info.IsDir() {
			return nil
		}

		// Simple extension check (case-insensitive)
		if strings.ToLower(filepath.Ext(path)) != ".mp3" {
			return nil
		}

		newPathFile, err := organizer.DetermineDestPath(path, outputDir)
		if err != nil {
			log.Printf("Skipping %s: %v\n", path, err)
			return nil
		}

		nBytes, err := organizer.CopyFile(path, newPathFile)
		if err != nil {
			log.Printf("Failed to copy %s: %v\n", path, err)
		} else {
			if nBytes > 0 {
				fmt.Printf("Copied: %s -> %s\n", path, newPathFile)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v\n", err)
	}

	fmt.Println("Done.")
}
