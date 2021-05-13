package main

import (
	"context"
)

// START OMIT
func (e EventHandler) HandleMenuOutdated(ctx context.Context, ev interface{}, h ...eventbus.Header) {
	// Don't pass the context to a long-running task.
	// The Event Handler cancellation timer it attached to this ctx.
	go longRunningTask(ctx)
	// Finish other work
	return nil
}

func longRunningTask() {
	// Do some long-running task
}

// END OMIT
