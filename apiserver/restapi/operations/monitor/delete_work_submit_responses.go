// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// DeleteWorkSubmitOKCode is the HTTP code returned for type DeleteWorkSubmitOK
const DeleteWorkSubmitOKCode int = 200

/*DeleteWorkSubmitOK A successful response.

swagger:response deleteWorkSubmitOK
*/
type DeleteWorkSubmitOK struct {

	/*
	  In: Body
	*/
	Payload *v1.WorkSubmit `json:"body,omitempty"`
}

// NewDeleteWorkSubmitOK creates DeleteWorkSubmitOK with default headers values
func NewDeleteWorkSubmitOK() *DeleteWorkSubmitOK {

	return &DeleteWorkSubmitOK{}
}

// WithPayload adds the payload to the delete work submit o k response
func (o *DeleteWorkSubmitOK) WithPayload(payload *v1.WorkSubmit) *DeleteWorkSubmitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete work submit o k response
func (o *DeleteWorkSubmitOK) SetPayload(payload *v1.WorkSubmit) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteWorkSubmitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
