package tenant

import "sync"

var (
	mu sync.RWMutex
)

func Create(t Tenant) {

	mu.Lock()
	defer mu.Unlock()

	Registry[t.APIKey] = t

	_ = SaveRegistry()
}

func Get(apiKey string) (Tenant, bool) {
	mu.RLock()
	defer mu.RUnlock()

	t, ok := Registry[apiKey]

	return t, ok
}
