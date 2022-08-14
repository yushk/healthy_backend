// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime"
)

// CreateWorkSubmitOKCode is the HTTP code returned for type CreateWorkSubmitOK
const CreateWorkSubmitOKCode int = 200

/*CreateWorkSubmitOK ok

swagger:response createWorkSubmitOK
*/
type CreateWorkSubmitOK struct {

	/*
	  In: Body
	*/
	Payload *v1.WorkSubmit `json:"body,omitempty"`
}

// NewCreateWorkSubmitOK creates CreateWorkSubmitOK with default headers values
func NewCreateWorkSubmitOK() *CreateWorkSubmitOK {

	return &CreateWorkSubmitOK{}
}

// WithPayload adds the payload to the create work submit o k response
func (o *CreateWorkSubmitOK) WithPayload(payload *v1.WorkSubmit) *CreateWorkSubmitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create work submit o k response
func (o *CreateWorkSubmitOK) SetPayload(payload *v1.WorkSubmit) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkSubmitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
