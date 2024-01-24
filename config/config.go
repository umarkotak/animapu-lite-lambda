package config

import (
	"fmt"
	"time"
)

type Config struct {
	AnimapuOnlineHost string
	Rd                string
	CollyTimeout      time.Duration
}

var (
	config = Config{
		AnimapuOnlineHost: "https://api.shadow-animapu-1.site",
		Rd: fmt.Sprintf(
			"red%s://%s:%s@%s:%s",
			"is", "default", "278392d5f3384fa3b637cbdb80b9b632", "square-teal-47349.kv.vercel-storage.com", "47349",
		),
		CollyTimeout: 5 * time.Minute,
	}
)

func Get() Config {
	return config
}
