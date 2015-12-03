package config

import "github.com/BurntSushi/toml"

//global Config variable
var Conf *Config

//Config this is a config structure
type Config struct {
	//[server]
	Server struct {
		Bind string `toml:"bind"`
	} `toml:"server"`

	//[rsa]
	RSA struct {
		Passphrase  string `toml:"passphrase"`
		PublicPath  string `toml:"public_key"`
		PrivatePath string `toml:"private_key"`
	} `toml:"rsa"`

	//[jwt]
	JWT struct {
		SecretKey string `toml:"jwt_key"`
		MaxAge    int    `toml:"max_age"`
		Path      string `toml:"path"`
		Domain    string `toml:"domain"`
		Secure    bool   `toml:"secure"`
	} `toml:"jwt"`

	//[db]
	DB struct {
		Database string   `toml:"database"`
		Hosts    []string `toml:"hosts"`
		Username string   `toml:"username"`
		Password string   `toml:"password"`
	} `toml:"db"`

	//[file_upload]
	FileUpload struct {
		MaxSize int64  `toml:"max_size"`
		Temp    string `toml:"temp"`
		Bundle  string `toml:"bundle"`
	} `toml:"file_upload"`
}

//New read a configuration file and returns a Config object
func New(configFile string, confEnv string) (*Config, error) {
	config := &Config{}

	if configFile == "" {
		configFile = confEnv
	}

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	Conf = config

	return config, nil
}
