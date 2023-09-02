package server

import (
	"github.com/gin-gonic/gin"
)

type Fn struct {
	HTTPMethod   string
	RelativePath string
	Handlers     []gin.HandlerFunc
}
