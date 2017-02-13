
  * flag - command line flags

example:

	...
	flag.Parse()
	persist.FlagSet(flag.CommandLine)  // added line: so flags parsed above are completed using stored attribs, and updated/added as needed.
	...
	
installation:

	go get github.com/splace/persist
     

docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/persist?status.svg)](https://godoc.org/github.com/splace/persist)  persist

     


