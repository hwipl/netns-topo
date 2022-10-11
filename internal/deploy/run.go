package deploy

// Run is a list of commands to run in a network namespace
type Run struct {
	Netns    string
	Commands []string
}

// Start executes the commands
func (r *Run) Start() {
	for _, c := range r.Commands {
		runNetns(r.Netns, "/bin/bash", "-c", c)
	}
}

// Stop stops the commands
func (r *Run) Stop() {
}

// NewRun returns a new list of commands to run
func NewRun() *Run {
	return &Run{}
}
