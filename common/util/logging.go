package util

import (
	"github.com/go-kit/kit/log"
	"shawn/gokbb_kit/services"
	"time"
)

type LoggingMiddleware struct {
	logger log.Logger
	next   services.StringService
}

func (mw LoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw LoggingMiddleware) Count(s string) (count int) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "count",
			"input", s,
			"n", count,
			"took", time.Since(begin),
		)
	}(time.Now())

	count = mw.next.Count(s)
	return
}
