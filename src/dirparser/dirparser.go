// Takes a directory recusively and creates a JSON representation of such.
//
// {
//    "ValueMap": {
//        "DirName": "OurDir",
//        "onefile.go": "onefile.go"
//        "another.go": "another.go"
//    },
//    "Children": [{
//        "ValueMap": {
//            "DirName": "ASubDir",
//            "afile.go": "ASubDir/afile.go"
//        },
//        "Children": []
//    }, {
//        "ValueMap": {
//            "DirName": "AnotherSubDir",
//            "anotherfile.go": "AontherSubDir/anotherfile.go"
//        },
//        "Children": []
//    }]
// }
//
package dirparser

import (
	"path/filepath"
	"fileutils"
	"strings"
	"log"
	"encoding/json"	
)

// Holds a map of the file names and content
// and a slice of DirTrees for the sub directories
type DirTree struct {
	ValueMap map[string]string
	Children []DirTree
}

// Takes a dir, and gives back a DirTree with all its files, recursively
// through the sub directories.
// The output will have paths for the files. relativePath says from where they should start.
func ParseDir(relativePath string, path string) DirTree {
	node := DirTree{make(map[string]string), []DirTree{}}

	files, _ := filepath.Glob(path + "*")
	if files==nil {
		log.Print("ERROR: No files found at " + path + "*")
		return node
	}
	
	node.ValueMap["DirName"] = fileutils.Create(path).Basename()	
	node = parseAllFiles(relativePath, files, node)
	
	return node
}

// Adds all the files, and directories, onto the DirTree passed
// relativePath is used to say from where the file path should start.
func parseAllFiles(relativePath string, files []string, currentNode DirTree) DirTree {
	for _, file := range files {
		of := fileutils.Create(file)
		if of==nil {
			continue
		}
		log.Println("PARSING: " + file)
		if of.IsDir() {
			child := ParseDir(relativePath, file + "/")
			currentNode.Children = append(currentNode.Children, child)
		} else {		
			file = strings.Replace(file, relativePath, "", -1)
			currentNode.ValueMap[of.Basename()] = file//of.GetContentOfFile()
		}	
	}
	return currentNode
}

// This will freak out if you're trying to print non-UTF8 characters
func (parsed *DirTree) ToJson() string {
	j, err := json.MarshalIndent(parsed, " ", "    ")
	if err!=nil {
		log.Println("JSON ERROR: " + err.Error())
		return "JSON ERROR: " + err.Error()
	}
	return string(j)
}