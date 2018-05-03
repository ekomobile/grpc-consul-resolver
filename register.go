package resolver

import (
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
)

func RegisterDefault(watchInterval time.Duration) {
	resolver.Register(&ConsulResolverBuilder{
		WatchInterval:      watchInterval,
		ConsulClientConfig: api.DefaultConfig(),
	})
}
