package process

// Process interface
type Process interface {
	Wait() error
	Stop() error
}
