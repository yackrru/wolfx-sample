package main

import (
	"flag"
	"github.com/yackrru/wolfx-sample/cli"
	"os"
)

var (
	jobName = flag.String("job", "", "jobname")
)

func main() {
	flag.Parse()
	os.Exit(cli.Execute(*jobName))
}
