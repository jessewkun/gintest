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
	ginMode    string
)

func initFlag() {
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.StringVar(&configfile, "c", "config.yml", "config file path\n")
	flag.StringVar(&ginMode, "m", "debug", `gin mode
"debug" indicates gin mode is debug.
"release" indicates gin mode is release.
"test" indicates gin mode is test.`+"\n")
	flag.Usage = config.Usage

	flag.Parse()
	if help {
		config.Usage()
		os.Exit(0)
	}
	if version {
		fmt.Println(config.PROJECT_NAME + "/" + config.VERSION)
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

	gin.SetMode(ginMode)
	r := gin.New()
	r.Use(gin.Recovery())

	routers.Setup(r)
	r.Run(common.CConfig.Port)
}
