package gopool

const defaultScaleScaleThreshold = 1

type Config struct {
	ScaleThreshold int32
}

func NewConfig() *Config {
	c := &Config{
		ScaleThreshold: defaultScaleScaleThreshold,
	}

	return c
}
