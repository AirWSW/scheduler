package main

import (
	// "fmt"
	"encoding/json"

	"github.com/op/go-logging"

	"github.com/AirWSW/scheduler/config"
	"github.com/AirWSW/scheduler/logger"
	"github.com/AirWSW/scheduler/node"
	"github.com/AirWSW/scheduler/slave"
	"github.com/AirWSW/scheduler/master"
)

var log = logging.MustGetLogger("schedule")

func init()  {
	logger.InitLogger()
	log.Debug("...importing the scheduler.go")
}

func main()  {
	log.Info("Starting scheduler...1.2.3.")
	// log.Infof(bigBanner, version, thisYear)
	log.Infof(bigBanner, version)

	// Load configuration from configStr
	log.Info("Load configuration from configStr...")
	cfg := &config.Config{}
	if err := cfg.LoadConfig(configStr); err != nil {
		log.Error(err)
		panic(err)
	}
	log.Debugf("Scheduler configuration: %s", cfg.String())

	// Load configuration to node
	log.Info("Load configuration to node...")
	n := node.Node{}
	if err := n.LoadConfig(cfg); err != nil {
		log.Error(err)
		panic(err)
	}
	log.Debugf("Node: %s", n.String())

	// Load configuration to schedule
	log.Info("Load configuration to schedule...")
	plans, err := master.LoadPlanConfig(cfg)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	bytes, _ := json.Marshal(plans)
	log.Debugf("Plans: %s", string(bytes))

	// Assign the node to a certain role
	s := slave.Slave{}
	m := master.Master{}
	log.Info("Assign node to a certain role...")
	if n.Role == "slave" {
		s.Node = &n
		log.Debugf("Slave: %s", s.String())
		if err := s.Run(); err != nil {
			log.Error(err)
			panic(err)
		}
	} else {
		m.Node = &n
		if n.Role == "master" || n.Role == "master-slave" {
			m.Schedule = plans
		}
		log.Debugf("Master: %s", m.String())
		if err := m.Run(); err != nil {
			log.Error(err)
			panic(err)
		}
	}
}