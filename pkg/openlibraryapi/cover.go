package openlibraryapi

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func (c *Client) FetchCoverById(id string, name string) error {
	sid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	url := buildCoverImageUrl(sid)
	imageFile := path.Join(name, id+".jpg")

	// Do nothing if cover image already exists.
	_, err = os.Stat(imageFile)
	if os.IsExist(err) {
		return nil
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
