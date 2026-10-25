package metrics

import (
	"sync"

	"github.com/xts-protocol/xtsd/internal/storage"
)

var (
	mu    sync.RWMutex
	Usage = map[string]uint64{}
)

func Load() {
	data, err := storage.LoadUsage()
	if err == nil {
		Usage = data
	}
}

func Increment(apiKey string) {
	mu.Lock()
	defer mu.Unlock()

	Usage[apiKey]++

	_ = storage.SaveUsage(Usage)
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
