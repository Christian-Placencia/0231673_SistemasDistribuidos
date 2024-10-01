package server

import (
	"context"
	"log"
	"net"

	. "github.com/Christian-Placencia/0231673_SistemasDistribuidos/internal"

	"github.com/Christian-Placencia/0231673_SistemasDistribuidos/api/v1"
	"google.golang.org/grpc"
)

// grpcServer is the implementation of the gRPC server for the Log service
type grpcServer struct {
	api.UnimplementedLogServer
	*Log
	grpcServer *grpc.Server
	listener   net.Listener
}

// NewGRPCServer initializes a new gRPC server with a CommitLog
func NewGRPCServer(commitlog *Log) (*grpcServer, error) {
	srv := &grpcServer{
		Log: commitlog,
	}
	return srv, nil
}

// Serve starts the gRPC server and listens on the given listener
func (s *grpcServer) Serve(l net.Listener) {
	// Assign the listener and create a new gRPC server
	s.listener = l
	s.grpcServer = grpc.NewServer()

	// Register the log service
	api.RegisterLogServer(s.grpcServer, s)

	log.Printf("gRPC server listening on %s", l.Addr().String())

	// Start serving
	if err := s.grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Stop gracefully stops the gRPC server
func (s *grpcServer) Stop() {
	log.Println("Stopping gRPC server...")
	s.grpcServer.GracefulStop()
	if s.listener != nil {
		s.listener.Close()
	}
	log.Println("gRPC server stopped")
}

func (s *grpcServer) Produce(ctx context.Context, req *api.ProduceRequest) (*api.ProduceResponse, error) {
	offset, err := s.Log.Append(req.Record)
	if err != nil {
		return nil, err
	}
	return &api.ProduceResponse{Offset: offset}, nil
}

func (s *grpcServer) Consume(ctx context.Context, req *api.ConsumeRequest) (*api.ConsumeResponse, error) {
	record, err := s.Log.Read(req.Offset)
	if err != nil {
		return nil, err
	}
	return &api.ConsumeResponse{Record: record}, nil
}

func (s *grpcServer) ProduceStream(stream api.Log_ProduceStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		res, err := s.Produce(stream.Context(), req)
		if err != nil {
			return err
		}
		if err = stream.Send(res); err != nil {
			return err
		}
	}
}

func (s *grpcServer) ConsumeStream(req *api.ConsumeRequest, stream api.Log_ConsumeStreamServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			res, err := s.Consume(stream.Context(), req)
			switch err.(type) {
			case nil:
			case api.ErrOffsetOutOfRange:
				continue
			default:
				return err
			}
			if err = stream.Send(res); err != nil {
				return err
			}
			req.Offset++
		}
	}
}
