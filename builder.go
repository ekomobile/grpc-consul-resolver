package resolver

import (
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
)

type ConsulResolverBuilder struct {
	WatchInterval      time.Duration
	ConsulClientConfig *api.Config
}

func (b *ConsulResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	consul, err := api.NewClient(b.ConsulClientConfig)
	if err != nil {
		return nil, err
	}

	r := ConsulResolver{
		target:        target,
		cc:            cc,
		consul:        consul,
		addr:          make(chan []resolver.Address, 1),
		done:          make(chan struct{}, 1),
		watchInterval: b.WatchInterval,
	}

	go r.updater()
	go r.watcher()
	r.resolve()

	return &r, nil
}

func (b *ConsulResolverBuilder) Scheme() string {
	return Scheme
}
