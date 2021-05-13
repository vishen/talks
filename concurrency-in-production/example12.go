package main

import (
	"context"
)

// START OMIT
func (e EventHandler) HandleMenuOutdated(ctx context.Context, ev interface{}, h ...eventbus.Header) {
	// Kick of an event to handle the long running task
	_ = e.producer.Emit(ctx, eventForLongRunningTask)

	// Finish other work
	return nil
}

// END OMIT
