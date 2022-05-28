package tools

import (
	"bytes"
	"fmt"
	"net"
	"regexp"
	"sync"

	"github.com/dyden/go-tcp-server/config"
	"github.com/dyden/go-tcp-server/global"
)

//GetFistWord returns the first word of a string
func GetFistWord(message []byte) []byte {
	r := regexp.MustCompile(`\w+`)
	data := r.Find(message)
	return data
}

//HandleAvaiblesMessage handle avaiblables messages sent by the client
func HandleMessage(message []byte, channels <-chan int, conn net.Conn, wg *sync.RWMutex, wg2 *sync.WaitGroup) {
	fmt.Println("Handling message: ", string(message))
	defer wg2.Done()
	//CHANNELS
	occurrencesChannels := make(chan int, config.MAX_OCCURRENCES_HANDLERS)
	statsChannels := make(chan int, config.MAX_STATS_HANDLERS)
	//WAIT GROUPS
	var wgOccurrences, wgStats sync.WaitGroup

	switch {
	case bytes.HasPrefix(message, global.Data):
		occurrencesChannels <- 1
		wgOccurrences.Add(1)
		go getAllOccurrences(message, wg, &wgOccurrences, occurrencesChannels)
	case bytes.HasPrefix(message, global.Cmd):
		statsChannels <- 1
		wgStats.Add(1)
		getStats(conn, wg, &wgStats, statsChannels)
	default:
		fmt.Println("UNKNOWN: ", string(message))
	}
	wgOccurrences.Wait()
	wgStats.Wait()
	<-channels
}

func getAllOccurrences(message []byte, wg *sync.RWMutex, wgOccurrences *sync.WaitGroup, occurrencesChannels <-chan int) {
	fmt.Println("Getting occurrences of: ", string(message))
	defer wgOccurrences.Done()
	occurrences := global.AvailableOcurrences.FindAll(message, -1)
	for _, test := range occurrences {
		global.AddToStats(string(test), wg)
	}
	<-occurrencesChannels
}

func getStats(conn net.Conn, wg *sync.RWMutex, wgStats *sync.WaitGroup, statsChannels <-chan int) {
	defer wgStats.Done()
	for k, v := range global.GetStats(wg) {
		conn.Write([]byte(fmt.Sprintf("%s: %d\n", k, v)))
	}
	<-statsChannels
}
