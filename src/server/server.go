package server

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"fileutils"
)

type ContentFunction func() string

type ServerInfo struct {
	IndexFileLocation string
	ContentFunction ContentFunction
	Port string
}

func (serverInfo *ServerInfo) StartServer() {
    var c chan string = make(chan string) //playing around with channels for no reason
    go checkForSomething(c)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("REQUESTED: " + r.URL.RequestURI())
		log.Println("REQUESTED: " + r.URL.Path)

		if r.URL.RequestURI() == "/index" {
			f, _ := os.Open(serverInfo.IndexFileLocation)
			of := fileutils.FileUtilType{*f}
			content := of.GetContentOfFile()
			fmt.Fprintf(w, "%s", string(content))
		} else if r.URL.RequestURI() == "/" {
			j := serverInfo.ContentFunction()
			fmt.Fprintf(w, "%s", j)
		} else if r.URL.RequestURI() == "/quit" {
			os.Exit(0)
		} else if r.URL.Path == "/s" { //playing around with channels for no reason
			c <- r.URL.RequestURI()
			fmt.Fprintf(w, "hiii")
		}
	})
	log.Println("Starting server...")
	http.ListenAndServe(":"+serverInfo.Port, nil)
}

func checkForSomething(c chan string) { //playing around with channels for no reason
	for {
		msg := <- c
		log.Println(msg)
	}
}