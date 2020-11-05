package main

import (
	"fmt"
	"os"
)

type logWriter interface {
	Write(fakeLog string)
}

type logToNullWriter struct{}

func (lw logToNullWriter) Write(fakeLog string) {
	_ = fakeLog
	return
}

type logToFileWriter struct {
	file string
}

func (lw logToFileWriter) Write(fakeLog string) {
	fmt.Println(lw.file)
	f, err := os.OpenFile(lw.file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		err = fmt.Errorf(
			"failed opening a file '%v': %w", lw.file, err,
		)
	}
	defer f.Close()

	fakeLogBytes := []byte(fakeLog)
	_, err = f.Write(fakeLogBytes)
	if err != nil {
		err = fmt.Errorf(
			"failed writing to a file: %w", err,
		)
	}
	return
}
