package endpoints

import (
	"context"

	"github.com/NaphatBuranathaworn/go-microservice/pkg/database"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
    GetEndpoint endpoint.Endpoint
    AddEndpoint endpoint.Endpoint
}

func NewEndPointSet(svc database.Service) Set {
    return Set{
        GetEndpoint: MakeGetEndpoint(svc),
    }
}

func MakeGetEndpoint(svc database.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetRequest)
        docs, err := svc.Get(ctx, req.Filters...)
        if err != nil {
            return GetResponse{docs, err.Error()}, nil
        }
        return GetResponse{docs, ""}, nil
    }
}

func MakeAddEndpoint(svc database.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		ticketID, err := svc.Add(ctx, req.Document)
		if err != nil {
			return AddResponse{TicketID: ticketID, Err: err.Error()}, nil
		}
		return AddResponse{TicketID: ticketID, Err: ""}, nil
	}
}
