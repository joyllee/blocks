package mongo

type Config struct {
	URI    string // mongodb://user:password@localhost:27017/?authSource=admin or mongodb://localhost:27017,localhost:27018,localhost:27018/?replicaSet=rs1
	Dbname string
}

var _Config Config

func Configuration() Config {
	return _Config
}
