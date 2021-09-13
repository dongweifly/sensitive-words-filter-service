package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	VERSION = 0.1

	port    = flag.Int("p", 8088, "http server port")
	dictDir = flag.String("d", "./data", "sensitive words file path")
)

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
	if err := GMatchService.Init(*dictDir); err != nil {
		log.Warningf("Service start fail %s", err.Error())
		return
	}
	httpServerRun(*port)
}
