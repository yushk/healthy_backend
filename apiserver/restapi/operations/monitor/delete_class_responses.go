// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime"
)

// DeleteClassOKCode is the HTTP code returned for type DeleteClassOK
const DeleteClassOKCode int = 200

/*DeleteClassOK A successful response.

swagger:response deleteClassOK
*/
type DeleteClassOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Class `json:"body,omitempty"`
}

// NewDeleteClassOK creates DeleteClassOK with default headers values
func NewDeleteClassOK() *DeleteClassOK {

	return &DeleteClassOK{}
}

// WithPayload adds the payload to the delete class o k response
func (o *DeleteClassOK) WithPayload(payload *v1.Class) *DeleteClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete class o k response
func (o *DeleteClassOK) SetPayload(payload *v1.Class) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}