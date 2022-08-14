// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime"
)

// DeleteMeasureResultOKCode is the HTTP code returned for type DeleteMeasureResultOK
const DeleteMeasureResultOKCode int = 200

/*DeleteMeasureResultOK A successful response.

swagger:response deleteMeasureResultOK
*/
type DeleteMeasureResultOK struct {

	/*
	  In: Body
	*/
	Payload *v1.MeasureResult `json:"body,omitempty"`
}

// NewDeleteMeasureResultOK creates DeleteMeasureResultOK with default headers values
func NewDeleteMeasureResultOK() *DeleteMeasureResultOK {

	return &DeleteMeasureResultOK{}
}

// WithPayload adds the payload to the delete measure result o k response
func (o *DeleteMeasureResultOK) WithPayload(payload *v1.MeasureResult) *DeleteMeasureResultOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete measure result o k response
func (o *DeleteMeasureResultOK) SetPayload(payload *v1.MeasureResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMeasureResultOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}