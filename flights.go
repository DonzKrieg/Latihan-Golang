package main

import (
	"time"
)

type Flight struct {
	id            int
	aircraft      string
	origin        string
	destination   string
	price         float64
	departureDate time.Time
}
