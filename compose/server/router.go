package server

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"machinelearning.one/go-htmx/compose/static"
	"machinelearning.one/go-htmx/frontend"
)

// Internal function to create a production ready instance of gin router.
func createRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	return engine
}

func Run(ctx context.Context, port uint, fns ...Fn) error {
	// Initiate a custom instance of gin router.
	router := createRouter()

	tmpl := template.Must(
		template.New("").
			Delims("{{", "}}").
			Funcs(router.FuncMap).
			ParseFS(frontend.Content, "dist/*.html"))

	router.SetHTMLTemplate(tmpl)
	// Create a group for all api routes.
	api := router.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
		for _, fn := range fns {
			api.Handle(fn.HTTPMethod, fn.RelativePath, fn.Handlers...)
		}

	}
	// Serve the index.html file for all routes except the api ones.
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.HTML(http.StatusOK, "index.html", nil)
		}
	})

	content, err := static.New(frontend.Content, "dist")
	if err != nil {
		return err
	}
	router.Use(static.Serve("/", *content))

	// Create a new http server and attach the router to it.
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Start the server in a separate goroutine.
	go func() {
		srv.ListenAndServe()
	}()

	// Await for context cancellation.
	<-ctx.Done()
	srv.Shutdown(ctx)

	return nil
}
