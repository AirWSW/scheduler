package connection

import (
	// "fmt"
	"net"
)

type Conn struct {
	ConnIn  *net.TCPConn
	ConnOut *net.TCPConn
}

