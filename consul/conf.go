package consul

type Config struct {
	Address    string
	Scheme     string
	DataCenter string
	Token      string
	TokenFile  string
}
