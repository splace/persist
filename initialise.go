package persist

import "os"

var thisFile *os.File

func init() {
	if file, err := os.Open(os.Args[0]);err == nil {
		thisFile=file
		}else{
		panic(err)
	}
}

