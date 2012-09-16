// Generic server that serves an index page
// To be removed/refactored
package server

import (
	"fileutils"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ContentFunction func() string

type ServerInfo struct {
	IndexFileLocation string
	Port              string
}

func (serverInfo *ServerInfo) StartServer() {
	//    var c chan string = make(chan string) //playing around with channels for no reason
	//    go checkForSomething(c)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("REQUESTED: " + r.URL.RequestURI())

		switch val := r.URL.Path; {
		case val == "/index":
			f, _ := os.Open(serverInfo.IndexFileLocation)
			of := fileutils.FileUtilType{*f}
			content := of.ContentOfFile()
			fmt.Fprintf(w, "%s", string(content))
		case val == "/quit":
			os.Exit(0)
			//			case val=="/s":
			//				if file := r.URL.Query()["thing"]; len(file)>0 {
			//					c <- r.URL.Query()["thing"][0]
			//					fmt.Fprintf(w, "hiii")
			//				}			
		}
	})
	log.Println("Starting server...")
	if err := http.ListenAndServe(":"+serverInfo.Port, nil); err != nil {
		log.Println(err)
	}
}

func checkForSomething(c chan string) { //playing around with channels for no reason
	for {
		msg := <-c
		log.Println(msg)
	}
}
