package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"os"
	"time"
)

func main() {
	//getting writer type
	lwType, err := getWriter()
	if err != nil {
		err = fmt.Errorf("failed to get Writer: %w", err)
		fmt.Println(err)
		return
	}
	//choosing correspondent writer
	var lw logWriter
	switch lwType {
	case 0:
		lw = logToNullWriter{}
	case 1:
		lw = logToFileWriter{file: os.Getenv("FILENAME")}
	}

	fmt.Printf("lw: %#v\n", lw)
	//generating logs and writing them to the writer of choice
	logGenerator(lw)
}

func logGenerator(lw logWriter) {
	fakeLog, err := logFaker()
	if err != nil {
		err = fmt.Errorf("failed to generate logs: %w", err)
		fmt.Println(err)
		return
	}
	lw.Write(fakeLog)
}

func logFaker() (fakeLog string, err error) {
	ut, err := getUptime()
	if err != nil {
		err = fmt.Errorf("failed to launch logs timer: %w", err)
		fmt.Println(err)
		return
	}

	rate, err := getRate()
	if err != nil {
		err = fmt.Errorf(
			"cannot start logFaker: %w", err,
		)
		return "", err
	}

	lineSize, err := getLineSize()
	if err != nil {
		err = fmt.Errorf(
			"cannot start lineFaker: %w", err,
		)
		fmt.Println(err) //TODO: replace w/ return statement
		return "", err
	}

	ticker := time.NewTicker(time.Second / time.Duration(rate))

	alive := true
	go func() {
		time.Sleep(time.Second * time.Duration(ut))
		alive = false
	}()

	for range ticker.C {
		if !alive {
			break
		}
		fakeLog += lineFaker(lineSize) + "\n" //TODO: replace with writing to a []byte buffer
	}
	return fakeLog, nil
}

func lineFaker(lineSize int64) (line string) {

	gofakeit.Seed(time.Now().UnixNano())

	var ip, httpMethod, path, httpVersion, referrer, userAgent string
	var statusCode, bodyBytesSent int
	var timeLocal time.Time

	httpVersion = "HTTP/1.1"
	referrer = "-"

	timeLocal = time.Now()

	ip = gofakeit.IPv4Address()
	httpMethod = weightedHTTPMethod(50, 20)
	path = randomPath(1, 5)
	statusCode = weightedStatusCode(80)
	bodyBytesSent = realisticBytesSent(statusCode)
	userAgent = gofakeit.UserAgent()

	for i := 0; int64(i) < lineSize; i++ {
		lineBit := fmt.Sprintf("%s - - [%s] \"%s %s %s\" %v %v \"%s\" \"%s\"", ip,
			timeLocal.Format("02/Jan/2006:15:04:05 -0700"), httpMethod,
			path, httpVersion, statusCode, bodyBytesSent, referrer, userAgent)
		line += lineBit //TODO: get rid of concatenation to optimize performance
	}
	return
}
