package gowc

import (
	"os"
	"strings"
	"unsafe"
)

func NewGoWC(path string) GoWC {
	return &goWC{
		path: path,
	}
}

type GoWC interface {
	Counte() (int, error)
}

type goWC struct {
	path string
}

func (wc *goWC) Counte() (total int, err error) {
	if wc == nil {
		err = ErrInvalidProgram
		return
	}

	if IsDir(wc.path) {
		var files []string
		files, err = DirWalk(wc.path)
		if err != nil {
			return
		}
		total = 0
		for _, f := range files {
			row, err := Counte(f)
			if err != nil {
				continue
			}
			total += row
		}
		return
	}

	total, err = Counte(wc.path)
	return
}

func Counte(file string) (n int, err error) {
	var f *os.File
	if f, err = os.Open(file); err != nil {
		return
	}

	var info os.FileInfo
	if info, err = f.Stat(); err != nil {
		return
	}

	buf := make([]byte, info.Size())
	if _, err = f.Read(buf); err != nil {
		return
	}

	firstRow := 1
	otherRow := strings.Count(*(*string)(unsafe.Pointer(&buf)), "\n")
	n = firstRow + otherRow
	return
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func IsDir(path string) bool {
	var err error
	var fileInfo os.FileInfo
	if fileInfo, err = os.Stat(path); err != nil {
		return false
	}
	return fileInfo.IsDir()
}
