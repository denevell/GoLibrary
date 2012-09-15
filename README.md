GoLibrary
=========

Some Golang libraries.

Here's the godoc output. Run go doc <packagename> for the latest

	PACKAGE
	
	package dirparser
	    import "dirparser"
	
	    Takes a directory recusively and creates a JSON representation of such.
	
	    {
	
		"ValueMap": {
		    "DirName": "OurDir",
		    "onefile.go": "file content"
		    "another.go": "file content"
		},
		"Children": [{
		    "ValueMap": {
		        "DirName": "ASubDir",
		        "afile.go": "file content"
		    },
		    "Children": []
		}, {
		    "ValueMap": {
		        "DirName": "AnotherSubDir",
		        "anotherfile.go": "file content"
		    },
		    "Children": []
		}]
	
	    }
	
	    This will obviously freak out if you parse a directory with too many
	    files
	
	TYPES
	
	type DirTree struct {
	    ValueMap map[string]string
	    Children []DirTree
	}
	    Holds a map of the file names and content and a slice of DirTrees for
	    the sub directories
	
	func ParseDir(path string) DirTree
	    Takes a dir, and gives back a DirTree with all its files, recursively
	    through the sub directories
	
	func (parsed *DirTree) ToJson() string
	    This will freak out if you're trying to print non-UTF8 characters I.e.
	    you're trying to print out a binary in JSON
	
	
	PACKAGE
	
	package server
	    import "server"
	
	
	TYPES
	
	type ContentFunction func() string
	
	type ServerInfo struct {
	    IndexFileLocation string
	    ContentFunction   ContentFunction
	    Port              string
	}
	
	func (serverInfo *ServerInfo) StartServer()
	
	
	PACKAGE
	
	package fileutils
	    import "fileutils"
	
	
	TYPES
	
	type FileUtilType struct {
	    os.File
	}
	
	func Create(file string) *FileUtilType
	    Create a new FileUtilType, a os.File. With FileUtilType, we've added
	    various utility methods in this package
	
	func (f *FileUtilType) Basename() string
	
	func (f *FileUtilType) GetContentOfFile() string
	
	func (f *FileUtilType) IsDir() bool
	
	
	PACKAGE
	
	package flagutils
	    import "flagutils"
	
	    Utility methods for the flag package
	
	FUNCTIONS
	
	func GetStringFlagOrExit(name string, desc string) string
	    Like a flag.String but if the flag is "", we exit the program altogether
