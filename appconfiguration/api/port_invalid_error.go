package api

type PortInvalidError struct{}

func (e *PortInvalidError) Error() string {
	return "Port must be in range between 1 and 65535."
}
