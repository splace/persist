package persist


import "github.com/splace/os/xattribs"
import "encoding/gob"
import "bytes"
import "log"

// need to register the type in interface{}
func SaveBinary(fs interface{},name string) {
	var buf bytes.Buffer
	gob.Register(fs)
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&fs)
	if err != nil {
		log.Fatal("encode:", err)
	}
	userFlagsFile := xattribs.FileNS{thisFile, "user.gobs"}
	userFlagsFile.Set(name, buf.Bytes())
}

func LoadBinary(name string)(fs interface{}) {
	userFlagsFile := xattribs.FileNS{thisFile, "user.gobs"}
	bs, err := userFlagsFile.Get(name)
	if err != nil {
		log.Fatal("decode:", err)
	}
	dec := gob.NewDecoder(bytes.NewBuffer(bs))
	err = dec.Decode(&fs)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return
}


