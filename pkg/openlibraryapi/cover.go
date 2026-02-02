package openlibraryapi

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func (c *Client) FetchCoverById(id string, name string) error {
	// skip if coverId is empty.
	if id == "" {
		return nil
	}
	sid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	url := buildCoverImageUrl(sid)
	imageFile := path.Join(name, id+".jpg")

	// Do nothing if cover image already exists.
	_, err = os.Stat(imageFile)
	if err == nil {
		return nil
	} else if !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received a %d reponse from the api\n GET %s\n", resp.StatusCode, url)
	}

	contentType := resp.Header.Get("content-type")
	if !strings.HasPrefix(contentType, "image") {
		return fmt.Errorf("Response was not image type")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(imageFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Takes the coverimageId and returns the url for the full sized cover image
func buildCoverImageUrl(coverId int) string {
	return fmt.Sprintf("https://covers.openlibrary.org/b/id/%s.jpg", strconv.Itoa(coverId))
}
