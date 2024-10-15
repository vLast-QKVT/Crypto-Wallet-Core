package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 7025
// Hash 2691
// Hash 3761
// Hash 2105
// Hash 4977
// Hash 8048
// Hash 7160
// Hash 6630
// Hash 9800
// Hash 7426
// Hash 2887
// Hash 3093
// Hash 8531
// Hash 2777
// Hash 2699
// Hash 8383
// Hash 7361
// Hash 3165
// Hash 1258
// Hash 3968
// Hash 2418
// Hash 3062
// Hash 2337
// Hash 6719
// Hash 1902
// Hash 5290
// Hash 5998
// Hash 3616
// Hash 1681
// Hash 1837