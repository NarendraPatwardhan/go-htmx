// Adapted from: https://github.com/moby/buildkit/blob/master/util/appcontext/appcontext.go
package context

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"machinelearning.one/go-htmx/compose/logger"
)

var (
	contextCache context.Context
	syncOnce     sync.Once
)

// Context returns a context.Context that is canceled when the process receives
// an os-specific termination signal. The result is cached and reused on subsequent calls.
// Suitable for use in main() as the top-level context
func Context() context.Context {
	syncOnce.Do(func() {
		signals := make(chan os.Signal, 2048)
		signal.Notify(signals, terminationSignals...)

		const exitLimit = 3
		retries := 0

		ctx := Empty()

		ctx, cancel := context.WithCancel(ctx)
		contextCache = ctx

		go func() {
			for {
				<-signals
				cancel()
				retries++
				if retries >= exitLimit {
					lg := logger.New(logger.DefaultLevel)
					lg.Info().Msg("received termination signal, exiting")
					os.Exit(1)
				}
			}
		}()
	})
	return contextCache
}
