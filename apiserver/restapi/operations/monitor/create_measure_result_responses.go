// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// CreateMeasureResultOKCode is the HTTP code returned for type CreateMeasureResultOK
const CreateMeasureResultOKCode int = 200

/*CreateMeasureResultOK ok

swagger:response createMeasureResultOK
*/
type CreateMeasureResultOK struct {

	/*
	  In: Body
	*/
	Payload *v1.MeasureResult `json:"body,omitempty"`
}

// NewCreateMeasureResultOK creates CreateMeasureResultOK with default headers values
func NewCreateMeasureResultOK() *CreateMeasureResultOK {

	return &CreateMeasureResultOK{}
}

// WithPayload adds the payload to the create measure result o k response
func (o *CreateMeasureResultOK) WithPayload(payload *v1.MeasureResult) *CreateMeasureResultOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create measure result o k response
func (o *CreateMeasureResultOK) SetPayload(payload *v1.MeasureResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMeasureResultOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
