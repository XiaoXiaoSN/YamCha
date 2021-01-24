package helper

import (
	"fmt"
	"runtime"
	"time"
	"yamcha/internal/config"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

// CmdRecover ...
func CmdRecover() {
	if r := recover(); r != nil {
		var msg string
		for i := 2; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			msg = msg + fmt.Sprintf("%s:%d\n", file, line)
			sentry.CaptureException(fmt.Errorf("panic: %+v", msg))
		}
		log.Errorf("==========\nPANIC: %s\n%s\n==========", msg, r)

		if config.Config().Sentry.SentryDSN != "" {
			sentry.Flush(time.Second * 5)
		}
	}
}
