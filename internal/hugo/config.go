package hugo

import "path"

type config struct {
	hugoPath   string
	dataDir    string
	contentDir string
}

type Hugo struct {
	Config config
}

func NewHugo(fullPath string) Hugo {
	return Hugo{
		Config: config{
			hugoPath:   fullPath,
			dataDir:    path.Join(fullPath, "data"),
			contentDir: path.Join(fullPath, "content"),
		},
	}
}
