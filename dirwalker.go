package gowc

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func DirWalk(dir string) (files []string, err error) {
	var fileInfos []os.FileInfo
	fileInfos, err = ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			files = append(files, filepath.Join(dir, fileInfo.Name()))
			continue
		}

		var _files []string
		_files, err = DirWalk(filepath.Join(dir, fileInfo.Name()))
		if err != nil {
			return
		}

		files = append(files, _files...)
	}
	return
}
