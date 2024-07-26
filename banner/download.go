package banner

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile downloads a file from the specified URL and saves it to the specified filepath.
func DownloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file from %v: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file from %v: %v", url, resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %v: %v", filepath, err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file %v: %v", filepath, err)
	}

	return nil
}
