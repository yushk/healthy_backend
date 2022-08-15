// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/healthy_backend/apiserver/v1"
)

// GetMeasureResultsOKCode is the HTTP code returned for type GetMeasureResultsOK
const GetMeasureResultsOKCode int = 200

/*GetMeasureResultsOK ok

swagger:response getMeasureResultsOK
*/
type GetMeasureResultsOK struct {

	/*
	  In: Body
	*/
	Payload *v1.MeasureResults `json:"body,omitempty"`
}

// NewGetMeasureResultsOK creates GetMeasureResultsOK with default headers values
func NewGetMeasureResultsOK() *GetMeasureResultsOK {

	return &GetMeasureResultsOK{}
}

// WithPayload adds the payload to the get measure results o k response
func (o *GetMeasureResultsOK) WithPayload(payload *v1.MeasureResults) *GetMeasureResultsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get measure results o k response
func (o *GetMeasureResultsOK) SetPayload(payload *v1.MeasureResults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeasureResultsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
