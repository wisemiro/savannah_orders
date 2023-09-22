package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Conf struct {
	Server struct {
		Address string
		Timeout int
	}
	Database struct {
		Address  string
		Port     string
		Name     string
		User     string
		Password string
	}
	Sms struct {
		Key  string
		Name string
		Code string
	}
	Security struct {
		Key string
	}
}

func ReadConfiguration(loc string) (Conf, error) {
	var config Conf
	if _, err := toml.DecodeFile(loc, &config); err != nil {
		return config, errors.Wrap(err, "Error reading config.toml")
	}
	return config, nil
}
