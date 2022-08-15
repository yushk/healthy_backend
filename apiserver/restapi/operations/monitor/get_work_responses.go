// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// GetWorkOKCode is the HTTP code returned for type GetWorkOK
const GetWorkOKCode int = 200

/*GetWorkOK A successful response.

swagger:response getWorkOK
*/
type GetWorkOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Work `json:"body,omitempty"`
}

// NewGetWorkOK creates GetWorkOK with default headers values
func NewGetWorkOK() *GetWorkOK {

	return &GetWorkOK{}
}

// WithPayload adds the payload to the get work o k response
func (o *GetWorkOK) WithPayload(payload *v1.Work) *GetWorkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get work o k response
func (o *GetWorkOK) SetPayload(payload *v1.Work) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWorkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
