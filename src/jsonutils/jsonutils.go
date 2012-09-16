package jsonutils

import (
	"regexp"
	"net/http"
)

// Converts 
//  I am
//  a multiline 
//  string
// to
//  "I am\n"+
//  "a multiline\n"+
//  "string"
func ConvertMultilineToJSONString(s string) string {
		r := regexp.MustCompile("\"")
		ooo := r.ReplaceAllString(s, "\\\"")
		r = regexp.MustCompile("\n")
		ooo = r.ReplaceAllString(ooo, "\\n\"+\n\"")	
		return  "\""+ooo+"\"";
}

func SetJsonHeader (w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type",  "application/javascript; charset=utf-8")
}
