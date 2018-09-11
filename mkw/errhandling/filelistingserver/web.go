package main

import (
	"net/http"
	"os"
	"github.com/gpmgo/gopm/modules/log"
	_ "net/http/pprof"//http://localhost:9191/debug/pprof/
	"awesomeProject/mkw/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Warn("Panic :%v", err)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request :%s", err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err): //404
				code = http.StatusNotFound
			case os.IsPermission(err): //403
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError //500
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}
func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":9191", nil)
	if err != nil {
		panic(err)
	}
}
