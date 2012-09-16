GoLibrary
=========

Some Golang packages.

Run go doc [packagename] for the latest docs on the packages. Or browse the source.

Overview
--------

dirparser        - parses a directory and display its files and directories as JSON
dirparser/server - serves the above with a jsonp interface, including serving the content of the files

fileutils        - basic file utilities, getting the basename easily, find if a file is under a path etc

flagutils        - get a command line argument or exist showing an error message

jsonutils        - convert a multiline string into json formatted multiline string, add the jsonp http header

server           - generic server soon to be removed since it's been imported into dirparser/server

