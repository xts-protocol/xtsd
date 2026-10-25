package metrics

import "sync"

var (
	mu    sync.RWMutex
	Usage = map[string]uint64{}
)

func Increment(apiKey string) {
	mu.Lock()
	defer mu.Unlock()

	Usage[apiKey]++
}

func GetUsage() map[string]uint64 {
	mu.RLock()
	defer mu.RUnlock()

	result := make(map[string]uint64)

	for k, v := range Usage {
		result[k] = v
	}

	return result
}
