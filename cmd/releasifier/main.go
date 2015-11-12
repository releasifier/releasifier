package main

import (
	"flag"
	"os"

	"github.com/alinz/releasifier"
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

	//load configuration from either confFile or Env's CONFIG variable
	conf, err = config.New(*confFile, os.Getenv("CONFIG"))
	if err != nil {
		panic(err)
	}

	//create a new Releasidier app.
	_, err = releasifier.New(conf)
	if err != nil {
		panic(err)
	}

	//start the Releasifier's App
	releasifier.App.Start()
}
