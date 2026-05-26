package testutil

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/failsafe-go/failsafe-go/internal/testutil/pbfixtures"
)

type pingService struct {
	pbfixtures.UnimplementedPingServiceServer
	responseFn func(ctx context.Context) (*pbfixtures.PingResponse, error)
}

func (s *pingService) Ping(ctx context.Context, req *pbfixtures.PingRequest) (*pbfixtures.PingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MockGrpcResponses(responses ...string) pbfixtures.PingServiceServer {
	_ = "STUB: not implemented"
	return *new(pbfixtures.PingServiceServer)
}

func MockDelayedGrpcResponse(response string, delay time.Duration) pbfixtures.PingServiceServer {
	_ = "STUB: not implemented"
	return *new(pbfixtures.PingServiceServer)
}

func MockGrpcError(err error) pbfixtures.PingServiceServer {
	_ = "STUB: not implemented"
	return *new(pbfixtures.PingServiceServer)
}

func MockFlakyGrpcServer(failTimes int, err error, finalResponse string) pbfixtures.PingServiceServer {
	_ = "STUB: not implemented"
	return *new(pbfixtures.PingServiceServer)
}

type Dialer func(context.Context, string) (net.Conn, error)

func GrpcServer(service pbfixtures.PingServiceServer, options ...grpc.ServerOption) (*grpc.Server, Dialer) {
	_ = "STUB: not implemented"
	return nil, *new(Dialer)
}

func GrpcClient(dialer Dialer, options ...grpc.DialOption) *grpc.ClientConn {
	_ = "STUB: not implemented"
	return nil
}
