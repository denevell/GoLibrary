package fileutils

import (
	"os"
	"log"
	"io/ioutil"
	"path/filepath"
)

type FileUtilType struct {
	os.File
}

// Create a new FileUtilType, a os.File.
// With FileUtilType, we've added various utility methods in this package 
func Create(file string) *FileUtilType {
	fi, err := os.Open(file)
	if err!=nil {
		log.Println("ERROR CREATING FILE:" + err.Error())
		return nil
	}
	f := &FileUtilType{*fi}
	return f
}

func (f *FileUtilType) IsDir() bool {
	fileInfo, err := f.Stat()
	if err!=nil {
		log.Println("ERROR DETECTING DIR:" + err.Error())
		return false
	}
	isDir := fileInfo.IsDir()
	return isDir
}

func (f *FileUtilType) GetContentOfFile() string {
	ff, err := ioutil.ReadAll(f)
	if err!=nil {
		log.Println("ERROR READING FILE:" + err.Error())
		return ""
	}	
	fff := string(ff)
	return fff
}

func (f *FileUtilType) Basename() string {
	filestat, _ := f.Stat()
	filename := filestat.Name()
	basename := filepath.Base(filename)
	return basename
}