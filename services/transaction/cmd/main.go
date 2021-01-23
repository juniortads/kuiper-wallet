package main

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main()  {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "transaction",
			"ts", log.DefaultTimestampUTC,
			"clr", log.DefaultCaller,
		)
	}
}
