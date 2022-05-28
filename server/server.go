package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	config "github.com/dyden/go-tcp-server/config"
	"github.com/dyden/go-tcp-server/handler"
	"github.com/dyden/go-tcp-server/tools"
	"golang.org/x/net/netutil"
)

//Start starts the server
func Start() {
	var wg sync.RWMutex

	fmt.Println("Starting server...")
	//LOAD CONFIG
	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Error loading config: ", err)
		os.Exit(1)
	}
	//SERVER LISTEN ON PORT...
	l, err := net.Listen(config.Type, config.Host+":"+config.Port)
	if err != nil {
		tools.HandleError(err)
		os.Exit(1)
	}
	//LIMIT CONNECTIONS
	l = netutil.LimitListener(l, config.MaxConnections)
	//ERROR HANDLING
	if err != nil {
		tools.HandleError(err)
		os.Exit(1)
	}
	//CLOSE LISTENER
	defer l.Close()
	//MESSAGE OF LISTENING ON PORT
	fmt.Println("Listening on " + config.Host + ":" + config.Port)
	//LISTEN FOR REQUESTS
	for {
		conn, err := l.Accept()
		if err != nil {
			tools.HandleError(err)
			return
		}
		fmt.Println("NEW CONNECTION -> ", conn.RemoteAddr())
		conn.Write([]byte("Connected to server.\nType 'STOP' to close connection\n"))
		//HANDLE REQUEST
		go handler.HandleRequest(&conn, &wg)
	}
}
