package pages

import (
	"fmt"
	"path/filepath"
	"strings"
)

func (srv *PagesSrv) ListPages() ([]string, error) {
	var list []string

	pagesDir := filepath.Join(srv.workdir, "pages")
	filenames, err := srv.repos.Fs.ListFiles(pagesDir)
	if err != nil {
		return list, fmt.Errorf("Error: failed to list files")
	}

	// trim workdir path
	for _, filename := range filenames {
		rel, err := filepath.Rel(pagesDir, filename)
		if err != nil {
			return list, err
		}
		path := strings.TrimSuffix(rel, ".json")
		list = append(list, path)
	}

	return list, err
}
