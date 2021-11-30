package redis

type Config struct {
	Addresses []string
	PoolSize  int
	DB        int
	Password  string
}
