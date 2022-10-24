package discovery

import (
	"google.golang.org/grpc/resolver"
)

const (
	EtcdScheme = "etcd"
)

func init() {
	resolver.Register(&etcdResolverBuilder{})
}

type etcdResolverBuilder struct {
}

func (*etcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	de := getDiscEtcd(target)
	de.loadServiceListenOn()

	r := &etcdResolver{
		cc: cc,
	}

	update := func() {
		addrStrs := de.addrs
		addrs := make([]resolver.Address, len(addrStrs))
		for i, s := range addrStrs {
			addrs[i] = resolver.Address{Addr: s}
		}
		r.cc.UpdateState(resolver.State{Addresses: addrs})
	}

	go de.monitorChange(update)
	update()

	return r, nil
}
func (*etcdResolverBuilder) Scheme() string {
	return EtcdScheme
}

type etcdResolver struct {
	cc resolver.ClientConn
}

func (*etcdResolver) ResolveNow(o resolver.ResolveNowOptions) {
}
func (*etcdResolver) Close() {
}
