package api_go_sdk

// Currently, unused error struct
type assertionError struct {
	s string
}

func (aerror *assertionError) Error() string {
	return aerror.s
}