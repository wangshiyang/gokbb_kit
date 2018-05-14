package util

import (
	"github.com/go-kit/kit/log"
	"shawn/gokbb_kit/services"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   services.StringService
}

func (mw LoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Uppercase(s)
	return
}

func (mw LoggingMiddleware) Count(s string) (count int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", count,
			"took", time.Since(begin),
		)
	}(time.Now())

	count = mw.Next.Count(s)
	return
}
