package receipt

import (
	"encoding/json"
	"os"
	"sync"
)

const ReceiptFile = "data/receipts.json"

var (
	mu       sync.RWMutex
	Receipts = map[string]Receipt{}
)

func Load() {

	file, err := os.Open(ReceiptFile)
	if err != nil {
		return
	}

	defer file.Close()

	var data map[string]Receipt

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return
	}

	Receipts = data
}

func Save(r Receipt) {

	mu.Lock()
	defer mu.Unlock()

	Receipts[r.MessageID] = r

	file, err := os.Create(ReceiptFile)
	if err != nil {
		return
	}

	defer file.Close()

	_ = json.NewEncoder(file).Encode(Receipts)
}

func Get(id string) (Receipt, bool) {

	mu.RLock()
	defer mu.RUnlock()

	r, ok := Receipts[id]

	return r, ok
}
