// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// GetWorkSubmitOKCode is the HTTP code returned for type GetWorkSubmitOK
const GetWorkSubmitOKCode int = 200

/*GetWorkSubmitOK A successful response.

swagger:response getWorkSubmitOK
*/
type GetWorkSubmitOK struct {

	/*
	  In: Body
	*/
	Payload *v1.WorkSubmit `json:"body,omitempty"`
}

// NewGetWorkSubmitOK creates GetWorkSubmitOK with default headers values
func NewGetWorkSubmitOK() *GetWorkSubmitOK {

	return &GetWorkSubmitOK{}
}

// WithPayload adds the payload to the get work submit o k response
func (o *GetWorkSubmitOK) WithPayload(payload *v1.WorkSubmit) *GetWorkSubmitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get work submit o k response
func (o *GetWorkSubmitOK) SetPayload(payload *v1.WorkSubmit) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWorkSubmitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
