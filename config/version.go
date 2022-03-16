package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	PROJECT_NAME = ""
	PROJECT_URL  = ""
	VERSION      = ""
	COMMIT_SHA1  = ""
	BUILD_TIME   = ""
)

func Usage() {
	fmt.Fprintf(os.Stderr, `%s
%s
Version: %s
Build Time: %s
Commit SHA1: %s
Usage: %s -h -c -m -v

Options:
`, PROJECT_NAME, PROJECT_URL, VERSION, BUILD_TIME, COMMIT_SHA1, PROJECT_NAME)
	flag.PrintDefaults()
}
