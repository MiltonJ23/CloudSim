package persist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

var lock sync.Mutex

var Marshall = func(o interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		fmt.Println("An error happened during the JSON Marshalling:", err.Error())
		return nil, err
	}
	return bytes.NewReader(b), nil
}
var Unmarshall = func(read io.Reader, v interface{}) error {
	return json.NewDecoder(read).Decode(v)
}

func Save(filepath string, o interface{}) {
	lock.Lock()
	defer lock.Unlock()
	// let's create a file
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println("An error occured while opening the file:", err.Error())
	}
	defer f.Close()
	r, er := Marshall(o)
	if er != nil {
		fmt.Println("An error occured during the JSON Marshalling:", er.Error())
	}
	_, errr := io.Copy(f, r)
	if errr != nil {
		fmt.Println("An error occured during the JSON Write:", errr.Error())
	}
}

func Load(filepath string, o interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("An error occured while opening the file:", err.Error())
	}
	defer f.Close()
	return Unmarshall(f, o)

}
