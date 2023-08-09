package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

func addDurationLoggerMiddleware(logger *zap.Logger, ep endpoint.Endpoint, epName string, elapsed time.Duration) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func(begin time.Time) {
			// _ = logger.Log(logconst.Layer, logconst.EndpointLayer, logconst.Endpoint, epName, "ElapsedTime", time.Since(begin))
		}(time.Now())

		return ep(ctx, request)
	}
}

func addErrLoggingMiddleware(logger *zap.Logger, ep endpoint.Endpoint, epName string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func() {
			if err != nil {
				// _ = logger.Log(logconst.Layer, logconst.EndpointLayer, logconst.Endpoint, epName, logconst.Error, err.Error())
			}
		}()

		return ep(ctx, request)
	}
}

// AddEndpointMiddleware ...
func AddEndpointMiddleware(logger *zap.Logger, ep endpoint.Endpoint, epName string, elapsed time.Duration) endpoint.Endpoint {
	ep = addErrLoggingMiddleware(logger, ep, epName)
	ep = addDurationLoggerMiddleware(logger, ep, epName, elapsed)
	return ep
}
