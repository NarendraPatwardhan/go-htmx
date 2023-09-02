// Adapted from: https://github.com/golang/go/blob/master/src/context/context.go
package context

import (
	"context"
	"time"
)

// An internal type implementing context.Context.
type emptyCtx struct{}

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) { return }
func (*emptyCtx) Done() <-chan struct{}                   { return nil }
func (*emptyCtx) Err() error                              { return nil }
func (*emptyCtx) Value(key any) any                       { return nil }

var empty = new(emptyCtx)

// Empty returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline.
func Empty() context.Context { return empty }
