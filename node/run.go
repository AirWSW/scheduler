package node

import (
	"fmt"
	// "strconv"
	"time"
	// "math/rand"
	"net"
	"net/http"
	// "flag"
	// "strings"
	"encoding/json"

	// "github.com/op/go-logging"

	// "github.com/AirWSW/scheduler/logger"
	// "github.com/AirWSW/scheduler/config"
)

// RunTCPServer runs a server to listen tcp request
func (n *Node) RunTCPServer(masterCh chan Requset) error {
	log.Debug("Run TCP Server...")
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", n.Server.Address, n.Server.Port))
	if err != nil {
		return err
	}

	for {
		connIn, err := ln.Accept()
		if err != nil {
			// if _, ok := err.(net.Error); ok {
			// 	fmt.Println("Error received while listening.", me.NodeId)
			// }
			log.Error(err)
		} else {
			// Receive any tcp request from ADDRESS:PORT
			var req Requset
			if err := json.NewDecoder(connIn).Decode(&req); err != nil {
				// Throw out invalid requset
				log.Error(err)
			} else {
				// Throw out invalid requset
				log.Debugf("Got request message: %s", req.String())

				masterCh <- req
				// Do something
			
				resp := response{
					ID: req.ID,
					Name: req.Name,
					Type: req.Type,
					Massage: "OK",
				}
				if json.NewEncoder(connIn).Encode(&resp); err != nil {
					log.Error(err)
				} else {
					log.Debugf("Send response message: %s", resp.String())
				}
			}
			connIn.Close()
		}
	}
}

func (n *Node) RunHTTPServer() {
	addr := fmt.Sprintf("%s:%d", n.Server.Address, n.Server.Port)

	httpServer := &http.Server{
		Addr:         addr,
		// Handler:      s.engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
