package persist

import "testing"
import "os"
import "io/ioutil"
import "log"
//import "encoding/gob"

type d struct{
	A int
	B int
}

func TestSaveLoad(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "xattrib_gob_test")
	defer os.Remove(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	d1:=d{A:1,B:2}
	SaveBinary(d1,"D1")
	d2:=LoadBinary("D1").(d)
	if d2.A!=1 || d2.B!=2 {
		t.Error("loaded back not the same as saved.")
	}
	
}

