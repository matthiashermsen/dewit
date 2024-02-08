package api

type Configuration struct {
	Port uint16
}

func New() Configuration {
	port := GetPort()
	c := Configuration{
		Port: port,
	}

	return c
}
