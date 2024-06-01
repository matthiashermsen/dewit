package cfg

import "os"

const (
	PortKey     = "PORT"
	DefaultPort = "8080"
)

func GetPort() string {
	port, hasPort := os.LookupEnv(PortKey)

	if !hasPort || port == "" {
		return DefaultPort
	}

	return port
}
