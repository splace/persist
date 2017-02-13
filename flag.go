package persist

import "github.com/splace/os/xattribs"
import "flag"

// FlagSet persists/completes the passed flag.FlagSet
func FlagSet(fs *flag.FlagSet) {
	userFlagsFile := xattribs.FileNS{thisFile, "user.flags"}
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
