package app

import (
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
}

type IApiServer interface {
	Start() error
	configLogger() error
	configRouter()
}

func (s *ApiServer) configLogger() error {
	s.logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: true,
	})

	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configLogger(); err != nil {
		return nil
	}

	s.configRouter()

	s.logger.Info("Starting api server...")

	// working
	return nil
}

func (s *ApiServer) configRouter() {
	panic("implement me")
}
