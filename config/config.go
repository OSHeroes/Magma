package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var (
	Config cfg
)

type cfg struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"DEBUG"`

	HTTPBind  string `envconfig:"HTTP_BIND" default:"127.0.0.1:80"`
	HTTPSBind string `envconfig:"HTTPS_BIND" default:"127.0.0.1:443"`

	CertificatePath string `envconfig:"CERTIFICATE_PATH" required:"true"`
	PrivateKeyPath  string `envconfig:"PRIVATE_KEY_PATH" required:"true"`
}

func LoadToMemory() {
	if err := envconfig.Process("", &Config); err != nil {
		logrus.WithError(err).Fatal("Initialize configs")
	}
}

// LogLevel parses a default log level from a string
func LogLevel() logrus.Level {
	lvl, err := logrus.ParseLevel(Config.LogLevel)
	if err != nil {
		logrus.WithError(err).Fatal("config: Parse log level")
	}
	return lvl
}
