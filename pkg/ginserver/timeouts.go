package ginserver

import "time"

const (
	DefTimeoutRead       = 50 * time.Second  // TLS + headers + body arrived
	DefTimeoutReadHeader = 5 * time.Second   // TLS + headers
	DefTimeoutWrite      = 100 * time.Second // Time Server has to respond (absolute)
	DefTimeoutIdle       = 120 * time.Second // Keepalives H1
)
