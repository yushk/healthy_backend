package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/sirupsen/logrus"
)

// Limiter  rate limit the valid requests to our swagger API.
func Limiter(handler http.Handler) http.Handler {
	lmt := tollbooth.NewLimiter(100, &limiter.ExpirableOptions{})
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	// Set a custom function for rejection.
	lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) { logrus.Errorln("A request was rejected") })
	return tollbooth.LimitHandler(lmt, handler)
}
