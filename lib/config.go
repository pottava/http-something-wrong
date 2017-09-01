package lib

import (
	"os"
	"strconv"
)

type config struct {
	Port            string // APP_PORT
	AccessLog       bool   // ACCESS_LOG
	ContentEncoding bool   // CONTENT_ENCODING
}

// Config represents its configurations
var Config *config

func init() {
	accessLog := true
	if b, err := strconv.ParseBool(os.Getenv("ACCESS_LOG")); err == nil {
		accessLog = b
	}
	contentEncoding := true
	if b, err := strconv.ParseBool(os.Getenv("CONTENT_ENCODING")); err == nil {
		contentEncoding = b
	}
	Config = &config{
		Port:            Env("APP_PORT", "80"),
		AccessLog:       accessLog,
		ContentEncoding: contentEncoding,
	}
}

// Env retrives a value from environmental variables
func Env(key, def string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return def
	}
	return value
}
