package deploy

// Veth is a veth device
type Veth struct {
	Name string
}

// NewVeth returns a new veth device
func NewVeth() *Veth {
	return &Veth{}
}
