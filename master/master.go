package master

import (
	// "fmt"
	"time"
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

	Interval int       `json:"interval"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	
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
			Interval: p.Interval,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
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

// Marshal returns the JSON encoding of a node info.
func (p *PlanConfig) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p *PlanConfig) String() string {
	bytes, err := p.Marshal()
	if err != nil {
		return string(bytes)
	}
	return string(bytes)
}

func (m *Master) Run() error {
	// go m.Node.RunTCPServer()
	masterCh := make(chan node.Requset, 20)
	go func () {
		switch m.Node.Role {
		case "master": m.RunAsMaster()
		case "master-standby": m.RunAsMasterStandby()
		case "master-slave": m.RunAsMasterSlave(masterCh)
		case "slave-master": m.RunAsSlaveMaster()
		}
	}()
	var err error
	if err = m.Node.RunTCPServer(masterCh); err != nil {
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

func (m *Master) RunAsMasterSlave(masterCh chan node.Requset) error {
	log.Debug("Run As Master Slave")
	// Register Request
	go RegisterToCluster(m)

	// master for for forever
	go MasterForLoop(m)
	
	// salve for for forever
	go SlaveForLoop(m)
	
	for v := range masterCh {
		log.Debug(v.String())
		switch v.Type {
			case "RegisterToCluster": 
			if err := m.AddNodeToCluster(&v.Node); err != nil {
				log.Error(err)
			}
			case "DelistFromCluster": 
			if err := m.RemoveNodeFromCluster(&v.Node); err != nil {
				log.Error(err)
			}
			case "AddTaskToNode": 
			plan := PlanConfig{}
			if err := json.Unmarshal([]byte(v.Massage.(string)), &plan); err != nil {
				log.Error(err)
				// return err
			}
			if err := m.AddTaskToNode(&plan); err != nil {
				log.Error(err)
			}
			case "UpdateTaskInfo": 
			t := slave.Task{}
			if err := json.Unmarshal([]byte(v.Massage.(string)), &t); err != nil {
				log.Error(err)
				// return err
			}
			if err := m.UpdateTaskInfo(&v.Node, &t); err != nil {
				log.Error(err)
			}
		}
		log.Debug(m.String())
	}
	return nil
}

func (m *Master) RunAsSlaveMaster() error {
	log.Debug("Run As Slave Master")
	return nil
}

func RegisterToCluster(m *Master) {
	m.Node.RegisterRequest()
	m.Node.DelistRequest()
	m.Node.RegisterRequest()
}

func MasterForLoop(m *Master) {
	for {
		for _, c := range m.Cluster {
			for j, s := range m.Schedule {
				if s.Status == "enabled" && s.NodeID == c.ID {
					m.Node.AddTaskRequest(s.String())
					s.Status = "start"
					log.Infof("Scheduling task: %s", s.String())
					m.Schedule[j] = s
				} else if s.Status == "wait" && s.NodeID == c.ID {
					if s.UpdateAt.Add(time.Duration(s.Interval)*time.Second).Before(time.Now()) {
						m.Node.AddTaskRequest(s.String())
						s.Status = "start"
						log.Infof("Scheduling task: %s", s.String())
						m.Schedule[j] = s
					}
				}
			}
		}
	}
}

func SlaveForLoop(m *Master) {
	for {
		HERE:
		for i, t := range m.Tasks {
			if t.Status == "enabled" {
				t.Status = "start"
				m.Node.UpdateTaskRequest(t.String())
				log.Infof("Starting task: %s", t.String())
				m.Tasks[i] = t
			} else if t.Status == "start" {
				t.Status = "success"
				m.Node.UpdateTaskRequest(t.String())
				log.Infof("Task success: %s", t.String())
				m.Tasks[i] = t
			} else if t.Status == "success" {
				t.Status = "delete"
				m.Node.UpdateTaskRequest(t.String())
				log.Infof("Delete task: %s", t.String())
				m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
				goto HERE
			}
		}
	}
}