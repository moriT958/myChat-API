package config

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	cfg  *config
	once sync.Once
)

type config struct {
	Version      string `json:"Version"`
	Address      string `json:"Address"`
	ReadTimeout  int64  `json:"ReadTimeout"`
	WriteTimeout int64  `json:"WriteTimeout"`
}

func Load(filename string) error {
	var err error

	once.Do(func() {
		cfg = new(config)

		file, openErr := os.Open(filename)
		if openErr != nil {
			err = openErr
			return
		}
		defer file.Close()

		if decodeErr := json.NewDecoder(file).Decode(&cfg); decodeErr != nil {
			err = decodeErr
			return
		}
	})

	return err
}

func Version() string {
	return cfg.Version
}

func Address() string {
	return cfg.Address
}

func ReadTimeout() int64 {
	return cfg.ReadTimeout
}

func WriteTimeout() int64 {
	return cfg.WriteTimeout
}
