package synchronization

import "sync"

// Supervising routine can be used in the case where channel can have
// multi senders or multi consumers. This function will provide you
// with a channel, that is closed when all goroutines tied to the
// Wait Group are done.
func ActivateSupervisingRoutine(wg *sync.WaitGroup) chan int {
	ch := make(chan int)
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
