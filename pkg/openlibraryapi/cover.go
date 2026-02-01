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

func (c *Client) FetchCoverById(id int, name string) error {
	url := buildCoverImageUrl(id)
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

	imageFile := path.Join(name, strconv.Itoa(id)+".jpg")
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
