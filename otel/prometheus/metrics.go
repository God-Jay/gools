package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
)

var once sync.Once

const PathMetrics = "/metrics"

func ServeMetrics(c *Config) {
	c.withDefault()

	once.Do(func() {
		go func() {
			mux := http.NewServeMux()
			mux.Handle(PathMetrics, promhttp.Handler())
			if err := http.ListenAndServe(c.ServeAddr, mux); err != nil {
				log.Println("failed to serve metrics exporter")
			}
		}()
	})
}
