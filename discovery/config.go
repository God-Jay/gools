package discovery

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Conf struct {
	Service     Service     `yaml:"Service"`
	Etcd        EtcdConf    `yaml:"Etcd"`
	RpcServices RpcServices `yaml:"RpcServices"`
}
type Service struct {
	Name     string `yaml:"Name"`
	ListenOn string `yaml:"ListenOn"`
}
type EtcdConf struct {
	Endpoints []string `yaml:"Endpoints"`
	// TODO
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}
type RpcServices map[string]string

func ResolveConf(configFile string) (*Conf, error) {
	var conf Conf
	content, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &conf)
	return &conf, err
}

func MustResolveConf(configFile string) *Conf {
	conf, err := ResolveConf(configFile)
	if err != nil {
		log.Fatalf("error resolve config file %s, %s", configFile, err.Error())
	}
	return conf
}
