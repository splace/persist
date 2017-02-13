
#persist

  * flag - command line flags.

example:

	...
	flag.Parse()
	persist.FlagSet(flag.CommandLine)  // added line: default FlagSet has any missing flags filled in with stored values, any included flags are added/updated to the store.
	...
	
installation:

	go get github.com/splace/persist
     

docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/persist?status.svg)](https://godoc.org/github.com/splace/persist)  persist

     


