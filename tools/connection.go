package tools

import (
	"net"
)

//CloseConnection close connection of server
func CloseConnection(conn net.Conn) {
	conn.Write([]byte("Closing connection...\n"))
	conn.Close()
}
