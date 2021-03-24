package counter

import (
	"context"
	"os"

	"github.com/NaphatBuranathaworn/go-microservice/pkg/counter"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)


type Set struct {
    AddEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc counter.Service) Set {
    return Set{
        AddEndpoint: MakeAddEndpoint(svc),
    }   
}

func MakeAddEndpoint(svc counter.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		v := svc.Add(req.V)
		return AddResponse{v}, nil
	}
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}