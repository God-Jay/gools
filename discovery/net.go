package discovery

import (
	"fmt"
	"net"
	"strings"
)

const allEths = "0.0.0.0"

// figureOutListenOn support listen on:
// ip:port, :port
func figureOutListenOn(listenOn string) string {
	fields := strings.Split(listenOn, ":")
	if len(fields) == 0 {
		return listenOn
	}

	host := fields[0]
	if len(host) > 0 && host != allEths {
		return listenOn
	}

	ip := getIP()

	return strings.Join(append([]string{ip}, fields[1:]...), ":")
}

func getIP() string {
	// TODO use net.Interfaces() ?
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
