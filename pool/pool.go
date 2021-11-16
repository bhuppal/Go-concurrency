/*
The purpose of the pool package is to show how you can use a buffered channel
to pool a set of resources that can be shared and individually used by
any number of goroutines. This pattern is useful when you have a statis set of
resources to share, such as database connections or memory buffers.
When a goroutine needs one of these resources from the pool, it can acquire
the resource, use it, and then return it to the pool
 */

// Package pool manages a user defined set of resources
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manages a set of resources that can be shared safely by multiple
// goroutines. The resource being managed must implement the io.closer interface
type Pool struct {
	m sync.Mutex
	resources chan io.Closer
	factory func() (io.Closer, error)
	closed bool
}

// ErrPoolClosed is returned when an acquire returns on a closed pool
var ErrPoolClosed = errors.New("Pool has been closed.")

// New creates a pool that manages resources. A pool requires a function that can
// allocate a new resource and the size of the pool.
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory: fn,
		resources: make(chan io.Closer, size),
	}, nil
}


// Acquire retrieves a resource from the pool
func (p *Pool) Acquire() (io.Closer, error) {
	select  {
	case r, ok := <-p.resources :
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release places a new resource onto the pool.
func (p *Pool) Release(r io.Closer){
	// Secure this operation with the Close operation.
	p.m.Lock()
	defer p.m.Unlock()

	// If the pool is closed, discard the resource
	if p.closed {
		r.Close()
		return
	}

	select {
	// Attempt to place the new resource on the queue
		case p.resources <- r:
			log.Println("Release:", "In Queue")

		// If the queue is already at capacity we close the resource
		default:
			log.Println("Release:", "Closing")
			r.Close()
	}
}

// Close will shutdown the pool and close all existing resources
func (p *Pool) Close() {
	// Secure this operation with the Release operation
	p.m.Lock()
	defer  p.m.Unlock()

	// If the pool is already closed, don't do anything.
	if p.closed {
		return
	}

	// Set the pool as closed.
	p.closed = true

	// Close the channel before we drain the channel of its
	// resources. If we don't do this, we will have a deadlock
	close(p.resources)

	// Close the resources
	for r := range p.resources {
		r.Close()
	}
}