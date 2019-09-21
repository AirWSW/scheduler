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

func (n *Node) RunTCPServer() error {
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
		} else {
			var req requset
			if err := json.NewDecoder(connIn).Decode(&req); err != nil {
				log.Error(err)
				// return err
			}
			log.Debugf("Got request message: %s", req.String())

			resp := response{
				ID: "123456",
				Name: "123456",
				Type: "responseType",
				Massage: "responseMessage",
			}
			if json.NewEncoder(connIn).Encode(&resp); err != nil {
				log.Error(err)
				// return err
			}
			log.Debugf("Send response message: %s", resp.String())

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
