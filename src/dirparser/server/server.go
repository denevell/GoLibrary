// Serves the DirTree on a server, which supports JSONP, 
// and getting the files specified in the DirTree
package server

import (
	"fileutils"
	"jsonutils"
	"fmt"
	"log"
	"net/http"
	"os"
	"dirparser"
)

type ContentFunction func() string

// The DirectoryTree to server, the port to serve it at, and the 
// RelativeFilePath from which to serve files
type ServerInfo struct {
	DirectoryTree     dirparser.DirTree
	Port              string
	RelativeFilePath  string
}

// Implementation of http.Handler, to pass to http.ListenAndServe, 
// which is used in ServerInfo.StartServer. 
// /               - Pass the JSON of the parsed directory
// /quit           - kill the server
// /file?p=file    - read the file after ? with the ServerInfo.RelativeFilePath directory 
// You can add q=func to / and /file to return the values as jsonp
func (serverInfo *ServerInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		log.Println("REQUESTED on dirserver: " + r.URL.RequestURI())
		jsonutils.SetJsonHeader(&w)

		switch val := r.URL.Path; {
		case val == "/":
			if jsonp := r.URL.Query().Get("q"); len(jsonp)>0 {
				fmt.Fprintf(w, "%s(%s)", jsonp, serverInfo.DirectoryTree.ToJson())
			} else {
				fmt.Fprintf(w, "%s", serverInfo.DirectoryTree.ToJson())
			}
			
		case val == "/quit":
			os.Exit(0)		

		case val == "/file":
			path := r.URL.Query().Get("p")
			if of := fileutils.Create(serverInfo.RelativeFilePath + path); of != nil && of.IsInDirectory(serverInfo.RelativeFilePath) {
				if jsonp := r.URL.Query().Get("q"); len(jsonp)>0 {
					o := of.ContentOfFile()
					fmt.Fprintf(w, "%s(%s)", jsonp, jsonutils.ConvertMultilineToJSONString(o))
				} else {
					fmt.Fprintf(w, of.ContentOfFile())
				}
			}
		}
}

// See ServerInfo.ServeHTTP
func (serverInfo *ServerInfo) StartServer() {
	log.Println("Starting dirparsed server...")
	if err := http.ListenAndServe(":"+serverInfo.Port, serverInfo); err != nil {
		log.Println(err)
	}
}
