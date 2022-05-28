package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"sync"

	"github.com/dyden/go-tcp-server/config"
	"github.com/dyden/go-tcp-server/global"
	"github.com/dyden/go-tcp-server/services"
	"github.com/dyden/go-tcp-server/tools"
)

func HandleRequest(conn *net.Conn, wg *sync.RWMutex) {
	fmt.Println("New connection request")
	c := make(chan int, config.MAX_MESSAGES_HANDLERS)
	b := bufio.NewReader(*conn)
	defer tools.CloseConnection(*conn)
	for {
		message, err := b.ReadBytes('\n')
		if err != nil {
			tools.HandleError(err)
			return
		}
		firstWord := tools.GetFistWord(message)
		if bytes.Equal(firstWord, global.Stop) {
			tools.CloseConnection(*conn)
			return
		}

		go services.HandleMessage(message, c, conn, wg)
		c <- 1
	}

}
