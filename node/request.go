package node

import (
	"fmt"
	// "strconv"
	"time"
	// "math/rand"
	"net"
	// "net/http"
	// "flag"
	// "strings"
	"encoding/json"

	// "github.com/op/go-logging"

	// "github.com/AirWSW/scheduler/logger"
	// "github.com/AirWSW/scheduler/config"
)

type requset struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Node Node   `json:"node"`

  Massage interface{} `json:"massage"`
}

// Marshal returns the JSON encoding of a node info.
func (r *requset) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *requset) String() string {
	bytes, err := r.Marshal()
	if err != nil {
		return string(bytes)
	}
	return string(bytes)
}

func (n *Node) Request(requestType string, requestMessage string) (error){
	HERE:
	connOut, err := net.DialTimeout(
		"tcp", 
		fmt.Sprintf("%s:%d", n.Master.Address, n.Master.Port),
		time.Duration(10) * time.Second,
	)
	if err != nil {
		log.Error(err)
		goto HERE
		return err
	} else {
		req := requset{
			ID: "123456",
			Name: "123456",
			Type: requestType,
			Node: Node{
				ID: n.ID,
				Name: n.Name,
				Role: n.Role,
				Debug: n.Debug,
				
				Master: serverConfig{
					Hostname: n.Master.Hostname,
					Address: n.Master.Address,
					Port: n.Master.Port,
				},
				Server: serverConfig{
					Hostname: n.Server.Hostname,
					Address: n.Server.Address,
					Port: n.Server.Port,
				},
		
				Token: tokenConfig{
					AccessToken: n.Token.AccessToken,
					RefreshToken: n.Token.RefreshToken,
					RegisterToken: n.Token.RegisterToken,
				},
			},
			Massage: requestMessage,
		}
		if err := json.NewEncoder(connOut).Encode(&req); err != nil {
			log.Error(err)
			// return err
		}
		log.Debugf("Send request message: %s", req.String())

		var resp response
		if err := json.NewDecoder(connOut).Decode(&resp); err != nil {
			log.Error(err)
			// return err
		}
		log.Debugf("Got response message: %s", resp.String())

		return nil
	}
}

func (n *Node) RegisterRequest() (string, error) {
	n.Request("RegisterToCluster", "RegisterRequest")
	bytes, err := n.Marshal()
	return string(bytes), err
}

func (n *Node) KeepAliveRequest() (string, error) {
	bytes, err := n.Marshal()
	return string(bytes), err
}

func (n *Node) UpdateRequest() (string, error) {
	bytes, err := n.Marshal()
	return string(bytes), err
}

func (n *Node) DelistRequest() (string, error) {
	bytes, err := n.Marshal()
	return string(bytes), err
}
