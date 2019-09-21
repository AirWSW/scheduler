package node

import (
	"encoding/json"

	"github.com/op/go-logging"

	"github.com/AirWSW/scheduler/logger"
	"github.com/AirWSW/scheduler/config"
)

var log = logging.MustGetLogger("schedule")

// Node represents a genernal node info
type Node struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Debug bool   `json:"debug"`
	
	Master serverConfig `json:"master"`
	Server serverConfig `json:"server"`

	Token tokenConfig `json:"token"`
}

type serverConfig struct {
	Hostname string `json:"hostname"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
}

type tokenConfig struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	RegisterToken string `json:"register_token"`
}

func init()  {
	logger.InitLogger()
	log.Debug("...importing the node.go")
}

// LoadConfig load node config to an existed node.
func (n *Node) LoadConfig(cfg *config.Config) error {
	*n = Node{
		ID: cfg.Node.ID,
		Name: cfg.Node.Name,
		Role: cfg.Node.Role,
		Debug: cfg.Node.Debug,
		
		Master: serverConfig{
			Hostname: cfg.Node.Master.Hostname,
			Address: cfg.Node.Master.Address,
			Port: cfg.Node.Master.Port,
		},
		Server: serverConfig{
			Hostname: cfg.Node.Server.Hostname,
			Address: cfg.Node.Server.Address,
			Port: cfg.Node.Server.Port,
		},

		Token: tokenConfig{
			AccessToken: cfg.Node.Token.AccessToken,
			RefreshToken: cfg.Node.Token.RefreshToken,
			RegisterToken: cfg.Node.Token.RegisterToken,
		},
	}
	return nil
}

// Marshal returns the JSON encoding of a node info.
func (n *Node) Marshal() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Node) String() string {
	bytes, err := n.Marshal()
	if err != nil {
		return string(bytes)
	}
	return string(bytes)
}
