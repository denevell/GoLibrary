// Takes a directory recusively and creates a JSON representation of such.
//
// {
//    "ValueMap": {
//        "DirName": "OurDir",
//        "onefile.go": "file content"
//        "another.go": "file content"
//    },
//    "Children": [{
//        "ValueMap": {
//            "DirName": "ASubDir",
//            "afile.go": "file content"
//        },
//        "Children": []
//    }, {
//        "ValueMap": {
//            "DirName": "AnotherSubDir",
//            "anotherfile.go": "file content"
//        },
//        "Children": []
//    }]
// }
//
// This will obviously freak out if you parse a directory with too many files
package dirparser

import (
	"path/filepath"
	"fileutils"
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
// through the sub directories 
func ParseDir(path string) DirTree {
	node := DirTree{make(map[string]string), []DirTree{}}

	files, _ := filepath.Glob(path + "*")
	if files==nil {
		log.Print("ERROR: No files found at " + path + "*")
		return node
	}
	
	node.ValueMap["DirName"] = fileutils.Create(path).Basename()	
	node = parseAllFiles(files, node)
	
	return node
}

// Adds all the files, and directories, onto the DirTree passed
func parseAllFiles(files []string, currentNode DirTree) DirTree {
	for _, file := range files {
		of := fileutils.Create(file)
		if of==nil {
			continue
		}
		log.Println("PARSING: " + file)
		if of.IsDir() {
			child := ParseDir(file + "/")
			currentNode.Children = append(currentNode.Children, child)
		} else {		
			currentNode.ValueMap[of.Basename()] = of.GetContentOfFile()
		}	
	}
	return currentNode
}

// This will freak out if you're trying to print non-UTF8 characters
// I.e. you're trying to print out a binary in JSON
func (parsed *DirTree) ToJson() string {
	j, err := json.Marshal(parsed)
	if err!=nil {
		log.Println("JSON ERROR: " + err.Error())
		return "JSON ERROR: " + err.Error()
	}
	return string(j)
}