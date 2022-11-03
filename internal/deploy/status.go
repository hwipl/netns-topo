package deploy

// Status codes for the deployment status
const (
	StatusUnknown Status = iota
	StatusInactive
	StatusPartial
	StatusActive
	StatusInvalid
)

// Status is a deployment status
type Status uint8

// String returns the status as a string
func (s Status) String() string {
	switch s {
	case StatusUnknown:
		return "unknown"
	case StatusInactive:
		return "inactive"
	case StatusPartial:
		return "partial"
	case StatusActive:
		return "active"
	default:
		return ""
	}
}
