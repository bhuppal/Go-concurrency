// Package runner manages the running and lifetime of a process
/*
A concurrency pattern for task-oriented programs that run unattended on a schedule
It's designed with 3 possible termiation points:
1) The program can finish the work within the allotted amount of time and terminate normally
2) The program doesn't finish in time and kills itself
3) An operating system interrupt event is received and the program attempts to immediately shut down cleanly
 */
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)
// Runner runs a set of tasks within a given timeout and can be shutdown on an operating system interrupt
type Runner struct {
	// interrupt channel reports a signal from the operating system
	interrupt chan os.Signal

	//Complete channel reports that processing is done.
	complete chan error

	// timeout reports that time has run out
	timeout <-chan time.Time

	//tasks holds a set of functions that are executed synchronously in index order
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the OS is received
var ErrInterrupt = errors.New("received interrupt")


// New returns a new ready-to-use Runner
func New(d time.Duration) *Runner {
	return &Runner {
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}
// Add attaches tasks to the Runner. A task is a function that takes an int ID
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}
// Start runs all tasks and monitors channel events
func (r *Runner) Start() error {
	// We want to receive all interrupt based signals.
	signal.Notify(r.interrupt, os.Interrupt)

	// Run the different tasks on a different goroutine
	go func() {
		r.complete <- r.run()
	}()

	select {
		// signaled when processing is done
		case err := <-r.complete:
			return err

		// Signaled when we run out of time.
		case <- r.timeout:
			return ErrTimeout
	}
}
// run executes each registered task
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt signal from the OS.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// Execute the registered tasks
		task(id)
	}

	return nil
}

// gotInterrupt verifies if the interrupt signal has  been issued.
func (r *Runner) gotInterrupt() bool {
	select {
		// Signaled when an interrupt event is sent.
		case <- r.interrupt:
			// Stop receiving any further signals
			signal.Stop(r.interrupt)
			return true

			// Continue running as normal
			default:
				return false
	}
}