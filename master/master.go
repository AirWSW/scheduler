package master

import (
	// "fmt"
	"encoding/json"

	"github.com/op/go-logging"

	"github.com/AirWSW/scheduler/logger"
	"github.com/AirWSW/scheduler/node"
	"github.com/AirWSW/scheduler/config"
	"github.com/AirWSW/scheduler/slave"
)

var log = logging.MustGetLogger("schedule")

type Master struct {
	Node     *node.Node   `json:"node"`
	Cluster  []node.Node  `json:"cluster"`
	Schedule []PlanConfig `json:"schedule"`
	Tasks    []slave.Task `json:"tasks"`
}

type PlanConfig struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	NodeID string `json:"node_id"`

	Task slave.Task `json:"task"`

	Status string `json:"status"`
}

func init()  {
	logger.InitLogger()
	log.Debug("...importing the master.go")
}

// LoadPlanConfig load node config to an existed node.
func LoadPlanConfig(cfg *config.Config) ([]PlanConfig, error) {
	plans := []PlanConfig{}
	for _, p := range cfg.Schedule {
		plan := PlanConfig{
			ID: p.ID,
			Name: p.Name,
			NodeID: p.NodeID,
			Task: slave.Task{
				ID: p.Task.ID,
				Name: p.Task.Name,
				Status: p.Task.Status,
			},
			Status: p.Status,
		}
		plans = append(plans, plan)
	}
	return plans, nil
}

// Marshal returns the JSON encoding of a node info.
func (m *Master) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Master) String() string {
	bytes, err := m.Marshal()
	if err != nil {
		return string(bytes)
	}
	return string(bytes)
}

func (m *Master) Run() error {
	// go m.Node.RunTCPServer()
	go func () {
		switch m.Node.Role {
		case "master": m.RunAsMaster()
		case "master-standby": m.RunAsMasterStandby()
		case "master-slave": m.RunAsMasterSlave()
		case "slave-master": m.RunAsSlaveMaster()
		}
	}()
	if err := m.Node.RunTCPServer(); err != nil {
		return err
	}
	return nil
}

func (m *Master) RunAsMaster() error {
	log.Debug("Run As Master Only")
	return nil
}

func (m *Master) RunAsMasterStandby() error {
	log.Debug("Run As Master Standby")
	return nil
}

func (m *Master) RunAsMasterSlave() error {
	log.Debug("Run As Master Slave")
	m.Node.RegisterRequest()
	
	return nil
}

func (m *Master) RunAsSlaveMaster() error {
	log.Debug("Run As Slave Master")
	return nil
}


// func main()  {
// 	// main func
// 	cfg, err := config.LoadConfig()
// 	n := node.Node{}

// 	err = n.LoadConfig(cfg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = n.RunTCPServer(); err != nil {
// 		panic(err)
// 	}
	
// }