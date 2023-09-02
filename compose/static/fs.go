package static

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// A wrapper around http.FileSystem for serving static files.
type FS struct {
	http.FileSystem
}

// Returns true iff an artifact has type file and exists at the given path.
func (e *FS) Exists(prefix string, path string) bool {
	f, err := e.Open(path)
	if err != nil {
		return false
	}

	// directory indexing is not supported
	s, _ := f.Stat()
	if s.IsDir() {
		return false
	}

	return true
}

// New creates a new static file system from the given folder implementing the http.FileSystem.
func New(folder embed.FS, path string) (*FS, error) {
	sub, err := fs.Sub(folder, path)
	if err != nil {
		return nil, err
	}
	return &FS{http.FS(sub)}, nil
}

// Serve checks existence of the file within static FS and serves it if it exists.
func Serve(urlPrefix string, fs FS) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if fs.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
