package main

import (
	"flag"
	"fmt"
	"gintest/common"
	"gintest/config"
	"gintest/routers"
	"gintest/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

var (
	help       bool
	version    bool
	configfile string
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.StringVar(&configfile, "c", "config.yml", "config file path\n")
	flag.Usage = config.Usage

	flag.Parse()
	if help {
		config.Usage()
		os.Exit(0)
	}
	if version {
		fmt.Println(config.PROJECT_NAME + "/" + config.COMMIT_SHA1)
		os.Exit(0)
	}
}

func main() {
	initFlag()
	if err := configor.Load(&common.CConfig, configfile); err != nil {
		panic(err)
	}

	utils.InitLogger(common.CConfig.Log)
	utils.InitDB(common.CConfig.Mysql)
	defer utils.CloseDb()

	gin.SetMode(common.CConfig.Env)
	r := gin.New()
	r.Use(gin.Recovery())

	routers.Setup(r)
	r.Run(common.CConfig.Port)
}
