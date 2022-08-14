// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	v1 "gitee.com/healthy/backend/apiserver/v1"
	"github.com/go-openapi/runtime"
)

// DeleteMeasureDetailOKCode is the HTTP code returned for type DeleteMeasureDetailOK
const DeleteMeasureDetailOKCode int = 200

/*DeleteMeasureDetailOK A successful response.

swagger:response deleteMeasureDetailOK
*/
type DeleteMeasureDetailOK struct {

	/*
	  In: Body
	*/
	Payload *v1.MeasureDetail `json:"body,omitempty"`
}

// NewDeleteMeasureDetailOK creates DeleteMeasureDetailOK with default headers values
func NewDeleteMeasureDetailOK() *DeleteMeasureDetailOK {

	return &DeleteMeasureDetailOK{}
}

// WithPayload adds the payload to the delete measure detail o k response
func (o *DeleteMeasureDetailOK) WithPayload(payload *v1.MeasureDetail) *DeleteMeasureDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete measure detail o k response
func (o *DeleteMeasureDetailOK) SetPayload(payload *v1.MeasureDetail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMeasureDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
