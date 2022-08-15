// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// GetCaseDatasOKCode is the HTTP code returned for type GetCaseDatasOK
const GetCaseDatasOKCode int = 200

/*GetCaseDatasOK ok

swagger:response getCaseDatasOK
*/
type GetCaseDatasOK struct {

	/*
	  In: Body
	*/
	Payload *v1.CaseDatas `json:"body,omitempty"`
}

// NewGetCaseDatasOK creates GetCaseDatasOK with default headers values
func NewGetCaseDatasOK() *GetCaseDatasOK {

	return &GetCaseDatasOK{}
}

// WithPayload adds the payload to the get case datas o k response
func (o *GetCaseDatasOK) WithPayload(payload *v1.CaseDatas) *GetCaseDatasOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get case datas o k response
func (o *GetCaseDatasOK) SetPayload(payload *v1.CaseDatas) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaseDatasOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
