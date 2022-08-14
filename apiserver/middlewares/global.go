package middlewares

import (
	"net/http"

	interpose "github.com/carbocation/interpose/middleware"
	"github.com/dre1080/recovr"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// HandlePanic handle panics from your API requests.
func HandlePanic(handler http.Handler) http.Handler {
	recovery := recovr.New()
	return recovery(handler)
}

// LogViaLogrus using interpose to integrate with logrus
func LogViaLogrus(handler http.Handler) http.Handler {
	logViaLogrus := interpose.NegroniLogrus()
	logrus.SetLevel(logrus.DebugLevel)
	return logViaLogrus(handler)
}

// Cross creates a new http.Handler that adds authentication logic to a given Handler
func Cross(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Debugln("Enter into cross handler")
		logrus.Debugln("request: ", r)

		if r.Method == "OPTIONS" {
			if r.Header.Get("Access-Control-Request-Method") != "" {
				logrus.Debugf("cors preflight detected")
				// cors preflight request/response
				w.Header().Add("Access-Control-Allow-Origin", "*")
				w.Header().Add("Access-Control-Allow-Methods", "*")
				w.Header().Add("Access-Control-Allow-Headers", "*")
				w.Header().Add("Access-Control-Max-Age", "86400")
				w.Header().Add("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(200)

				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
				return
			}
		}

		logrus.Debugln("Writing to header in callBefore \"Access-Control-Allow-Origin: *\"")

		w.Header().Add("Access-Control-Allow-Origin", "*")

		h.ServeHTTP(w, r)
	})
}

// UI default url host:port or host:port/ui or host:port/ui/
// func UI(handler http.Handler) http.Handler {
// 	// return http.FileServer(assetFS())
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path == "/dashboard" || r.URL.Path == "/" {
// 			http.Redirect(w, r, "/dashboard/", http.StatusFound)
// 			return
// 		}
// 		// Serving ./dashboard/
// 		if strings.Index(r.URL.Path, "/dashboard/") == 0 {
// 			fsys, err := fs.Sub(WebUI, "dist")
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusNotFound)
// 				return
// 			}
// 			http.StripPrefix("/dashboard/", http.FileServer(http.FS(fsys))).ServeHTTP(w, r)
// 			return
// 		}
// 		handler.ServeHTTP(w, r)
// 	})
// }

// RedocUI default url host:port/docs
func RedocUI(handler http.Handler) http.Handler {
	// return http.FileServer(assetFS())
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		opts := middleware.RedocOpts{
			SpecURL:  r.URL.Host + "/dashboard/static/swagger/swagger.yaml",
			RedocURL: r.URL.Host + "/dashboard/static/js/redoc.min.js",
			Title:    "User Manager",
		}
		middleware.Redoc(opts, handler).ServeHTTP(w, r)
	})
}
