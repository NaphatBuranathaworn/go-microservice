package counter

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	endpoint "github.com/NaphatBuranathaworn/go-microservice/pkg/counter/endpoints"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoint.Set) http.Handler {
    m := http.NewServeMux()
    m.Handle("/add", httptransport.NewServer(
        ep.AddEndpoint,
        decodeHTTPServiceAddRequest,
        encodeResponse,
    ))

    return m
}

func decodeHTTPServiceAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req endpoint.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}