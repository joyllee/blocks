package redis

type Config struct {
	Addresses  []string
	PoolSize   int
	ClientType string // "" or "cluster" or "failover"
	MasterName string
	DB         int
	Password   string
}
