package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	gateway "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/status"
)

func Error(err error) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		s := status.Convert(err)
		code := gateway.HTTPStatusFromCode(s.Code())
		if code == 500 {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(code)
		}
		payload, _ := json.Marshal(map[string]interface{}{
			"error":   s.Message(),
			"message": s.Message(),
			"code":    code,
			"details": s.Proto().GetDetails(),
		})
		w.Write(payload)
	})
}
