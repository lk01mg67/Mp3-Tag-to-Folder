package organizer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/dhowden/tag"
)

var (
	replacer = strings.NewReplacer(
		":", "：",
		"/", "／",
		"?", "？",
		"*", "＊",
		">", "＞",
		"<", "＜",
		"\"", "＂",
		"|", "｜",
		"...", "…",
		"..", "…",
	)
)

func DetermineDestPath(oldPathFile string, outputDir string) (string, error) {
	f, err := os.Open(oldPathFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	m, err := tag.ReadFrom(io.ReadSeeker(f))
	if err != nil {
		return "", err // Return error if tag reading fails
	}

	// Default values if tags are missing
	artist := "Unknown Artist"
	album := "Unknown Album"
	// title := "Unknown Title"
	disc := 1

	if m != nil {
		if strings.TrimSpace(m.AlbumArtist()) != "" {
			artist = m.AlbumArtist()
		} else if strings.TrimSpace(m.Artist()) != "" {
			artist = m.Artist()
		}

		if strings.TrimSpace(m.Album()) != "" {
			album = m.Album()
		}

		// if strings.TrimSpace(m.Title()) != "" {
		// 	title = m.Title()
		// }

		d, _ := m.Disc()
		if d > 0 {
			disc = d
		}
	}

	_, file := filepath.Split(oldPathFile)

	cleanArtist := replacer.Replace(strings.TrimSpace(artist))
	cleanAlbum := replacer.Replace(strings.TrimSpace(album))

	newPath := filepath.Join(outputDir, cleanArtist, cleanAlbum, fmt.Sprintf("%d", disc))

	// 0755 for directories (rwxr-xr-x)
	err = os.MkdirAll(newPath, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	return filepath.Join(newPath, file), nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CopyFile(src, dst string) (int64, error) {
	if fileExists(dst) {
		return 0, nil
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
