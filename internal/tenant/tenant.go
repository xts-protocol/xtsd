package tenant

type Tenant struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
}

var Registry = map[string]Tenant{
	"xts_demo_key": {
		ID:     "tenant-001",
		Name:   "XTS Demo",
		APIKey: "xts_demo_key",
	},
}
