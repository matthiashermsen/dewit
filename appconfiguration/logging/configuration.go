package logging

import "log/slog"

type Configuration struct {
	Level slog.Level
}

func New() (Configuration, error) {
	level, err := GetLevel()

	if err != nil {
		return Configuration{}, err
	}

	c := Configuration{
		Level: level,
	}

	return c, err
}
