// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// Config represents a configuration file.
type Config struct {
	filename string
	cache    *map[string]interface{}
}

// DefaultConfig is the config object shared by all modules
var DefaultConfig *Config

// NewConfig returns a new Config object.
func NewConfig(filename string) *Config {
	config := Config{filename, nil}
	config.Reload()
	go config.watch()
	return &config
}

// GetDefaultConfig returns the default config path file
func GetDefaultConfig() *Config {
	if DefaultConfig != nil {
		return DefaultConfig
	}
	return NewConfig("./config.json")
}

func (config *Config) Read(key string) string {
	var v string
	config.Get(key, &v)
	return v
}

// Get retreives a Config option into a passed in pointer or returns an error.
func (config *Config) Get(key string, v interface{}) error {
	var val interface{}

	env, set := os.LookupEnv(key)

	if set {
		switch v.(type) {
		case *float64:
			val, err := strconv.ParseFloat(env, 64)

			if err != nil {
				return err
			}

			*v.(*float64) = val
			return nil
		case *int:
			val, err := strconv.ParseInt(env, 10, 64)

			if err != nil {
				return err
			}

			*v.(*int) = int(val)
			return nil
		default:
			val = env
		}
	} else if config.cache != nil {
		val = (*config.cache)[key]
	}

	// Cast JSON values
	switch v.(type) {
	case *string:
		if val == nil {
			val = ""
		}

		if b, ok := val.(bool); ok {
			*v.(*string) = strconv.FormatBool(b)
		} else if f, ok := val.(float64); ok {
			*v.(*string) = strconv.FormatFloat(f, 'f', -1, 64)
		} else {
			*v.(*string) = val.(string)
		}
	case *bool:
		switch val {
		case nil, 0, false, "", "0", "false":
			// falsey
			val = false
		default:
			// truthy
			val = true
		}

		*v.(*bool) = val.(bool)
	case *float64:
		if val == nil {
			val = float64(0)
		}

		if s, ok := val.(string); ok {
			pf, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return err
			}

			*v.(*float64) = pf
		} else {
			*v.(*float64) = val.(float64)
		}
	case *int:
		if val == nil {
			val = float64(0)
		}

		*v.(*int) = int(val.(float64))
	default:
		return errors.New("Type not supported")
	}

	return nil
}

// Reload clears the config cache.
func (config *Config) Reload() error {
	cache, err := primeCacheFromFile(config.filename)
	config.cache = cache

	if err != nil {
		return err
	}

	return nil
}

func (config *Config) watch() {
	l := log.New(os.Stderr, "", 0)

	// Catch SIGHUP to automatically reload cache
	sighup := make(chan os.Signal, 1)
	signal.Notify(sighup, syscall.SIGHUP)

	for {
		<-sighup
		l.Println("Caught SIGHUP, reloading config...")
		config.Reload()
	}
}

func primeCacheFromFile(file string) (*map[string]interface{}, error) {
	// File exists?
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, err
	}

	// Read file
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal
	var config map[string]interface{}
	if err := json.Unmarshal(raw, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
