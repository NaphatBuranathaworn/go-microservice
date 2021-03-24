package counter

import (
	"os"

	"github.com/go-kit/kit/log"
)

type counterService struct{}

func NewService() Service { return &counterService{} }

func (c *counterService) Add(v int) int { 
    v += v
    return v
}

var logger log.Logger

func init() {
    logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
    logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}