package target

import (
	"context"
	"fmt"

	"github.com/iptecharch/data-server/config"
	"github.com/iptecharch/data-server/schema"
	"google.golang.org/grpc"

	sdcpb "github.com/iptecharch/sdc-protos/sdcpb"
)

type Target interface {
	Get(ctx context.Context, req *sdcpb.GetDataRequest) (*sdcpb.GetDataResponse, error)
	Set(ctx context.Context, req *sdcpb.SetDataRequest) (*sdcpb.SetDataResponse, error)
	Subscribe()
	//
	Sync(ctx context.Context, syncConfig *config.Sync, syncCh chan *SyncUpdate)
}

func New(ctx context.Context, name string, cfg *config.SBI, schemaClient schema.Client, schema *sdcpb.Schema, opts ...grpc.DialOption) (Target, error) {
	switch cfg.Type {
	case "gnmi":
		return newGNMITarget(ctx, name, cfg, opts...)
	case "nc":
		return newNCTarget(ctx, name, cfg, schemaClient, schema)
	case "redis":
		return newRedisTarget(ctx, cfg)
	case "nats":
		return newNATSTarget(ctx, cfg)
	case "noop", "":
		return newNoopTarget(ctx, name)
	}
	return nil, fmt.Errorf("unknown DS target type %q", cfg.Type)
}

type SyncUpdate struct {
	Tree   string
	Update *sdcpb.Notification
}
