package discovery

import (
	"fmt"
	"strings"
)

// RegService register the current server service to etcd
func RegService(conf *Conf) {
	keepalive(conf)
}

func keepalive(conf *Conf) {
	re := getRegEtcd(strings.Join(conf.Etcd.Endpoints, ","), conf.Service.Name, figureOutListenOn(conf.Service.ListenOn))
	re.storeService()

	ch, kaerr := re.keepalive()
	if kaerr != nil {
		// handle error!
	}
	go func() {
		for {
			select {
			case ka, ok := <-ch:
				if !ok {
					re.revoke()
					keepalive(conf)
					return
				} else {
					// TODO
					if ka != nil {
						fmt.Println("ttl:", ka.TTL)
					} else {
						fmt.Println("Unexpected NULL")
					}
				}
			}
		}
	}()
}
