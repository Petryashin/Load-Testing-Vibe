package config

import "time"

type Config struct {
	URL      string
	Workers  int
	Duration time.Duration
}
