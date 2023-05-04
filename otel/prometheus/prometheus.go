package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ServeAddr string `yaml:"ServeAddr"`
}

func ResolveConf(config string) (*Config, error) {
	var conf Config
	content, err := os.ReadFile(config)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &conf)
	return &conf, err
}

func (c *Config) withDefault() {
	if c.ServeAddr == "" {
		c.ServeAddr = ":9100"
	}
}

func Simple() {
	// Create non-global registry.
	reg := prometheus.NewRegistry()

	// Add go runtime metrics and process collectors.
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}
