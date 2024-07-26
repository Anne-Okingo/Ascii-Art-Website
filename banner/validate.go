package banner

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Expected SHA-256 hashes for banner files
const (
	standardHash   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowHash     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyHash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

// Map of banner files to their expected hashes
var BannerHashes = map[string]string{
	"bannerfiles/standard.txt":   standardHash,
	"bannerfiles/shadow.txt":     shadowHash,
	"bannerfiles/thinkertoy.txt": thinkertoyHash,
}

// Map of banner files to their download URLs
var BannerURLs = map[string]string{
	"bannerfiles/standard.txt":   "https://learn.zone01kisumu.ke/git/root/public/raw/commit/07e39a68787f57edd0ea8b16caa9ee8f7360869b/subjects/ascii-art/standard.txt",
	"bannerfiles/shadow.txt":     "https://learn.zone01kisumu.ke/git/root/public/raw/commit/07e39a68787f57edd0ea8b16caa9ee8f7360869b/subjects/ascii-art/shadow.txt",
	"bannerfiles/thinkertoy.txt": "https://learn.zone01kisumu.ke/git/root/public/raw/commit/07e39a68787f57edd0ea8b16caa9ee8f7360869b/subjects/ascii-art/thinkertoy.txt",
}

// ValidateBanner checks if a specific banner file is present and corrects it if necessary.
func ValidateBanner(fileName string) {
	expectedHash, hashOk := BannerHashes[fileName]
	if !hashOk {
		fmt.Printf("Error: Unknown banner file %v\n", fileName)
		os.Exit(1)
	}

	// Check if the banner file exists and is valid
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: %v is missing.\n", fileName)
			fmt.Println("File downloaded succesfully")

		} else {
			fmt.Printf("Error: Could not retrieve file information for %v: %v\n", fileName, err)
		}
		DownloadBannerFile(fileName)
		return
	}

	if fileInfo.Size() == 0 {
		fmt.Printf("Error: %v is empty.\n", fileName)
		fmt.Println("File re-downloaded succesfully")
		DownloadBannerFile(fileName)
		return
	}

	// Check file hash
	fileHash, err := computeFileHash(fileName)
	if err != nil {
		fmt.Printf("Error: Could not compute hash for %v: %v\n", fileName, err)
		DownloadBannerFile(fileName)
		return
	}
	if fileHash != expectedHash {
		fmt.Printf("Error: %v  does not match expected file (Content Modified).\n", fileName)
		fmt.Println("File re-downloaded succesfully")
		DownloadBannerFile(fileName)
		return
	}
}

// computeFileHash computes the SHA-256 hash of the given file.
func computeFileHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// downloadBannerFile attempts to download a new banner file.
func DownloadBannerFile(bannerName string) {
	// fmt.Println("Attempting to download a new banner file...")
	downloadURL, ok := BannerURLs[bannerName]
	if !ok {
		fmt.Printf("Error: No download URL for banner file %v\n", bannerName)
		os.Exit(1)
	}
	err := DownloadFile(downloadURL, bannerName)
	if err != nil {
		fmt.Printf("Failed to download the banner file: %v\n", err)
		os.Exit(1)
	}
}
