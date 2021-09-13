package main

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
)

func isASCIISpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

// TrimString returns s without leading and trailing ASCII space.
func TrimString(s string) string {
	for len(s) > 0 && isASCIISpace(s[0]) {
		s = s[1:]
	}
	for len(s) > 0 && isASCIISpace(s[len(s)-1]) {
		s = s[:len(s)-1]
	}
	return s
}

func readWordsFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Warnf("LocalWordLoad %s fail : %s", fileName, err.Error())
		return nil, err
	}

	defer file.Close()

	fi, _ := file.Stat()

	//主要是为性能考虑，append []string
	resp := make([]string, 0, int(fi.Size()))

	scanner := bufio.NewScanner(file)
	//注意: optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//把不需要的字符过滤掉
		text := TrimString(scanner.Text())
		resp = append(resp, text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return resp, nil
}

func PanicRecovery(quit bool) {
	var err error
	if r := recover(); r != nil {
		switch x := r.(type) {
		case string:
			err = errors.New(x)
			break
		case error:
			err = x
			break
		default:
			err = errors.New("Unknown panic")
		}

		debug.PrintStack()
		log.Warn(string(debug.Stack()))
		log.Infoln("Panic :", err.Error())

		if quit {
			os.Exit(101)
		}
	}
}
