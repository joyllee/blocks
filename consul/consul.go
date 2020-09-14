package consul

import (
	"github.com/hashicorp/consul/api"
)

func New(conf Config) (*api.Client, error) {
	consulConf := api.DefaultConfig()
	if len(conf.Address) > 0 {
		consulConf.Address = conf.Address
	}
	if len(conf.Scheme) > 0 {
		consulConf.Scheme = conf.Scheme
	}
	if len(conf.DataCenter) > 0 {
		consulConf.Datacenter = conf.DataCenter
	}
	if len(conf.Token) > 0 {
		consulConf.Token = conf.Token
	}
	if len(conf.TokenFile) > 0 {
		consulConf.TokenFile = conf.TokenFile
	}
	return api.NewClient(consulConf)
}
