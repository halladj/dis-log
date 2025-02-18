package loadbalance

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"

	api "github.com/halladj/dis-log/api/v1"
)

type Resolver struct {
	mu           sync.Mutex
	clientConn   resolver.ClientConn
	resolverConn *grpc.ClientConn
	serverConfig *serviceconfig.ParseResult
	logger       *zap.Logger
}

// Build implements resolver.Builder.
func (r *Resolver) Build(
	target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions,
) (resolver.Resolver, error) {

	r.logger = zap.L().Named("resolver")
	r.clientConn = cc
	var dialOpts []grpc.DialOption
	if opts.DialCreds != nil {
		dialOpts = append(
			dialOpts,
			grpc.WithTransportCredentials(
				opts.DialCreds),
		)
	}

	r.serverConfig = r.clientConn.ParseServiceConfig(
		fmt.Sprintf(`{"loadBalancingConfig:[{"%s":{}}]"}`, Name),
	)

	var err error
	r.resolverConn, err = grpc.NewClient(
		target.Endpoint(),
		dialOpts...,
	)

	if err != nil {
		return nil, err
	}

	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

const Name = "dis-log"

func (r *Resolver) Scheme() string {
	return Name
}

func init() {
	resolver.Register(&Resolver{})
}

var _ resolver.Builder = (*Resolver)(nil)

var _ resolver.Resolver = (*Resolver)(nil)

// Close implements resolver.Resolver.
func (r *Resolver) Close() {
	if err := r.resolverConn.Close(); err != nil {
		r.logger.Error(
			"failed to close conn",
			zap.Error(err),
		)
	}
}

// ResolveNow implements resolver.Resolver.
func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {
	r.mu.Lock()
	defer r.mu.Unlock()

	client := api.NewLogClient(r.resolverConn)

	//gets cluster & sets cc.
	ctx := context.Background()
	res, err := client.GetServers(
		ctx, &api.GetServersRequest{},
	)
	if err != nil {
		r.logger.Error(
			"failed to resolve server",
			zap.Error(err),
		)
		return
	}

	var addrs []resolver.Address
	for _, server := range res.Servers {
		addrs = append(addrs, resolver.Address{
			Addr: server.RpcAddr,
			Attributes: attributes.New(
				"is_leader",
				server.IsLeader,
			),
		})
	}

	r.clientConn.UpdateState(resolver.State{
		Addresses:     addrs,
		ServiceConfig: r.serverConfig,
	})

}
