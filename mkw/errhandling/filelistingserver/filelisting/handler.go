package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	index := strings.Index(request.URL.Path, prefix)
	if index != 0 {
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
