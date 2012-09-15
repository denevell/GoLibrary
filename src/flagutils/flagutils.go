//Utility methods for the flag package
package flagutils

import (
	"os"
	"flag"
)

// Like a flag.String but if the flag is "", we exit
// the program altogether
func GetStringFlagOrExit(name string, desc string) string {
	arg := flag.String(name, "", desc)
	flag.Parse()
	if *arg=="" {
		flag.Usage()
		os.Exit(1)
	}
	return *arg
}