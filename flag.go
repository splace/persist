package persist

import "os"
import "github.com/splace/os/xattribs"
import "flag"

var userFlagsFile xattribs.FileNS

func init() {
	thisFile, err := os.Open(os.Args[0])
	if err != nil {
		panic(err)
	}
	userFlagsFile = xattribs.FileNS{thisFile, "user.flags"}
}

// FlagSet persists/completes the passed flag.FlagSet
func FlagSet(fs *flag.FlagSet) {
	fs.VisitAll(func(f *flag.Flag) {
		bs, err := userFlagsFile.Get(f.Name)
		if err == nil {
			f.Value.Set(string(bs))
		}else{
			if f.Value.String()!=f.DefValue {
				userFlagsFile.Set(f.Name, []byte(f.Value.String()))
			}
		}
	})
}
