// Some utility methods for os.File
package fileutils

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"path/filepath"
)

// Simply extends os.File
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

func (f *FileUtilType) ContentOfFile() string {
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

// Cleans all the relative paths, eg .., and then checks to see
// if we're in such a directory.
// Used for security purposes - we don't want the user going back in the directory stack
func (f *FileUtilType) IsInDirectory(dir string) bool{
	filename := f.Name()
	p := filepath.Clean(filename)
	b := strings.HasPrefix(p, dir)
	return b;
}