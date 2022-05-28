package services

import (
	"fmt"
	"net"
	"sync"

	"github.com/dyden/go-tcp-server/global"
	"github.com/dyden/go-tcp-server/tools"
)

//HandleRequest handles the message sent by the client
func HandleMessage(data []byte, c <-chan int, conn *net.Conn, wg *sync.RWMutex) {
	fmt.Println("New message from ", (*conn).RemoteAddr())
	channels := make(chan int, 1)
	availableMessages := global.AvailableMessageRegexp.FindAll(data, -1)
	var wgHandleMessage sync.WaitGroup
	for _, message := range availableMessages {
		channels <- 1
		wgHandleMessage.Add(1)
		go tools.HandleMessage(message, channels, *conn, wg, &wgHandleMessage)
	}
	wgHandleMessage.Wait()
	<-c

}
