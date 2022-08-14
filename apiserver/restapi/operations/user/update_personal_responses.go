// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime"
)

// UpdatePersonalOKCode is the HTTP code returned for type UpdatePersonalOK
const UpdatePersonalOKCode int = 200

/*UpdatePersonalOK A successful response.

swagger:response updatePersonalOK
*/
type UpdatePersonalOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Personal `json:"body,omitempty"`
}

// NewUpdatePersonalOK creates UpdatePersonalOK with default headers values
func NewUpdatePersonalOK() *UpdatePersonalOK {

	return &UpdatePersonalOK{}
}

// WithPayload adds the payload to the update personal o k response
func (o *UpdatePersonalOK) WithPayload(payload *v1.Personal) *UpdatePersonalOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update personal o k response
func (o *UpdatePersonalOK) SetPayload(payload *v1.Personal) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePersonalOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
