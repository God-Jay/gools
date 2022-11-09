package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ResolveYaml(configFile string, conf interface{}) error {
	content, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, conf)
	return err
}

func MustResolveYaml(configFile string, conf interface{}) {
	if err := ResolveYaml(configFile, conf); err != nil {
		log.Fatalf("err resolve config file: %s, %v", configFile, err)
	}
}
