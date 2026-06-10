package work

import (
	"context"
)

func ContextDemo() {
	ctx := context.Background()
	ctx.Done()
	// subCtx, func := context.WithTimeout(ctx, time.Duration(3, time.Second))
}