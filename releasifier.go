package releasifier

import (
	"time"

	"github.com/alinz/releasifier/config"
	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/logme"
	"github.com/alinz/releasifier/web"
	"github.com/alinz/releasifier/web/security"
	"github.com/tylerb/graceful"
)

//Releasifier main structuire for releasifier's app
type Releasifier struct {
	conf *config.Config
}

//Start starts Releasifier App, listeing to specified port
func (r *Releasifier) Start() {
	graceful.Run(r.conf.Server.Bind, 10*time.Second, web.New())
}

//Exit stops the app
func (r *Releasifier) Exit() {
	data.DB.Close()
}

//New makes a new and setup releasifer app's settings
func New(conf *config.Config) (*Releasifier, error) {
	logme.Info("Releasifier started at " + conf.Server.Bind)

	app := &Releasifier{conf: conf}

	//setup security
	security.Setup(conf)

	//set SecureIDKey from config
	data.SetSecureIDKey(conf.AES.SecureKey)

	//Start a new DB session
	_, err := data.NewDBWithConfig(conf)
	if err != nil {
		logme.Fatal(err)
	}

	return app, nil
}
