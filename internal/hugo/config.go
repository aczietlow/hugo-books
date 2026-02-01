package hugo

import (
	"github.com/aczietlow/hugo-books/internal/config"
)

type configuration struct {
	hugoPath   string
	dataDir    string
	contentDir string
	imageDir   string
}

type Hugo struct {
	Config configuration
}

func NewHugo(conf *config.Config) Hugo {
	return Hugo{
		Config: configuration{
			hugoPath:   conf.Hugo.BasePath,
			dataDir:    conf.Hugo.DataDir,
			contentDir: conf.Hugo.ContentDir,
			imageDir:   conf.Hugo.ImageDir,
		},
	}
}
