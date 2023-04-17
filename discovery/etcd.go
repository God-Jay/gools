package discovery

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/resolver"
	"strings"
	"time"
)

const (
	Delimiter         = "/"
	dialTimeout       = 5 * time.Second
	dialKeepAliveTime = 5 * time.Second
)

// servers use register struct
type regEtcd struct {
	*clientv3.Client
	leaseID     clientv3.LeaseID
	serviceName string
	addr        string
}

// clients use discovery struct
type discEtcd struct {
	*clientv3.Client
	serviceName string
	addrs       []string
}

func newEtcdClient(endpoints string) (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            strings.Split(endpoints, ","),
		DialTimeout:          dialTimeout,
		DialKeepAliveTime:    dialKeepAliveTime,
		DialKeepAliveTimeout: dialKeepAliveTime,
	})
	go monitorEtcdConnState(client)
	return client, err
}
func monitorEtcdConnState(client *clientv3.Client) {
	cc := client.ActiveConnection()
	for {
		state := cc.GetState()
		switch state {
		case connectivity.Idle:
			cc.Connect()
		}
		if !cc.WaitForStateChange(client.Ctx(), state) {
			return
		}
	}
}

func getRegEtcd(endpoints string, serviceName string, listenOn string) *regEtcd {
	client, err := newEtcdClient(endpoints)
	if err != nil {
		// TODO err
	}
	lease, err := client.Grant(client.Ctx(), 5)
	if err != nil {
		// handle error!
	}
	return &regEtcd{Client: client, leaseID: lease.ID, serviceName: serviceName, addr: listenOn}
}

// store the service to etcd
func (re *regEtcd) storeService() {
	serviceName := makeEtcdKey(re.serviceName, int64(re.leaseID))
	_, err := re.Put(re.Ctx(), serviceName, re.addr,
		clientv3.WithLease(re.leaseID))
	if err != nil {
		// handle error!
		panic(err)
	}
}

func (re *regEtcd) keepalive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	return re.KeepAlive(re.Ctx(), re.leaseID)
}

func (re *regEtcd) revoke() {
	re.Revoke(re.Ctx(), re.leaseID)
}

func getDiscEtcd(target resolver.Target) *discEtcd {
	client, err := newEtcdClient(target.URL.Host)
	if err != nil {
		// TODO err
	}
	return &discEtcd{Client: client, serviceName: strings.Trim(target.URL.Path, Delimiter)}
}

func (de *discEtcd) loadServiceListenOn() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	resp, err := de.Get(ctx, makeKeyPrefix(de.serviceName), clientv3.WithPrefix())
	if err != nil {
		// TODO
	}

	cancel()

	var kvs []string
	for _, kv := range resp.Kvs {
		kvs = append(kvs, string(kv.Value))
	}
	de.addrs = kvs
}

func (de *discEtcd) monitorChange(update func()) {
	rch := de.Client.Watch(context.Background(), makeKeyPrefix(de.serviceName), clientv3.WithPrefix())
	for range rch {
		de.loadServiceListenOn()
		update()
		//for _, ev := range wresp.Events {
		//	fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		//}
	}
}

func makeEtcdKey(key string, id int64) string {
	return fmt.Sprintf("%s%s%d", key, Delimiter, id)
}
func makeKeyPrefix(key string) string {
	return fmt.Sprintf("%s%s", key, Delimiter)
}
