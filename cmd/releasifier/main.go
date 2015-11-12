package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alinz/releasifier/config"
)

var (
	flags    = flag.NewFlagSet("releasifier", flag.ExitOnError)
	confFile = flag.String("config", "", "path to config file")
)

func main() {
	flags.Parse(os.Args[1:])

	var err error
	var conf *config.Config

	conf, err = config.New(*confFile, os.Getenv("CONFIG"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", conf)
}
