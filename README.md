
  * flag - command line flags
  
 example:
 
 	...
 	flag.Parse()
	persist.FlagSet(flag.CommandLine)  // flags parsed above are completed using stored attribs, and updated/added as needed.
	...
