package log

import (
	"context"
	"errors"

	"github.com/Christian-Placencia/0231673_SistemasDistribuidos/lab_2_commit_log/api/v1" // Cambiar por la ruta correcta en tu proyecto
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Asegura que grpcServer implemente api.LogServer
var _ api.LogServer = (*grpcServer)(nil)

// Definimos el servidor gRPC que implementa api.LogServer
type grpcServer struct {
	api.UnimplementedLogServer
	*CommitLog
}

// newgrpcServer inicializa un nuevo servidor gRPC con un CommitLog
func newgrpcServer(commitlog *CommitLog) (srv *grpcServer, err error) {
	srv = &grpcServer{
		CommitLog: commitlog,
	}
	return srv, nil
}

// Produce recibe un ProduceRequest y agrega un record al log
func (s *grpcServer) Produce(ctx context.Context, req *api.ProduceRequest) (*api.ProduceResponse, error) {
	offset, err := s.CommitLog.Append(req.Record)
	if err != nil {
		return nil, err
	}
	return &api.ProduceResponse{Offset: offset}, nil
}

// Consume recibe un ConsumeRequest y devuelve el record correspondiente del log
func (s *grpcServer) Consume(ctx context.Context, req *api.ConsumeRequest) (*api.ConsumeResponse, error) {
	record, err := s.CommitLog.Read(req.Offset)
	if err != nil {
		if errors.Is(err, ErrOffsetOutOfRange) {
			return nil, status.Error(codes.OutOfRange, "offset out of range")
		}
		return nil, err
	}
	return &api.ConsumeResponse{Record: record}, nil
}

// ProduceStream maneja un stream bidireccional para producir records en el log
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

// ConsumeStream maneja un stream para consumir records a partir de un offset
func (s *grpcServer) ConsumeStream(req *api.ConsumeRequest, stream api.Log_ConsumeStreamServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			res, err := s.Consume(stream.Context(), req)
			if err != nil {
				if status.Code(err) == codes.OutOfRange {
					continue
				}
				return err
			}
			if err = stream.Send(res); err != nil {
				return err
			}
			req.Offset++
		}
	}
}

// NewGRPCServer inicializa el servidor gRPC
func NewGRPCServer(config *Config) (*grpc.Server, error) {
	gsrv := grpc.NewServer()
	srv, err := newgrpcServer(config.CommitLog)
	if err != nil {
		return nil, err
	}
	api.RegisterLogServer(gsrv, srv)
	return gsrv, nil
}
