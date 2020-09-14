package consul

import (
	"github.com/hashicorp/consul/api"
)

var defaultConsulClient *api.Client

func InitDefault(config Config) (err error) {
	defaultConsulClient, err = New(config)
	if err != nil {
		return err
	}
	return nil
}

func Client() *api.Client {
	return defaultConsulClient
}
