package main

import (
	"flag"
	"fmt"
	match "github.com/dongweifly/sensitive-words-match"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"path/filepath"

	"os"
)

var (
	VERSION = 0.1

	port    = flag.Int("p", 8088, "http server port")
	dictDir = flag.String("d", "./data", "sensitive words file path")

	MatchService = match.NewMatchService()
)

func InitMatchService(dictDir string) error {
	var files []string
	//只关注rootDir下面的文件，暂时不支持子文件夹的形式
	err := filepath.Walk(dictDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Warnf("load %s fail %s", dictDir, err.Error())
			return nil
		}

		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})

	if err != nil {
		return err
	}

	var words []string
	for _, file := range files {
		if w, err := readWordsFromFile(dictDir + "/" + file); err == nil {
			log.Infof("load senvitive files  %s words size %d", dictDir+"/"+file, len(w))
			words = append(words, w...)
		} else {
			return err
		}
	}

	MatchService.Build(words)
	return nil
}

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			fmt.Println("version: ", VERSION)
			os.Exit(0)
		}

		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			flag.Usage()
			os.Exit(0)
		}
	}
	log.SetOutput(os.Stdout)
	gin.SetMode(gin.ReleaseMode)

	log.Info("Use sensitive words DIR ", *dictDir)
	if err := InitMatchService(*dictDir); err != nil {
		fmt.Println("init sensitive dict error: ", err.Error())
	}

	httpServerRun(*port)
}
