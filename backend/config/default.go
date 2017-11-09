package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const DEFAULT_CONFIG = `
[server]
host = "0.0.0.0"
port = 8080

[middlewares]

    [middlewares.cors]
    active = true
    [middlewares.logger]
	active = true
	[middlewares.limit]
	active = true
	max = "4MB"

[databases]

    [databases.mysql]
    host = "127.0.0.1"
    port = 3306
    username = "root"
    password = "123456"
    dbname = "scratch"
`

type Config struct {
	Server      serverConfig
	Databases   map[string]databaseConfig
	Middlewares map[string]middlewareConfig
}

type serverConfig struct {
	Host string
	Port int
}

type databaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type middlewareConfig struct {
	Active bool
	Max    string
}

func check(conf *Config) error {
	return nil
}

func Initialize() (Config, error) {

	var conf Config

	_, err := toml.DecodeFile(filepath.Join(filepath.Dir(os.Args[0]), "config/default.conf"), &conf)

	if err == nil {

		log.Print("Custom configutation has been loaded successfully")

	} else {
		_, err := toml.Decode(DEFAULT_CONFIG, &conf)

		if err != nil {
			return Config{}, err
		}

		log.Print("Failed to load custom configuration. Use default")

	}

	err = nil
	err = check(&conf)

	return conf, err

}
