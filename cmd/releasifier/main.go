package main

import (
	"flag"
	"os"

	"github.com/alinz/releasifier"
	"github.com/alinz/releasifier/config"
	"github.com/alinz/releasifier/logme"
)

var (
	flags    = flag.NewFlagSet("releasifier", flag.ExitOnError)
	confFile = flags.String("config", "", "path to config file")
)

func main() {
	flags.Parse(os.Args[1:])

	var err error
	var conf *config.Config

	//load configuration from either confFile or Env's CONFIG variable
	conf, err = config.New(*confFile, os.Getenv("CONFIG"))
	if err != nil {
		logme.Fatal(err)
	}

	//create a new Releasidier app.
	app, err := releasifier.New(conf)
	if err != nil {
		logme.Fatal(err)
	}

	//start the Releasifier's App.
	//this will block until app stops, either by panic or exit signal
	app.Start()

	logme.Info("App is shutting down.")
	app.Exit()
}
