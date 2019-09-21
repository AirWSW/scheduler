package master

import (
	"errors"
	"time"
	// "fmt"

	"github.com/AirWSW/scheduler/node"
	"github.com/AirWSW/scheduler/slave"
)

func (m *Master) AddNodeToCluster(n *node.Node) error {
	for _, v := range m.Cluster {
		if v.ID == n.ID {
			return errors.New("Node existed: " + n.String())
		}
	}
	m.Cluster = append(m.Cluster, *n)
	log.Infof("Add node to cluster: %s", n.String())

	return nil
}

func (m *Master) RemoveNodeFromCluster(n *node.Node) error {
	for i, v := range m.Cluster {
		if v.ID == n.ID {
			m.Cluster = append(m.Cluster[:i], m.Cluster[i+1:]...)
			log.Infof("Remove node from cluster: %s", n.String())
			return nil
		}
	}
	return errors.New("Node not existed: " + n.String())
}

func (m *Master) AddTaskToNode(t *PlanConfig) error {
	for _, v := range m.Tasks {
		if v.ID == t.Task.ID {
			return errors.New("Tasks existed: " + t.String())
		}
	}
	m.Tasks = append(m.Tasks, t.Task)
	log.Infof("Add task to node: %s", t.String())

	return nil
}

func (m *Master) UpdateTaskInfo(n *node.Node, t *slave.Task) error {
	for i, s := range m.Schedule {
		if s.NodeID == n.ID && s.Task.ID == t.ID {
			s.Task = *t
			s.UpdateAt = time.Now()
			if s.Task.Status == "delete" {
				s.Task.Status = "enabled"
				s.Status = "wait"
			}
			log.Infof("Updated task info: %s", s.String())
			m.Schedule[i] = s
			return nil
		}
	}
	return errors.New("Task not existed: " + t.String())
}
