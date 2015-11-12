package releasifier

import (
	"github.com/alinz/releasifier/config"
	"github.com/alinz/releasifier/data"
)

//Global App pointer
var App *Releasifier

//Releasifier main structuire for releasifier's app
type Releasifier struct {
	Config *config.Config
}

//Start starts Releasifier App, listeing to specified port
func (r *Releasifier) Start() {

}

//Exit stops the app
func (r *Releasifier) Exit() {
	data.DB.Close()
}

//New makes a new and setup releasifer app's settings
func New(conf *config.Config) (*Releasifier, error) {
	//make sure that App is replaced properly.
	if App != nil {
		App.Exit()
	}

	app := &Releasifier{Config: conf}

	//Start a new DB session
	_, err := data.NewDBWithConfig(conf)
	if err != nil {
		panic(err)
	}

	App = app
	return app, nil
}
