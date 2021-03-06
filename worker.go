package thdpool

import (
	"sync"
)

// Worker is a interface include a Work function that called in thread pool.
type Worker interface {
	Work(thdID int, mutex sync.Locker) error
}
