package node

import (
	// "fmt"
	// "strconv"
	// "time"
	// "math/rand"
	// "net"
	// "net/http"
	// "flag"
	// "strings"
	"encoding/json"

	// "github.com/op/go-logging"

	// "github.com/AirWSW/scheduler/logger"
	// "github.com/AirWSW/scheduler/config"
)

type response struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

  Massage interface{} `json:"massage"`
}

// Marshal returns the JSON encoding of a node info.
func (r *response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *response) String() string {
	bytes, err := r.Marshal()
	if err != nil {
		return string(bytes)
	}
	return string(bytes)
}

// func (n *Node) Response(responseType string, responseMessage string) (error){
// 	HERE:
// 	connOut, err := net.DialTimeout(
// 		"tcp", 
// 		fmt.Sprintf("%s:%d", n.Master.Address, n.Master.Port),
// 		time.Duration(10) * time.Second,
// 	)
// 	if err != nil {
// 		log.Error(err)
// 		goto HERE
// 		return err
// 	} else {
// 		var resp response
// 		if err := json.NewDecoder(connOut).Decode(&resp); err != nil {
// 			log.Error(err)
// 			// return err
// 		}
// 		log.Debugf("Got response message: %s", resp.String())

// 		return nil
// 	}
// }

func (n *Node) RegisterResponse() (string, error) {
	n.SendRequest("RegisterToCluster", "RegisterRequest")
	bytes, err := n.Marshal()
	return string(bytes), err
}
