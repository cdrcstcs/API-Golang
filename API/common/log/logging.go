package log

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// TODO update the log package here

func CtxInfo(ctx context.Context, message string, any ...interface{}) {
	message = fmt.Sprintf(message, any...)
	log.Info().Ctx(ctx).Str("log_message", message)
}

//func Info()
