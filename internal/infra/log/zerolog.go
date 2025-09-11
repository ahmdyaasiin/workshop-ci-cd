package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZeroLog() error {
	timezone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeFormat:   time.RFC1123,
		PartsOrder:   []string{"level", "message", "time"},
		TimeLocation: timezone,
		FormatCaller: func(i interface{}) string {
			var c string
			if cc, ok := i.(string); ok {
				c = cc
			}
			if len(c) > 0 {
				if cwd, err := os.Getwd(); err == nil {
					if rel, err := filepath.Rel(cwd, c); err == nil {
						c = rel
					}
				}
			}
			if c != "" {
				return fmt.Sprintf("[%v]", c)
			}
			return c
		},
	}
	log.Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()

	return nil
}
