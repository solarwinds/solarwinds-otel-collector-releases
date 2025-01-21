package synchronization

import "context"

func IsContextClosed(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// Cancels given context ctx by given cancel function cancelFn
// passing cause of cancellation only if context ctx is not closed.
// If context was already closed nothing is done here.
func CancelContextWithCauseIfNotClosed(
	ctx context.Context,
	cancelFn context.CancelCauseFunc,
	cause error,
) {
	if !IsContextClosed(ctx) {
		cancelFn(cause)
	}
}
