package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"machinelearning.one/go-htmx/compose/context"
	"machinelearning.one/go-htmx/compose/logger"
	"machinelearning.one/go-htmx/compose/server"
)

func main() {
	ctx := context.Context()
	lg := logger.New("trace")
	ctx = lg.WithContext(ctx)
	port := uint(8080)

	data := struct {
		Count int
	}{0}

	fn := server.Fn{
		HTTPMethod:   "POST",
		RelativePath: "/clicked",
		Handlers: []gin.HandlerFunc{
			func(c *gin.Context) {
				data.Count += 1
				lg.Trace().Int("count", data.Count).Msg("clicked")
				c.HTML(http.StatusOK, "clicked.frag.html", data)
			},
		},
	}

	lg.Trace().Msgf("starting server on port %d", port)
	err := server.Run(ctx, port, fn)
	if err != nil {
		lg.Fatal().Err(err).Msg("failed to start server")
	}
}
