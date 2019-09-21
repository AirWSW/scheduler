package config

import (
	// "errors"
	// "fmt"
	// "os"
	"encoding/json"
	// "path/filepath"
	
	"github.com/op/go-logging"
	"github.com/BurntSushi/toml"

	"github.com/AirWSW/scheduler/logger"
)

var log = logging.MustGetLogger("schedule")

// Config represents a genernal config file
type Config struct {
	Node     nodeConfig   `toml:"node"`
	Schedule []planConfig `toml:"schedule"`
	Tasks    []taskConfig `toml:"tasks"`
}

type nodeConfig struct {
	ID    string `toml:"id"`
	Name  string `toml:"name"`
	Role  string `toml:"role"`
	Debug bool   `toml:"debug"`
	
	Master serverConfig `toml:"master"`
	Server serverConfig `toml:"server"`

	Token tokenConfig `toml:"token"`
}

type serverConfig struct {
	Hostname string `toml:"hostname"`
	Address  string `toml:"address"`
	Port     int    `toml:"port"`
}

type tokenConfig struct {
	AccessToken   string `toml:"access_token"`
	RefreshToken  string `toml:"refresh_token"`
	RegisterToken string `toml:"register_token"`
}

type planConfig struct {
	ID       string `toml:"id"`
	Name     string `toml:"name"`
	NodeID   string `toml:"node_id"`
	Interval int    `toml:"interval"`
	
	Task     taskConfig `toml:"task"`

	Status   string     `toml:"status"`
}

type taskConfig struct {
	ID     string `toml:"id"`
	Name   string `toml:"name"`
	Status string `toml:"status"`
}

func init()  {
	logger.InitLogger()
	log.Debug("...importing the config.go")
}

// CheckConfig checks configuration's legality.
func (c *Config) CheckConfig() (error) {

	return nil
}

// LoadConfig loads configuration from string.
func (c *Config) LoadConfig(s string) (error) {
	// Read config anyway
	if _, err := toml.Decode(s, c); err != nil {
		return err
	}
  // Set DEBUG flag from config
	if c.Node.Debug {
		logging.SetLevel(logging.DEBUG, "")
	}
	// Check KEY configuration's legality
	if err := c.CheckConfig(); err != nil {
		return err
	}
	return nil
}

// Marshal returns the JSON encoding of a node info.
func (c *Config) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Config) String() string {
	bytes, err := c.Marshal()
	if err != nil {
		return string(bytes)
	}
	return string(bytes)
}

// LoadConfig loads configuration
// func LoadConfig() (*Config, error) {
// 	cfg := new(Config)
// 	err := cfg.LoadConfig()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return cfg, nil
// }
