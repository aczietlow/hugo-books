package config

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Hugo        hugoConfig        `yaml:"hugo"`
	OpenLibrary openLibraryConfig `yaml:"openLibrary"`
}

type hugoConfig struct {
	BasePath   string `yaml:"basePath"`
	DataDir    string `yaml:"dataDir"`
	ContentDir string `yaml:"contentDir"`
	ImageDir   string `yaml:"imageDir"`
}

type openLibraryConfig struct {
	HTTPTimeout int    `yaml:"httpTimeout"`
	CacheTTL    int    `yaml:"cacheTTL"`
	UserAgent   string `yaml:"userAgent"`
	BaseUrl     string `yaml:"baseUrl"`
}

func LoadConfig(filepath string) (*Config, error) {
	return loadConfigFromFile(os.DirFS("."), filepath)
}

func loadConfigFromFile(fsys fs.FS, name string) (*Config, error) {
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s", string(data))
	c := Config{}
	if err = yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func LoadConfigFromBytes(b []byte) (*Config, error) {
	var conf Config
	if err := yaml.Unmarshal(b, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
