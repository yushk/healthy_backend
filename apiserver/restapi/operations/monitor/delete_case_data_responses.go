// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime"
)

// DeleteCaseDataOKCode is the HTTP code returned for type DeleteCaseDataOK
const DeleteCaseDataOKCode int = 200

/*DeleteCaseDataOK A successful response.

swagger:response deleteCaseDataOK
*/
type DeleteCaseDataOK struct {

	/*
	  In: Body
	*/
	Payload *v1.CaseData `json:"body,omitempty"`
}

// NewDeleteCaseDataOK creates DeleteCaseDataOK with default headers values
func NewDeleteCaseDataOK() *DeleteCaseDataOK {

	return &DeleteCaseDataOK{}
}

// WithPayload adds the payload to the delete case data o k response
func (o *DeleteCaseDataOK) WithPayload(payload *v1.CaseData) *DeleteCaseDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete case data o k response
func (o *DeleteCaseDataOK) SetPayload(payload *v1.CaseData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCaseDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
