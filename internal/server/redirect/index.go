package redirect

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var indexHTML = ReadHTML()

func ReadHTML() []byte {
	_, file, _, _ := runtime.Caller(0)
	index := fmt.Sprintf("%s/%s", filepath.Dir(file), "index.html")
	buf, err := ioutil.ReadFile(index)
	if err != nil {
		panic(err)
	}
	return buf
}
