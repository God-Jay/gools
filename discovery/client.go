package discovery

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

type Client struct {
	*grpc.ClientConn
}

// NewClient rpcServiceName is the key of Conf.RpcServices, not the Conf.Service.Name
// use round_robin balancing policy
func NewClient(c *Conf, rpcServiceName string) *Client {
	client, err := grpc.Dial(
		fmt.Sprintf("%s://%s/%s", EtcdScheme, strings.Join(c.Etcd.Endpoints, ","), c.RpcServices[rpcServiceName]),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), // This sets the initial balancing policy.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		// TODO
	}
	return &Client{client}
}
