package app

import (
	"sync"
)

const (
	PORT    = 8080
	DEBUG   = true
	VERSION = "0.1.0"
)

var (
	count    uint64
	once     sync.Once
	instance *Application
)
