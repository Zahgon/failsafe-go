package pbfixtures

import (
	"context"
)

type MockPingServer struct {
	UnimplementedPingServiceServer
	OnPing func(context.Context, *PingRequest) (*PingResponse, error)
}

func (m *MockPingServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
