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
// Hash 8058
// Hash 1871
// Hash 2963
// Hash 8854
// Hash 5429
// Hash 3961
// Hash 2836
// Hash 2441
// Hash 3507
// Hash 4637
// Hash 1455
// Hash 1560
// Hash 3926
// Hash 5469
// Hash 9167
// Hash 3503
// Hash 6908
// Hash 8954
// Hash 2275
// Hash 9918
// Hash 7238
// Hash 4741
// Hash 7588
// Hash 6084
// Hash 7338
// Hash 8117
// Hash 9482
// Hash 4092
// Hash 3965
// Hash 4643
// Hash 1831
// Hash 7817
// Hash 2009
// Hash 7971
// Hash 7579
// Hash 7173
// Hash 3896
// Hash 1381
// Hash 2094
// Hash 7955
// Hash 4965
// Hash 3955
// Hash 6482
// Hash 7769
// Hash 7710
// Hash 6661
// Hash 7320
// Hash 5065
// Hash 3904
// Hash 8136
// Hash 3332
// Hash 7858