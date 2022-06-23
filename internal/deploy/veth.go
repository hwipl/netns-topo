package deploy

// Veth is a veth device
type Veth struct {
	Name string
}

// Start starts the veth device
func (v *Veth) Start() {
}

// Stop stops the veth device
func (v *Veth) Stop() {
}

// NewVeth returns a new veth device
func NewVeth() *Veth {
	return &Veth{}
}
