package openlibraryapi

import (
	"fmt"
	"strings"
)

type author struct {
	Name           string   `json:"name"`
	PersonalName   string   `json:"personal_name"`
	FullerName     string   `json:"fuller_name"`
	Key            string   `json:"key"`
	EntityType     string   `json:"entity_type"`
	BirthDate      string   `json:"birth_date"`
	AlternateNames []string `json:"alternate_names"`
	Title          string   `json:"title"`
	Photos         []int    `json:"photos"`
}

func (a *author) getAuthorId() (string, error) {
	if strings.Contains(a.Key, "/authors/") {
		return strings.Trim(a.Key, "/authors/"), nil
	}
	return "", fmt.Errorf("%s not a valid key", a.Key)
}
