package slave

import (
	// "errors"
	// "fmt"
	// "os/exec"
	// "net/http"
	// "time"
	"encoding/json"

	"github.com/op/go-logging"

	"github.com/AirWSW/scheduler/logger"
	"github.com/AirWSW/scheduler/node"
)

var log = logging.MustGetLogger("schedule")

type Slave struct {
	Node    *node.Node  `json:"node"`
	Cluster []node.Node `json:"cluster"`
	Tasks   []Task      `json:"tasks"`
}

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func init()  {
	logger.InitLogger()
	log.Debug("...importing the slave.go")
}

// Marshal returns the JSON encoding of a node info.
func (s *Slave) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Slave) String() string {
	bytes, err := s.Marshal()
	if err != nil {
		return string(bytes)
	}
	return string(bytes)
}

func (s *Slave) Run() error {
	log.Debug("Run As Slave Only")
	if err := s.Node.RunTCPServer(); err != nil {
		return err
	}
	s.Node.RegisterRequest()
	return nil
}