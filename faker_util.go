package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"os"
	"strconv"
	"strings"
)

func getLineSize() (lineSize int64, err error) {
	lineSizeString := os.Getenv("LINE_SIZE")
	fmt.Printf("LINE_SIZE: %v\n", lineSizeString)
	lineSize, err = strconv.ParseInt(lineSizeString, 0, 32)
	if err != nil {
		err = fmt.Errorf(
			"failed parsing env. variable 'LINE_SIZE' into an int: %w", err,
		)
		return 0, err
	}
	if lineSize <= 0 {
		err = fmt.Errorf("LINE_SIZE cannot be non-positive")
	}
	return lineSize, err
}

func getRate() (rate float64, err error) {
	rateString := os.Getenv("RATE")
	fmt.Printf("RATE: %v\n", rateString)
	rate, err = strconv.ParseFloat(rateString, 32)
	if err != nil {
		err = fmt.Errorf(
			"failed parsing env. variable 'RATE' into a float: %w", err,
		)
		return 0, err
	}
	if rate <= 0 {
		err = fmt.Errorf("RATE cannot be non-positive")
	}
	return rate, err
}

func getUptime() (uptime int64, err error) {
	uptimeString := os.Getenv("UPTIME")
	fmt.Printf("UPTIME: %v\n", uptimeString)
	uptime, err = strconv.ParseInt(uptimeString, 0, 32)
	if err != nil {
		err = fmt.Errorf(
			"failed parsing env. variable 'UPTIME' into an int: %w", err,
		)
		return 0, err
	}
	if uptime <= 0 {
		err = fmt.Errorf("UPTIME cannot be non-positive")
	}
	return uptime, err
}

func getWriter() (writerType int, err error) {
	writerString := os.Getenv("WRITER")
	fmt.Printf("WRITER: %v\n", writerString)
	switch writerString {
	case "null":
		writerType = 0
	case "file":
		writerType = 1
	default:
		err = fmt.Errorf("WRITER is not selected")
		writerType = 0
	}
	return
}

func realisticBytesSent(statusCode int) int {
	if statusCode != 200 {
		return gofakeit.Number(30, 120)
	}

	return gofakeit.Number(800, 3100)
}

func weightedStatusCode(percentageOk int) int {
	roll := gofakeit.Number(0, 100)
	if roll <= percentageOk {
		return 200
	}

	return gofakeit.SimpleStatusCode()
}

func weightedHTTPMethod(percentageGet, percentagePost int) string {
	if percentageGet+percentagePost > 100 {
		panic("percentageGet and percentagePost add up to more than 100%")
	}

	roll := gofakeit.Number(0, 100)
	if roll <= percentageGet {
		return "GET"
	} else if roll <= percentagePost {
		return "POST"
	}

	return gofakeit.HTTPMethod()
}

func randomPath(min, max int) string {
	var path strings.Builder
	length := gofakeit.Number(min, max)

	path.WriteString("/")

	for i := 0; i < length; i++ {
		if i > 0 {
			path.WriteString(gofakeit.RandString([]string{"-", "-", "_", "%20", "/", "/", "/"}))
		}
		path.WriteString(gofakeit.BuzzWord())
	}

	path.WriteString(gofakeit.RandString([]string{".html", ".php", ".htm", ".jpg", ".png", ".gif", ".svg", ".css",
		".js"}))

	result := path.String()
	return strings.Replace(result, " ", "%20", -1)
}
