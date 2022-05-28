package global

import (
	"fmt"
	"sync"
)

var (
	stats = make(map[string]int)
)

func GetStats(wg *sync.RWMutex) (allStats map[string]int) {
	fmt.Println("Reading stats")
	wg.RLock()
	allStats = stats
	wg.RUnlock()
	return
}

func AddToStats(key string, wg *sync.RWMutex) {
	fmt.Println("Adding to stats")
	wg.Lock()
	stats[key]++
	wg.Unlock()
}
