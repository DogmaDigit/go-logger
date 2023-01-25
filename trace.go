package go_logger

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// инициализация логгера, создание необходимого файла
func (gl *GoLogger) loggerWriter(content string, level string) {
	if len(gl.FolderPath) == 0 {
		return
	}
	if err := checkFolder(gl.FolderPath); err != nil {
		if err := createFolder(gl.FolderPath); err != nil {
			log.Fatal("error create logs folder")
		}
	}
	fileName := level + "." + "log"
	file, err := os.OpenFile(gl.FolderPath+"/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatal("error opening file: ", err)
	}
	file.WriteString(content + "\n")
	file.Close()
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func getTrace(content string, err error) []string {
	if err == nil {
		return []string{}
	}
	traceErr := errors.Wrap(err, content)
	trace := []string{}
	if err, ok := traceErr.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			trace = append(trace, fmt.Sprintf("%+s:%d", f, f))
		}
	}
	return trace
}
